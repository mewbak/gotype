package style

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/andybalholm/cascadia"
	"github.com/npillmayer/gotype/core/config/tracing"
	"golang.org/x/net/html"
)

/*
TODO
+ implement shortcut-properties: "border", "background" etc.
+ style from header: link and sytle tags
+ Locally scoped style (inline or <style> in body)
+ Specifity
+ be independent from goquery
+ be independent form douceur.css
- Fix API:
  + create StyledTree using factory
  + define extension points for future properties: group "X"; registerCompound(key, func(...))
- write lots of tests
- make concurrent
+ Document code
- Create Diagram -> Wiki
- API for export to DOT
*/

/* -----------------------------------------------------------------
BSD License

Copyright (c) 2017–18, Norbert Pillmayer

All rights reserved.

Redistribution and use in source and binary forms, with or without
modification, are permitted provided that the following conditions
are met:

1. Redistributions of source code must retain the above copyright
notice, this list of conditions and the following disclaimer.

2. Redistributions in binary form must reproduce the above copyright
notice, this list of conditions and the following disclaimer in the
documentation and/or other materials provided with the distribution.

3. Neither the name of Norbert Pillmayer nor the names of its contributors
may be used to endorse or promote products derived from this software
without specific prior written permission.

THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
"AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
(INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
----------------------------------------------------------------- */

// CSSOM is the "CSS Object Model", similar to the DOM for HTML.
// Our CSSOM consists of a set of stylesheets, each relevant for a sub-tree
// of the DOM. This sub-tree is called the "scope" of the stylesheet.
// Sub-trees are identified through the top node.
//
// Stylesheets are wrapped into an internal rules tree.
type CSSOM struct {
	rules             map[*html.Node]rulesTree
	defaultProperties *PropertyMap
	compoundSplitters []CompoundPropertiesSplitter
}

// RulesTree holds the styling rules of a stylesheet.
//
// Status: Currently this is not really a tree.
// Optimize some day (see
// https://hacks.mozilla.org/2017/08/inside-a-super-fast-css-engine-quantum-css-aka-stylo/).
type rulesTree struct {
	stylesheet StyleSheet                   // stylesheet holding the rules
	selectors  map[string]cascadia.Selector // cache of compiled selectors
	source     PropertySource               // where do these rules come from?
}

// Compound properties are properties which abbreviate the
// setting of more fine grained propertes. An example is
//    padding: 10px 20px
// which sets the following detail properties:
//    padding-top:    10px
//    padding-right:  20px
//    padding-bottom: 10px
//    padding-left:   20px
// Standard CSS compound properties are known by default, but clients are
// allowed to extend the set of compound properties.
type CompoundPropertiesSplitter func(string, Property) ([]KeyValue, error)

// NewCSSOM creates an empty CSSOM.
// Clients are allowed to supply a map of additional/custom CSS property values.
// These may override values of the default ("user-agent") style sheet,
// or introduce completely new styling properties.
func NewCSSOM(additionalProperties []KeyValue) CSSOM {
	cssom := CSSOM{}
	cssom.rules = make(map[*html.Node]rulesTree)
	cssom.defaultProperties = initializeDefaultPropertyValues(additionalProperties)
	cssom.compoundSplitters = make([]CompoundPropertiesSplitter, 1)
	cssom.compoundSplitters[0] = splitCompoundProperty
	return cssom
}

// Properties may be defined at different places in HTML: as a sytlesheet
// reference link, within a <script> element in the HTML file, or in an
// attribute value.
//
// PropertySource affects the specifity of rules.
type PropertySource uint8

// Values for property sources, used when adding style sheets.
const (
	Global    PropertySource = iota + 1 // "browser" globals
	Author                              // CSS author (stylesheet link)
	Script                              // <script> element
	Attribute                           // in an element's attribute(s)
)

