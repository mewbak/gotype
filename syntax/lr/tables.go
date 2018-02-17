package lr

import (
	"fmt"
	"io"
	"os"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/emirpasic/gods/sets/hashset"
	"github.com/emirpasic/gods/sets/treeset"
	"github.com/emirpasic/gods/utils"
	"github.com/npillmayer/gotype/syntax/lr/sparse"
	"github.com/npillmayer/gotype/syntax/runtime"
)

// TODO: Improve documentation...

// === Items and Item Sets ===================================================

// A set of Earley items ( A -> B * C D ).
type itemSet struct {
	*hashset.Set
}

func newItemSet() *itemSet {
	s := hashset.New()
	iset := &itemSet{s}
	return iset
}

func (iset *itemSet) union(iset2 *itemSet) {
	iset.Add(iset2.Values()...)
}

func (iset *itemSet) equals(iset2 *itemSet) bool {
	if iset.Size() == iset2.Size() {
		for _, i := range iset.Values() {
			if !iset2.Contains(i) {
				return false
			}
		}
		return true
	}
	return false
}

func (iset *itemSet) String() string {
	items := iset.Values()
	if len(items) == 0 {
		return "{ }"
	} else {
		s := "{\n"
		for _, i := range items {
			s = s + fmt.Sprintf("\t%v\n", i)
		}
		s += "}"
		return s
	}
}

// Prepare an item set for export to Graphviz.
func (iset *itemSet) forGraphviz() string {
	items := iset.Values()
	if len(items) == 0 {
		return "err\\n"
	} else {
		s := ""
		for _, i := range items {
			s = s + fmt.Sprintf("%v\\l", i)
		}
		return s
	}
}

// Debugging helper
func (iset *itemSet) Dump() {
	items := iset.Values()
	//T.Debug("--- item set ------------")
	for k, i := range items {
		T.Debugf("item %2d = %v", k, i)
	}
	//T.Debug("-------------------------")
}

var _ *itemSet = newItemSet() // verify assignability

// === Closure and Goto-Set Operations =======================================

// Refer to "Crafting A Compiler" by Charles N. Fisher & Richard J. LeBlanc, Jr.
// Section 6.2.1 LR(0) Parsing

func (g *Grammar) closure(i *item, A Symbol) *itemSet {
	iset := newItemSet()
	iset.Add(i)
	if A == nil {
		A = i.peekSymbol() // get symbol after dot
	}
	if A != nil {
		T.Debugf("pre closure(%v) = %v", i, iset)
		iset = g.closureSet(iset)
		T.Debugf("    closure(%v) = %v", i, iset)
		return iset
	}
	return iset
}

// https://www.cs.bgu.ac.il/~comp151/wiki.files/ps6.html#sec-2-7-3
func (g *Grammar) closureSet(iset *itemSet) *itemSet {
	cset := newItemSet()
	cset.union(iset)
	for _, x := range iset.Values() {
		i := x.(*item)
		if A := i.peekSymbol(); A != nil {
			// iterate through all rules
			// is LHS = A ?
			// create item A ::= * RHS  ? How to proceed with eps-rules?
			if !A.IsTerminal() {
				iiset := g.findNonTermRules(A)
				//T.Debugf("found %d items for closure", iiset.Size())
				cset.union(iiset)
			}
		}
	}
	return cset
}

func (g *Grammar) gotoSet(closure *itemSet, A Symbol) (*itemSet, Symbol) {
	// for every item in closure C
	// if item in C:  N -> ... * A ...
	//     advance N -> ... A * ...
	gotoset := newItemSet()
	for _, x := range closure.Values() {
		i := x.(*item)
		if i.peekSymbol() == A {
			ii, _ := i.advance()
			T.Debugf("goto(%s) -%s-> %s", i, A, ii)
			gotoset.Add(ii)
		}
	}
	//gotoset.Dump()
	return gotoset, A
}

func (g *Grammar) gotoSetClosure(i *itemSet, A Symbol) (*itemSet, Symbol) {
	gotoset, _ := g.gotoSet(i, A)
	//T.Infof("gotoset  = %v", gotoset)
	gclosure := g.closureSet(gotoset)
	//T.Infof("gclosure = %v", gclosure)
	T.Debugf("goto(%s) --%s--> %s", i, A, gclosure)
	return gclosure, A
}

// === CFSM Construction =====================================================

// CFSM state
type cfsmState struct {
	id    int      // serial ID of this state
	items *itemSet // configuration items within this state
}

// CFSM edge between 2 states, directed and labeled with a terminal
type cfsmEdge struct {
	from  *cfsmState
	to    *cfsmState
	label Symbol
}

