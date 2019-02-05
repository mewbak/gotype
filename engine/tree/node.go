package tree

/*
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
*/

import (
	"fmt"
	"sync"
)

// Node is the base type our tree is built of.
type Node struct {
	parent   *Node
	children childrenSlice
	Payload  interface{}
}

// NewNode creates a new tree node with a payload.
func NewNode(payload interface{}) *Node {
	return &Node{Payload: payload}
}

// String is a simple Stringer which the node's Payload packaged in a string.
func (node *Node) String() string {
	return fmt.Sprintf("(Node #ch=%d)", node.ChildCount())
}

// AddChild inserts a new child node into the tree.
// The newly inserted node is connected to this node as its parent.
// It returns the parent node to allow for chaining.
//
// This operation is concurrency-safe.
func (node *Node) AddChild(ch *Node) *Node {
	if ch != nil {
		node.children.addChild(ch, node)
	}
	return node
}

// SetChildAt inserts a new child node into the tree.
// The newly inserted node is connected to this node as its parent.
// The child is set at a given position in relation to other children,
// replacing the child at position i if it exists.
// It returns the parent node to allow for chaining.
//
// This operation is concurrency-safe.
func (node *Node) SetChildAt(i int, ch *Node) *Node {
	if ch != nil {
		node.children.setChild(i, ch, node)
	}
	return node
}

// ParentNode returns the parent node or nil (for the root of the tree).
func (node Node) Parent() *Node {
	return node.parent
}

// ChildCount returns the number of children-nodes for a node
// (concurrency-safe).
func (node Node) ChildCount() int {
	return node.children.length()
}

// Child is a concurrency-safe way to get a children-node of a node.
func (node Node) Child(n int) (*Node, bool) {
	if node.children.length() <= n {
		return nil, false
	}
	return node.children.child(n), true
}

// --- Slices of concurrency-safe sets of children ----------------------

type childrenSlice struct {
	sync.RWMutex
	slice []*Node
}

func (chs *childrenSlice) length() int {
	chs.RLock()
	defer chs.RUnlock()
	return len(chs.slice)
}

func (chs *childrenSlice) addChild(child *Node, parent *Node) {
	if child == nil {
		return
	}
	chs.Lock()
	defer chs.Unlock()
	chs.slice = append(chs.slice, child)
	child.parent = parent
}

func (chs *childrenSlice) setChild(i int, child *Node, parent *Node) {
	if child == nil {
		return
	}
	chs.Lock()
	defer chs.Unlock()
	if len(chs.slice) <= i {
		l := len(chs.slice)
		c := make([]*Node, i+1)
		for j := 0; j <= i; j++ {
			if j < l {
				c[j] = chs.slice[j]
			} else {
				c[j] = nil
			}
		}
		chs.slice = c
	}
	chs.slice[i] = child
	child.parent = parent
}

func (chs *childrenSlice) child(n int) *Node {
	if chs.length() == 0 || n < 0 || n >= chs.length() {
		return nil
	}
	chs.RLock()
	defer chs.RUnlock()
	return chs.slice[n]
}