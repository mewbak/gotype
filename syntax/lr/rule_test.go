package lr

import (
	"testing"

	"github.com/npillmayer/gotype/gtcore/config/tracing"
	"github.com/npillmayer/gotype/syntax/runtime"
)

func traceOn() {
	T.SetLevel(tracing.LevelDebug)
}

func TestBuilder1(t *testing.T) {
	traceOn()
	b := NewGrammarBuilder("G")
	b.LHS("S").N("A").End()
	if len(b.Grammar().rules) != 1 {
		t.Fail()
	}
}

func TestBuilder2(t *testing.T) {
	b := NewGrammarBuilder("G")
	b.LHS("S").Epsilon()
	if len(b.Grammar().rules) != 1 {
		t.Fail()
	}
}

func TestItems1(t *testing.T) {
	b := NewGrammarBuilder("G")
	r := b.LHS("S").N("E").EOF()
	i1, _ := r.startItem()
	i2, _ := r.startItem() // items are cashed to not get duplicates
	if i1 != i2 {
		t.Fail()
	}
}

func TestClosure1(t *testing.T) {
	b := NewGrammarBuilder("G")
	r1 := b.LHS("S").N("E").EOF()
	r2 := b.LHS("E").N("E").T("+", 1).N("E").End()
	if len(b.Grammar().rules) != 2 {
		t.Fail()
	}
	item1, _ := r1.startItem()
	item2, _ := r2.startItem()
	T.Debug(item1)
	T.Debug(item2)
	if item1.dot != 0 {
		t.Fail()
	}
	item2, _ = item2.advance()
	T.Debug(item2)
	if item2.dot != 1 {
		t.Fail()
	}
}

func TestClosure2(t *testing.T) {
	b := NewGrammarBuilder("G")
	r0 := b.LHS("S").N("E").EOF()
	b.LHS("E").N("E").T("+", 1).T("(", 2).N("E").T(")", 3).End()
	b.LHS("E").T("a", 4).End()
	g := b.Grammar()
	b.Grammar().Dump()
	closure0 := g.closure(r0.startItem())
	closure0.Dump()
}

func TestItemSetEquality(t *testing.T) {
	b := NewGrammarBuilder("G")
	r0 := b.LHS("S").N("E").EOF()
	b.LHS("E").N("E").T("+", 1).T("(", 2).N("E").T(")", 3).End()
	b.LHS("E").T("a", 4).End()
	g := b.Grammar()
	b.Grammar().Dump()
	closure0 := g.closure(r0.startItem())
	closure0.Dump()
	closure1 := g.closure(r0.startItem())
	if !closure0.equals(closure1) {
		t.Fail()
	}
}

func TestClosure4(t *testing.T) {
	b := NewGrammarBuilder("G")
	r0 := b.LHS("S").N("E").EOF()
	b.LHS("E").N("E").T("+", 1).T("(", 2).N("E").T(")", 3).End()
	b.LHS("E").T("a", 4).End()
	g := b.Grammar()
	b.Grammar().Dump()
	i, A := r0.startItem()
	closure0 := g.closure(i, A)
	closure0.Dump()
	g.gotoSet(closure0, A)
}

func TestStateRetrieval(t *testing.T) {
	b := NewGrammarBuilder("G")
	r0 := b.LHS("S").N("E").EOF()
	b.LHS("E").N("E").T("+", 1).T("(", 2).N("E").T(")", 3).End()
	b.LHS("E").T("a", 4).End()
	g := b.Grammar()
	cfsm := emptyCFSM()
	closure0 := g.closure(r0.startItem())
	s0 := cfsm.addState(closure0)
	s0.Dump()
	s1 := cfsm.addState(closure0)
	if s0.id != s1.id {
		t.Fail()
	}
}

func TestBuildCFSM(t *testing.T) {
	b := NewGrammarBuilder("G")
	b.LHS("S").N("E").EOF()
	b.LHS("E").N("E").T("+", 1).T("(", 2).N("E").T(")", 3).End()
	b.LHS("E").T("a", 4).End()
	g := b.Grammar()
	c := g.buildCFSM()
	cfsm2dot(c)
}

func TestDerivesEps(t *testing.T) {
	b := NewGrammarBuilder("G")
	b.LHS("S").N("E").EOF()
	b.LHS("E").N("T").T("a", 1).End()
	b.LHS("E").N("F").End()
	b.LHS("F").Epsilon()
	g := b.Grammar()
	//g.Dump()
	ga := NewGrammarAnalysis(g)
	ga.markEps()
	cnt := 0
	g.symbols.Each(func(name string, sym runtime.Symbol) {
		A := sym.(Symbol)
		T.Debugf("%s => eps  : %v", name, ga.derivesEps[A])
		if ga.derivesEps[A] {
			cnt++
		}
	})
	if cnt != 2 {
		t.Fail() // E and F should => eps
	}
}

func TestFirstSet(t *testing.T) {
	b := NewGrammarBuilder("G")
	b.LHS("S").N("E").EOF()
	b.LHS("E").N("T").End()
	b.LHS("T").T("a", 1).End()
	b.LHS("T").Epsilon()
	g := b.Grammar()
	g.Dump()
	ga := NewGrammarAnalysis(g)
	ga.markEps()
	ga.initFirstSets()
	for key, value := range ga.firstSets.sets {
		T.Debugf("key = %v     value = %v", key, value)
	}
}

func TestFollowSet(t *testing.T) {
	b := NewGrammarBuilder("G")
	b.LHS("S").N("E").End()
	b.LHS("E").N("T").EOF()
	b.LHS("T").T("a", 1).End()
	b.LHS("T").Epsilon()
	g := b.Grammar()
	g.Dump()
	ga := NewGrammarAnalysis(g)
	ga.markEps()
	ga.initFirstSets()
	for key, value := range ga.firstSets.sets {
		T.Debugf("key = %v     value = %v", key, value)
	}
	T.Debug("-------")
	ga.initFollowSets()
	for key, value := range ga.followSets.sets {
		T.Debugf("key = %v     value = %v", key, value)
	}
}