// AddStylesFor includes a stylesheet to a CSSOM and sets the scope for
// the stylesheet. If a stylesheet for the scope already exists, the
// styles are merged. css may be nil. If scope is nil then scope is the
// body (i.e., top-level content element) of a future document.
//
// source hints to where the stylesheet comes from.
// Its value will affect the calculation of specifity for rules of this
// stylesheet.
func (cssom CSSOM) AddStylesFor(scope *html.Node, css StyleSheet, source PropertySource) error {
	rules, exists := cssom.rules[scope]
	if exists && css != nil {
		rules.stylesheet.AppendRules(css)
	} else {
		if scope == nil {
			scope = bodyElement
		} else {
			if scope.Type != html.ElementNode {
				return errors.New("Can style element nodes only")
			}
		}
		cssom.rules[scope] = rulesTree{css, nil, source}
	}
	return nil
}

// bodyElement is a symbolic node to denote the body element of a future
// HTML document. AddStylesFor(...) with nil as a scope will replace it
// with this marker for scoping the complete document body.
var bodyElement *html.Node = &html.Node{}

// RegisterCompoundSplitter allows clients to handle additional compound
// properties. See type CompoundPropertiesSplitter.
func (cssom CSSOM) RegisterCompoundSplitter(splitter CompoundPropertiesSplitter) {
	if splitter != nil {
		cssom.compoundSplitters = append(cssom.compoundSplitters, splitter)
	}
}

// Empty is a predicate wether a stylesheet is empty, i.e. does not contain
// any rules.
func (rt rulesTree) Empty() bool {
	return rt.stylesheet == nil || rt.stylesheet.Empty()
}

// Internal helper for applying rules to an HTML node.
// In a first step it holds all the rules matching for node.
// In a second step it collects all the properties set in those rules,
// then orderes them by specifity.
type matchesList struct {
	matchingRules   []Rule
	propertiesTable []specifity
}

// Rules matches are collected from more than one stylesheet. Matching
// rules from these stylesheets will be merged to one list.
func (mlist *matchesList) mergeMatchesWith(m *matchesList) *matchesList {
	if mlist == nil {
		return m
	}
	if m != nil {
		for _, r := range m.matchingRules {
			mlist.matchingRules = append(mlist.matchingRules, r)
		}
	}
	return mlist
}

// sorter
type byHighestSpecifity []specifity

// make specifities sortable by highest sp.spec
func (sp byHighestSpecifity) Len() int           { return len(sp) }
func (sp byHighestSpecifity) Swap(i, j int)      { sp[i], sp[j] = sp[j], sp[i] }
func (sp byHighestSpecifity) Less(i, j int) bool { return sp[i].spec > sp[j].spec }

func (m *matchesList) String() string {
	s := fmt.Sprintf("match of %d rules:\n", len(m.matchingRules))
	s += "Src +-- Spec. --+------------- Key --------------+------- Value ---------------\n"
	for _, sp := range m.propertiesTable {
		s += fmt.Sprintf("%3d | %9d | %30s | %s\n", sp.source, sp.spec, sp.propertyKey, sp.propertyValue)
	}
	return s
}

func (rt *rulesTree) FilterMatchesFor(node *html.Node) *matchesList {
	list := &matchesList{}
	for _, rule := range rt.stylesheet.Rules() {
		selectorString := rule.Selector()
		var sel cascadia.Selector
		found := false
		if sel, found = rt.selectors[selectorString]; !found {
			var err error
			sel, err = cascadia.Compile(selectorString)
			if err != nil {
				tracing.EngineTracer.Errorf("CSS selector seems not to work: %s", selectorString)
				break
			}
			if rt.selectors == nil {
				rt.selectors = make(map[string]cascadia.Selector)
			}
			rt.selectors[selectorString] = sel
		}
		if sel.Match(node) {
			list.matchingRules = append(list.matchingRules, rule)
		}
	}
	tracing.EngineTracer.Debugf("matching rules for %s", nodePath(node))
	return list
}