// Debugging helper
func (s *cfsmState) Dump() {
	T.Debugf("--- state %03d -----------", s.id)
	s.items.Dump()
	T.Debug("-------------------------")
}

func (s *cfsmState) isErrorState() bool {
	return s.items.Size() == 0
}

// Create a state from an item set
func state(id int, iset *itemSet) *cfsmState {
	s := &cfsmState{id: id}
	if iset == nil {
		s.items = newItemSet()
	} else {
		s.items = iset
	}
	return s
}

func (s *cfsmState) allItems() []interface{} {
	vals := s.items.Values()
	return vals
}

func (s *cfsmState) String() string {
	return fmt.Sprintf("(state %d | [%d])", s.id, s.items.Size())
}

// Create an edge
func edge(from, to *cfsmState, label Symbol) *cfsmEdge {
	return &cfsmEdge{
		from:  from,
		to:    to,
		label: label,
	}
}

// We need this for the set of states. It sorts states by serial ID.
func stateComparator(s1, s2 interface{}) int {
	c1 := s1.(*cfsmState)
	c2 := s2.(*cfsmState)
	return utils.IntComparator(c1.id, c2.id)
}

// Add a state to the CFSM. Checks first if state is present.
func (c *CFSM) addState(iset *itemSet) *cfsmState {
	s := c.findStateByItems(iset)
	if s == nil {
		s = state(c.cfsmIds, iset)
		c.cfsmIds++
	}
	c.states.Add(s)
	return s
}

// Find a CFSM state by the contained item set.
func (c *CFSM) findStateByItems(iset *itemSet) *cfsmState {
	it := c.states.Iterator()
	for it.Next() {
		s := it.Value().(*cfsmState)
		if s.items.equals(iset) {
			return s
		}
	}
	return nil
}

func (c *CFSM) addEdge(s0, s1 *cfsmState, sym Symbol) *cfsmEdge {
	e := edge(s0, s1, sym)
	c.edges.Add(e)
	return e
}

func (c *CFSM) allEdges(s *cfsmState) []*cfsmEdge {
	it := c.edges.Iterator()
	r := make([]*cfsmEdge, 0, 2)
	for it.Next() {
		e := it.Value().(*cfsmEdge)
		if e.from == s {
			r = append(r, e)
		}
	}
	return r
}

// LR(0) state diagram for a grammar, i.e. the characteristic finite
// state automata CFSM. Will be constructed by a LRTableGenerator.
// Clients normally do not use it directly. Nevertheless, there are some methods
// defined on it, e.g, for debugging purposes, or even to
// compute your own tables from it.
type CFSM struct {
	g       *Grammar        // this CFSM is for Grammar g
	states  *treeset.Set    // all the states
	edges   *arraylist.List // all the edges between states
	cfsmIds int             // serial IDs for CFSM states
}

// create an empty (initial) CFSM automata.
func emptyCFSM(g *Grammar) *CFSM {
	c := &CFSM{g: g}
	c.states = treeset.NewWith(stateComparator)
	c.edges = arraylist.New()
	return c
}

/*
Generator object to construct LR parser tables.
Clients usually create a Grammar G, then a GrammarAnalysis-object for G,
and then a table generator. LRTableGenerator.CreateTables() constructs
the CFSM and parser tables for an LR-parser recognizing grammar G.
*/
type LRTableGenerator struct {
	g           *Grammar
	ga          *GrammarAnalysis
	dfa         *CFSM
	gototable   *sparse.IntMatrix
	actiontable *sparse.IntMatrix
}

/*
Create a new LRTableGenerator for a (previously analysed) grammar.
*/
func NewLRTableGenerator(ga *GrammarAnalysis) *LRTableGenerator {
	lrgen := &LRTableGenerator{}
	lrgen.g = ga.Grammar()
	lrgen.ga = ga
	return lrgen
}

/*
Return the characteristic finite state machine (CFSM) for a grammar.
Usually clients call lrgen.CreateTables() beforehand, but it is possible
to call lrgen.CFSM() directly. The CFSM will be created, if it has not
been constructed previously.
*/
func (lrgen *LRTableGenerator) CFSM() *CFSM {
	if lrgen.dfa == nil {
		lrgen.dfa = lrgen.buildCFSM()
	}
	return lrgen.dfa
}

/*
Create the necessary data structures for an SLR parser.
*/
func (lrgen *LRTableGenerator) CreateTables() {
	lrgen.dfa = lrgen.buildCFSM()
	lrgen.gototable = lrgen.BuildGotoTable()
	lrgen.actiontable = lrgen.BuildSLR1ActionTable()
}

