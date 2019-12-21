package layout

import (
	"fmt"
	"strings"

	"github.com/npillmayer/gotype/engine/dom/cssom/style"
	"github.com/npillmayer/gotype/engine/frame/box"
	"github.com/npillmayer/gotype/engine/tree"
	"golang.org/x/net/html"
)

// Flags for box context and display level.
const (
	NoMode      uint8 = iota // in most cases an error condition
	BlockMode                // CSS block context / display level
	InlineMode               // CSS inline context / display level
	DisplayNone              // CSS display = none
)

// Container is a (CSS-)styled box which may contain other boxes and/or
// containers.
type Container struct {
	tree.Node                        // a container is a node within the layout tree
	Box                box.StyledBox // styled box wich is layed out
	contextOrientation uint8         // context of children (block or inline)
	displayLevel       uint8         // container lives in this mode (block or inline)
	StyleNode          *tree.Node    // the DOM node this Container refers to
}

// newContainer creates either a block-level container or an inline-level container
func newContainer(orientation uint8, displayLevel uint8) *Container {
	c := &Container{
		contextOrientation: orientation,
		displayLevel:       displayLevel,
	}
	c.Payload = c // always points to itself
	return c
}

// IsPrincipal is a predicate if this is a principal box.
//
// Some HTML elements create a mini-hierachy of boxes for rendering. The outermost box
// is called the principal box. It will always refer to the styled node.
// An example would be an "li"-element: it will create two sub-boxes, one for the
// list item marker and one for the item's text/content. Another example are anonymous
// boxes, which will be generated for reconciling context/level-discrepancies.
func (c *Container) IsPrincipal() bool {
	return (c.StyleNode != nil)
}

func (c *Container) IsBlock() bool {
	return c.contextOrientation == BlockMode
}

// newVBox creates a block-level container with block context.
func newBlockBox(sn *tree.Node) *Container {
	c := newContainer(BlockMode, BlockMode)
	c.StyleNode = sn
	return c
}

// newHBox creates an inline-level container with inline context.
func newInlineBox(sn *tree.Node) *Container {
	c := newContainer(InlineMode, InlineMode)
	c.StyleNode = sn
	return c
}

// Node retrieves the payload of a tree node as a Container.
// Will be called from clients as
//
//    container := layout.Node(n)
//
func Node(n *tree.Node) *Container {
	if n == nil {
		return nil
	}
	return n.Payload.(*Container)
}

func (c *Container) String() string {
	if c == nil {
		return "<empty>"
	}
	n := "_"
	b := "inline-box"
	if c.contextOrientation == BlockMode {
		b = "block-box"
	}
	return fmt.Sprintf("\\%s<%s>", b, n)
}

// Add appends a child container as the last sibling of existing
// child containers for c. Does nothing if child is nil.
// Will wrap the child in an anonymous container, if necessary.
//
// Returns the container itself (for chaining).
func (c *Container) Add(child *Container) *Container {
	T().Debugf("Adding child to %s", c)
	if child == nil {
		return c
	}
	if requiresAnonBox(c.contextOrientation, child.displayLevel) {
		anon := newContainer(child.displayLevel, c.contextOrientation)
		anon.AddChild(&child.Node)
		c.AddChild(&anon.Node)
		return c
	}
	c.AddChild(&child.Node)
	return c
}

func getDisplayLevelForStyledNode(sn *tree.Node, toStyler style.StyleInterf) (uint8, string) {
	if sn == nil {
		return NoMode, ""
	}
	styler := toStyler(sn)
	T().Infof("display(%s/%s) = ?", styler.HtmlNode().Data, nodeTypeString(styler.HtmlNode().Type))
	pmap := styler.ComputedStyles()
	dispProp, isSet := style.GetLocalProperty(pmap, "display")
	if !isSet {
		if styler.HtmlNode().Type != html.ElementNode &&
			styler.HtmlNode().Type != html.DocumentNode {
			T().Errorf("Have styled node for non-element ?!?")
			T().Errorf("type of node = %s", nodeTypeString(styler.HtmlNode().Type))
			return InlineMode, ""
		}
	}
	dispProp = style.DisplayPropertyForHtmlNode(styler.HtmlNode())
	if strings.HasPrefix(dispProp.String(), "none") {
		return DisplayNone, dispProp.String()
	} else if strings.HasPrefix(dispProp.String(), "block") {
		return BlockMode, dispProp.String()
	} else if strings.HasPrefix(dispProp.String(), "inline") {
		return InlineMode, dispProp.String() // inline or inline-block
	}
	return NoMode, ""
}

func nodeTypeString(nt html.NodeType) string {
	switch nt {
	case html.ErrorNode:
		return "error-node"
	case html.TextNode:
		return "text-node"
	case html.CommentNode:
		return "comment-node"
	case html.DocumentNode:
		return "doc-node"
	case html.ElementNode:
		return "element-node"
	case html.DoctypeNode:
		return "doctype-node"
	}
	return "?-node"
}