func (matches *matchesList) SortProperties() {
	var proptable []specifity
	for rno, rule := range matches.matchingRules {
		for _, propertyKey := range rule.Properties() {
			value := Property(rule.Value(propertyKey))
			props, err := splitCompoundProperty(propertyKey, value)
			if err != nil {
				sp := specifity{Author, rule, propertyKey, value, rule.IsImportant(propertyKey), 0}
				sp.calcSpecifity(rno)
				proptable = append(proptable, sp)
			} else {
				tracing.EngineTracer.Debugf("%s is a compound style", propertyKey)
				for _, kv := range props {
					key := kv.Key
					val := kv.Value
					sp := specifity{Author, rule, key, val, rule.IsImportant(propertyKey), 0}
					sp.calcSpecifity(rno)
					proptable = append(proptable, sp)
				}
			}
		}
	}
	if len(proptable) > 0 {
		sort.Sort(byHighestSpecifity(proptable))
		matches.propertiesTable = proptable
	}
	if tracing.EngineTracer.GetTraceLevel() >= tracing.LevelDebug {
		tracing.EngineTracer.Debugf(matches.String())
	}
}

func (matches *matchesList) createStyleGroups(parent StyledNode) map[string]*propertyGroup {
	m := make(map[string]*propertyGroup)
	done := make(map[string]bool, len(matches.propertiesTable))
	for _, pspec := range matches.propertiesTable { // for every specifity entry
		groupname, found := groupNameFromPropertyKey[pspec.propertyKey]
		if !found {
			tracing.EngineTracer.Debugf("Don't know about CSS property: %s", pspec.propertyKey)
			groupname = "X" // assume client uses custom properties
		}
		if done[pspec.propertyKey] {
			// already present in current properties map
			// this must be from previous set with higher specifity
			// => do nothing
			break
		}
		if group, exists := m[groupname]; exists {
			group.Set(pspec.propertyKey, pspec.propertyValue)
		} else {
			_, pg := findAncestorWithPropertyGroup(parent, groupname) // must succeed
			if pg == nil {
				panic(fmt.Sprintf("Cannot find ancestor with prop-group %s -- did you create global properties?", groupname))
			}
			group, isNew := pg.SpawnOn(pspec.propertyKey, pspec.propertyValue, true)
			if isNew { // a new property group has been created
				m[groupname] = group // put it into the group map
			}
		}
		done[pspec.propertyKey] = true // remember we're done with this property
	}
	if len(m) == 0 { // no property groups created, not properties set
		return nil
	}
	if tracing.EngineTracer.GetTraceLevel() >= tracing.LevelDebug {
		for _, v := range m {
			tracing.EngineTracer.Debugf(v.String())
		}
	}
	return m
}

// --- Specifity of rules -----------------------------------------------

type specifity struct {
	source        PropertySource // where the property has been defined
	rule          Rule           // the rule containing the property definition
	propertyKey   string         // CSS property name
	propertyValue Property       // raw string value
	important     bool           // marked as !IMPORTANT ?
	spec          uint32         // specifity value to calculate; higher is more
}

// CalcSpecifity calculates an appromiation to the true W3C specifity.
// https://www.smashingmagazine.com/2007/07/css-specificity-things-you-should-know/
//
// no is a sequence number for rules, ensuring that later rules override
// previously defined rules / properties.
func (sp *specifity) calcSpecifity(no int) {
	if sp.rule.IsImportant(sp.propertyKey) {
		sp.spec = 99999 // max
		return
	}
	sp.spec = uint32(sp.source-1) * 1000
	selectorstring := sp.rule.Selector()
	// simple "parsing" = rough estimate...
	// alternatively use code from cascadia or from
	// https://godoc.org/github.com/ericchiang/css
	sels := strings.Fields(selectorstring)
	var selcnt uint32
	var idcnt uint32
	var classcnt uint32
	for _, sel := range sels {
		selcnt++
		if strings.ContainsRune(sel, ':') {
			selcnt++ // count double
		}
		if strings.ContainsAny(sel, ".[:") {
			classcnt++
		}
		if strings.HasPrefix(sel, "#") {
			idcnt++
		}
	}
	sp.spec += selcnt*10 + classcnt*100 + idcnt*1000 + uint32(no)
}