// Construct the characteristic finite state machine CFSM for a grammar.
func (lrgen *LRTableGenerator) buildCFSM() *CFSM {
	T.Debug("=== build CFSM ==================================================")
	g := lrgen.g
	r0 := g.rules[0]
	closure0 := g.closure(r0.startItem())
	cfsm := emptyCFSM(g)
	s0 := cfsm.addState(closure0)
	s0.Dump()
	S := treeset.NewWith(stateComparator)
	S.Add(s0)
	for S.Size() > 0 {
		s := S.Values()[0].(*cfsmState)
		S.Remove(s)
		g.symbols.Each(func(name string, sym runtime.Symbol) {
			T.Debugf("sym %s = %v", name, sym)
			A := sym.(Symbol)
			gotoset, _ := g.gotoSetClosure(s.items, A)
			snew := cfsm.findStateByItems(gotoset)
			if snew == nil {
				snew = cfsm.addState(gotoset)
				if !snew.isErrorState() {
					S.Add(snew)
				}
			}
			if !snew.isErrorState() {
				cfsm.addEdge(s, snew, A)
			}
			snew.Dump()
		})
		T.Debug("-----------------------------------------------------------------")
	}
	return cfsm
}

// Export an CFSM to the Graphviz Dot format, given a filename.
func (cfsm *CFSM) CFSM2GraphViz(filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(fmt.Sprintf("file open error: %v", err.Error()))
	}
	defer f.Close()
	f.WriteString(`digraph {
node [shape=record];

`)
	for _, x := range cfsm.states.Values() {
		s := x.(*cfsmState)
		f.WriteString(fmt.Sprintf("s%03d [label=\"{%03d | %s}\"]\n", s.id, s.id, s.items.forGraphviz()))
	}
	it := cfsm.edges.Iterator()
	for it.Next() {
		x := it.Value()
		edge := x.(*cfsmEdge)
		f.WriteString(fmt.Sprintf("s%03d -> s%03d [label=\"%s\"]\n", edge.from.id, edge.to.id, edge.label))
	}
	f.WriteString("}\n")
}

// ===========================================================================

/*
Build the GOTO table. This is normally not called directly, but rather
via CreateTables().
*/
func (lrgen *LRTableGenerator) BuildGotoTable() *sparse.IntMatrix {
	statescnt := lrgen.dfa.states.Size()
	maxtok := 0
	lrgen.g.symbols.Each(func(n string, sym runtime.Symbol) {
		A := sym.(Symbol)
		if A.Token() > maxtok { // find maximum token value
			maxtok = A.Token()
		}
	})
	T.Infof("GOTO table of size %d x %d", statescnt, maxtok)
	gototable := sparse.NewIntMatrix(statescnt, maxtok, sparse.DefaultNullValue)
	states := lrgen.dfa.states.Iterator()
	for states.Next() {
		state := states.Value().(*cfsmState)
		edges := lrgen.dfa.allEdges(state)
		for _, e := range edges {
			//T.Debugf("edge %s --%v--> %v", state, e.label, e.to)
			//T.Debugf("GOTO (%d , %d ) = %d", state.id, symvalue(e.label), e.to.id)
			gototable.Set(state.id, symvalue(e.label), int32(e.to.id))
		}
	}
	return gototable
}

func GotoTableAsHTML(lrgen *LRTableGenerator, w io.Writer) {
	if lrgen.gototable == nil {
		T.Errorf("GOTO table not yet created, cannot export to HTML")
		return
	}
	parserTableAsHTML(lrgen, "GOTO", lrgen.gototable, w)
}

func ActionTableAsHTML(lrgen *LRTableGenerator, w io.Writer) {
	if lrgen.actiontable == nil {
		T.Errorf("ACTION table not yet created, cannot export to HTML")
		return
	}
	parserTableAsHTML(lrgen, "ACTION", lrgen.actiontable, w)
}

func parserTableAsHTML(lrgen *LRTableGenerator, tname string, table *sparse.IntMatrix, w io.Writer) {
	var symvec []Symbol = make([]Symbol, lrgen.g.symbols.Size())
	io.WriteString(w, "<html><body>\n")
	io.WriteString(w, "<img src=\"cfsm.png\"/><p>")
	io.WriteString(w, fmt.Sprintf("%s table of size = %d<p>", tname, table.ValueCount()))
	io.WriteString(w, "<table border=1 cellspacing=0 cellpadding=5>\n")
	io.WriteString(w, "<tr bgcolor=#cccccc><td></td>\n")
	j := 0
	lrgen.g.symbols.Each(func(n string, sym runtime.Symbol) {
		A := sym.(Symbol)
		io.WriteString(w, fmt.Sprintf("<td>%s</td>", A))
		symvec[j] = A
		j++
	})
	io.WriteString(w, "</tr>\n")
	states := lrgen.dfa.states.Iterator()
	for states.Next() {
		state := states.Value().(*cfsmState)
		io.WriteString(w, fmt.Sprintf("<tr><td>state %d</td>\n", state.id))
		for _, A := range symvec {
			v := table.Value(state.id, symvalue(A))
			td := fmt.Sprintf("%d", v)
			if v == table.NullValue() {
				td = "&nbsp;"
			}
			io.WriteString(w, "<td>")
			io.WriteString(w, td)
			io.WriteString(w, "</td>\n")
		}
		io.WriteString(w, "</tr>\n")
	}
	io.WriteString(w, "</table></body></html>\n")
}

