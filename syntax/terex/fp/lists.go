package fp

import (
	"fmt"

	"github.com/npillmayer/gotype/syntax/terex"
)

/*
The current implementation always pre-fetches the first value.
This could be optimized. It would be a problem with long-running ops in the
atom-creation, in case the value is never fetched by an output call.
For now, we will leave it this way.
*/

type ListSeq struct {
	atom terex.Atom
	seq  ListGenerator
}

func Seq(l *terex.GCons) ListSeq {
	var S ListGenerator
	S = func() ListSeq {
		if l == nil {
			return ListSeq{terex.NilAtom, nil}
		}
		atom := l.Car
		l = l.Cdr
		return ListSeq{atom, S}
	}
	atom := l.Car
	return ListSeq{atom, S}
}

func (seq *ListSeq) Break() {
	seq.seq = nil
}

func (seq *ListSeq) Done() bool {
	return seq.seq == nil
}

func (seq ListSeq) First() (terex.Atom, ListSeq) {
	return seq.atom, seq
}

func (seq *ListSeq) Next() terex.Atom {
	if seq.Done() {
		return terex.NilAtom
	}
	next := seq.seq()
	seq.atom = next.atom
	if seq.atom == terex.NilAtom {
		seq.seq = nil
	} else {
		seq.seq = next.seq
	}
	return seq.atom
}

type ListGenerator func() ListSeq

func NSeq() ListSeq {
	var n int64
	var S ListGenerator
	S = func() ListSeq {
		n++
		atom := terex.Atomize(n)
		return ListSeq{atom, S}
	}
	atom := terex.Atomize(n)
	return ListSeq{atom, S}
}

type ListMapper func(terex.Atom) terex.Atom

func (seq ListSeq) Map(mapper ListMapper) ListSeq {
	var F ListGenerator
	//inner := seq
	atom, inner := seq.atom, seq
	//n, inner := seq.First()
	v := mapper(atom)
	F = func() ListSeq {
		//fmt.Printf("F  called, n=%d\n", n)
		atom = inner.Next()
		v = mapper(atom)
		//fmt.Printf("F' n=%d, v=%d\n", n, v)
		return ListSeq{v, F}
	}
	return ListSeq{v, F}
}

func (seq ListSeq) List() *terex.GCons {
	if seq.Done() {
		return nil
	}
	var start, end *terex.GCons
	//atom, S := seq.First()
	//fmt.Printf("first atom=%s\n", atom)
	S := seq
	// var atom terex.Atom
	for atom := seq.Next(); !S.Done(); atom = S.Next() {
		fmt.Printf("next atom=%s, S=%v\n", atom, S)
		fmt.Printf("  done=%v\n", S.Done())
		if start == nil {
			start = terex.Cons(atom, nil)
			end = start
		} else {
			end.Cdr = terex.Cons(atom, nil)
			end = end.Cdr
		}
		fmt.Printf("result list = %s\n", start.ListString())
	}
	return start
}

// --- Trees -----------------------------------------------------------------

type TreeSeq struct {
	stack []*terex.GCons
}

func Tree(l *terex.GCons) *TreeSeq {
	tseg := &TreeSeq{stack: make([]*terex.GCons, 0, 32)}
	if l == nil {
		return tseg
	}
	tseg.stack = append(tseg.stack, l) // push root
	node := tseg.stack[len(tseg.stack)-1]
	if node.IsLeaf() {
		node = nil
		tseg.stack = tseg.stack[:len(tseg.stack)-1] // pop node
		if len(tseg.stack) == 0 {
			tseg.seq = nil // have returned to root
		}
	} else if node.Car.Type() == terex.ConsType {
		if node.Car.Data != nil {
			tseg.stack = append(tseg.stack, node.Tee()) // push left child node
		} else {
			node = nil
			tseg.stack = tseg.stack[:len(tseg.stack)-1] // pop node
			if len(tseg.stack) == 0 {
				tseg.seq = nil // have returned to root
			}
		}
	} else { // Cdr != nil
		tseg.stack = append(tseg.stack, node.Cdr) // push right child node
	}
	return tseg
}

func (seq *TreeSeq) Break() {
	seq.seq = nil
}

func (seq *TreeSeq) Done() bool {
	return seq.seq == nil
}

func (seq *TreeSeq) First() (*terex.GCons, *TreeSeq) {
	return seq.stack[len(seq.stack)-1], seq
}

func (seq *TreeSeq) Next() *terex.GCons {
	if seq.Done() {
		return nil
	}
	// next := seq.seq()
	// seq.list = next.list
	// seq.seq = next.seq
	return seq.stack[len(seq.stack)-1]
}

type TreeGenerator func() *TreeSeq