// --- Styled Node Tree -------------------------------------------------

// StyledNode is the node type for the styled tree to build. The CSSOM
// styler builds a tree of styled nodes, more or less corresponding to the
// stylable DOM nodes. We de-couple the Styler from the type of the resulting
// styled nodes by using a factory (see NodeFactory) and an interface type.
type StyledNode interface {
	Parent() StyledNode
	LinkToParent(StyledNode)
	ComputedStyles() *PropertyMap
	SetComputedStyles(*PropertyMap)
}

// NodeFactory is a factory to create conrete implementations of interface
// StyleNode (see type StyledNode).
type NodeFactory interface {
	NodeFor(*html.Node) StyledNode
}

func setupStyledNodeTree(domRoot *html.Node, defaults *PropertyMap, factory NodeFactory) StyledNode {
	viewport := factory.NodeFor(domRoot)
	viewport.SetComputedStyles(defaults)
	return viewport
}

func findAncestorWithPropertyGroup(sn StyledNode, group string) (StyledNode, *propertyGroup) {
	var pg *propertyGroup
	it := sn
	for it != nil && pg == nil {
		pg = it.ComputedStyles().m[group]
		it = it.Parent()
	}
	return it, pg
}

// Style() gets things rolling. It styles a DOM, referred to by the root
// node, and returns a tree of styled nodes.
// For an explanation what's going on here, refer to
// https://hacks.mozilla.org/2017/08/inside-a-super-fast-css-engine-quantum-css-aka-stylo/
// and
// https://limpet.net/mbrubeck/2014/08/23/toy-layout-engine-4-style.html
//
// If either dom or factory are nil, no tree is returned (but an error).
func (cssom CSSOM) Style(dom *html.Node, factory NodeFactory) (StyledNode, error) {
	if dom == nil {
		return nil, errors.New("Nothing to style: empty document")
	}
	if factory == nil {
		return nil, errors.New("Cannot style: no factory to create styles nodes")
	}
	styledRootNode := setupStyledNodeTree(dom, cssom.defaultProperties, factory)
	runner, err := cssom.attachStylesheets(dom)
	if err != nil {
		return nil, err
	}
	runner.splitters = cssom.compoundSplitters // hack, sorry
	runner.doStyle(runner.startNode, styledRootNode, factory)
	return styledRootNode, nil
}

// Scopes for all ruletrees have to be HTML element nodes.
// This is enforced during construction of the rules trees (adding to CSSDOM).
func (cssom CSSOM) attachStylesheets(dom *html.Node) (*stylingRunner, error) {
	if dom == nil {
		return nil, errors.New("Nothing to style: empty document")
	}
	runner := newStylingRunner()
	for scope, rulestree := range cssom.rules {
		var stylingRootElement *html.Node
		if scope == bodyElement { // scope is body element, i.e. whole document
			stylingRootElement := findBodyElement(dom)
			if stylingRootElement == nil { // no body element found
				tracing.EngineTracer.Infof("Misconstructed DOM: cannot find <body>. Proceeding.")
				stylingRootElement = dom // root of fragment, try our best
			}
			runner.activeStylers[stylingRootElement] = rulestree
			runner.startNode = stylingRootElement
		} else {
			stylingRootElement = findThisNode(dom, scope)
			if stylingRootElement == nil { // scope not not found in DOM
				// do not try to proceed: client meant something else...
				return nil, errors.New(fmt.Sprintf("Scope '%s' not found in DOM", scope.Data))
			}
			if stylingRootElement.Type != html.ElementNode {
				return nil, errors.New(fmt.Sprintf("Scope '%s' is not of element type", shorten(scope.Data)))
			}
			runner.inactiveStylers[stylingRootElement] = rulestree
		}
	}
	if runner.startNode == nil {
		runner.startNode = dom
	}
	return runner, nil
}

type stylingRunner struct {
	activeStylers   map[*html.Node]rulesTree
	inactiveStylers map[*html.Node]rulesTree
	startNode       *html.Node
	splitters       []CompoundPropertiesSplitter
}