// ===========================================================================

/*
Build the LR(0) Action table. This method is not called by CreateTables(),
as we normally use an SLR(1) parser and therefore an action table with
lookahead included. This method is provided as an add-on.
*/
func (lrgen *LRTableGenerator) BuildLR0ActionTable() *sparse.IntMatrix {
	statescnt := lrgen.dfa.states.Size()
	T.Infof("ACTION.0 table of size %d x 1", statescnt)
	actions := sparse.NewIntMatrix(statescnt, 1, sparse.DefaultNullValue)
	return lrgen.buildActionTable(actions, false)
}

/*
Build the SLR(1) Action table. This method is normally not called by
clients, but rather via CreateTables(). It builds an action table including
lookahead (using the FOLLOW-set created by the grammar analyzer).
*/
func (lrgen *LRTableGenerator) BuildSLR1ActionTable() *sparse.IntMatrix {
	statescnt := lrgen.dfa.states.Size()
	maxtok := 0
	lrgen.g.symbols.Each(func(n string, sym runtime.Symbol) {
		A := sym.(Symbol)
		if A.Token() > maxtok { // find maximum token value
			maxtok = A.Token()
		}
	})
	T.Infof("ACTION.1 table of size %d x %d", statescnt, maxtok)
	actions := sparse.NewIntMatrix(statescnt, maxtok, sparse.DefaultNullValue)
	return lrgen.buildActionTable(actions, true)
}

/*
	states := lrgen.dfa.states.Iterator()
	for states.Next() {
		state := states.Value().(*cfsmState)
		for _, v := range state.items.Values() {
			T.Infof("item in s%d = %v", state.id, v)
			i, _ := v.(*item)
			sym := i.peekSymbol()
			prefix := i.getPrefix()
			//sid := state.id
			T.Infof("    symbol at dot = %v, prefix = %v", sym, prefix)
			if sym != nil && sym.IsTerminal() { // create a shift entry
				T.Info("    creating shift action entry")
				actions.Add(state.id, 1, 1) // general shift (no lookahead)
			}
			if len(prefix) > 0 && sym == nil {
				rule, inx := lrgen.g.matchesRHS(prefix, false)
				if inx >= 0 { // create a reduce entry
					lookaheads := lrgen.ga.Follow(rule.lhs[0])
					T.Infof("    Follow(%v) = %v", rule.lhs[0], lookaheads)
					for _, la := range lookaheads {
						actions.Add(state.id, la, int32(-inx)) // reduce rule[inx]
						T.Infof("    creating reduce_%d action entry @ %v for %v", inx, la, rule)
					}
				}
			}
		}
	}
	return actions
}
*/

func (lrgen *LRTableGenerator) buildActionTable(actions *sparse.IntMatrix, slr1 bool) *sparse.IntMatrix {
	states := lrgen.dfa.states.Iterator()
	for states.Next() {
		state := states.Value().(*cfsmState)
		for _, v := range state.items.Values() {
			T.Debugf("item in s%d = %v", state.id, v)
			i, _ := v.(*item)
			sym := i.peekSymbol()
			prefix := i.getPrefix()
			//sid := state.id
			T.Debugf("    symbol at dot = %v, prefix = %v", sym, prefix)
			if sym != nil && sym.IsTerminal() { // create a shift entry
				T.Debug("    creating shift action entry")
				actions.Add(state.id, 1, 1) // general shift (no lookahead)
			}
			if len(prefix) > 0 && sym == nil {
				rule, inx := lrgen.g.matchesRHS(prefix, false)
				if inx >= 0 { // create a reduce entry
					if slr1 {
						lookaheads := lrgen.ga.Follow(rule.lhs[0])
						T.Debugf("    Follow(%v) = %v", rule.lhs[0], lookaheads)
						for _, la := range lookaheads {
							actions.Add(state.id, la, int32(-inx)) // reduce rule[inx]
							T.Debugf("    creating reduce_%d action entry @ %v for %v", inx, la, rule)
						}
					} else {
						T.Debugf("    creating reduce_%d action entry for %v", inx, rule)
						actions.Add(state.id, 1, int32(-inx)) // reduce rule[inx]
					}
				}
			}
		}
	}
	return actions
}