func newStylingRunner() *stylingRunner {
	runner := &stylingRunner{}
	runner.activeStylers = make(map[*html.Node]rulesTree)
	runner.inactiveStylers = make(map[*html.Node]rulesTree)
	return runner
}

func (runner *stylingRunner) activateStylesheetsFor(node *html.Node) {
	for scope, r := range runner.inactiveStylers {
		if scope == node {
			runner.activeStylers[scope] = r
			delete(runner.inactiveStylers, scope)
		}
	}
}

func (runner *stylingRunner) doStyle(node *html.Node, parent StyledNode, factory NodeFactory) {
	runner.activateStylesheetsFor(node)
	if createsStyledNode(node.Type) {
		sn := factory.NodeFor(node)
		sn.LinkToParent(parent)
		var matchingRules *matchesList
		for _, rulesTree := range runner.activeStylers {
			matches := rulesTree.FilterMatchesFor(node)
			matchingRules = matchingRules.mergeMatchesWith(matches)
		}
		if matchingRules != nil {
			matchingRules.SortProperties()
			groups := matchingRules.createStyleGroups(parent)
			sn.SetComputedStyles(&PropertyMap{groups}) // groups may be nil
		}
		parent = sn // continue with newly created styled node
	}
	node = node.FirstChild // now recurse into children
	for node != nil {
		if node.Type == html.ElementNode {
			runner.doStyle(node, parent, factory) // TODO make this concurrent
		}
		node = node.NextSibling
	}
}

func (runner *stylingRunner) splitCompoundProperties(key string, value Property) ([]KeyValue, error) {
	for _, splitter := range runner.splitters {
		kv, err := splitter(key, value)
		if err != nil {
			return kv, err
		}
	}
	return nil, errNoSuchCompoundProperty
}

var errNoSuchCompoundProperty error = errors.New("No such compound property")

// ----------------------------------------------------------------------

// GetCascaded gets the value of a property. The search cascades to
// parent property maps, if available.
//
// This is normally called on a tree of styled nodes and it will cascade
// all the way up to the default properties, if necessary.
func GetCascadedProperty(sn StyledNode, key string) Property {
	group, found := groupNameFromPropertyKey[key]
	if !found {
		group = "X"
	}
	_, pg := findAncestorWithPropertyGroup(sn, group) // must succeed
	if pg == nil {
		panic(fmt.Sprintf("Cannot find ancestor with prop-group %s -- did you create global properties?", group))
	}
	return pg.Cascade(key).Get(key) // must succeed
}

// --- Helpers ----------------------------------------------------------

// Which HTML node type needs a corresponding styled node?
func createsStyledNode(nodeType html.NodeType) bool {
	if nodeType == html.ElementNode || nodeType == html.TextNode {
		return true
	}
	return false
}

// Helper to find nodes matching a predicate. Currently works recursive.
// Returns a node or nil.
func findNode(node *html.Node, matcher func(n *html.Node) bool) *html.Node {
	if node == nil {
		return nil
	}
	if matcher(node) {
		return node
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		if f := findNode(c, matcher); f != nil {
			return f
		}
	}
	return nil
}

func findThisNode(tree *html.Node, nodeToFind *html.Node) *html.Node {
	return findNode(tree, func(n *html.Node) bool {
		return n == nodeToFind
	})
}

func findBodyElement(tree *html.Node) *html.Node {
	return findNode(tree, func(n *html.Node) bool {
		return n.Type == html.ElementNode && n.Data == "body"
	})
}

// shorten a string
func shorten(s string) string {
	if len(s) > 10 {
		return s[:10] + "..."
	}
	return s
}

// small helper to debug-print out a node. TODO
func nodePath(node *html.Node) string {
	s := ""
	if node.Type == html.TextNode {
		s += "(text)"
	} else if node.Type == html.ElementNode {
		s += fmt.Sprintf("%s", node.Data)
	} else {
		s += "(unknown)"
	}
	return s
}
