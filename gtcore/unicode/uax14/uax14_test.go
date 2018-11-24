package uax14

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/npillmayer/gotype/gtcore/unicode"
)

func TestAddPenalties(t *testing.T) {
	total := make([]int, 0, 5)
	penalties := []int{17, 23}
	total = unicode.AddPenalties(total, penalties)
	fmt.Printf("total = %v\n", total)
}

func TestClassForRune1(t *testing.T) {
	SetupUAX14Classes()
	var r rune
	r = 'A'
	c := UAX14ClassForRune(r)
	fmt.Printf("%+q = %s\n", r, c)
}

func TestClassForRune2(t *testing.T) {
	SetupUAX14Classes()
	var r rune
	//r = 'A'
	r = '世'
	c := UAX14ClassForRune(r)
	fmt.Printf("%+q = %s\n", r, c)
}

func TestLineWrapNL(t *testing.T) {
	SetupUAX14Classes()
	publisher := unicode.NewRunePublisher()
	lw := NewLineWrap()
	lw.InitFor(publisher)
	lw.StartRulesFor('\n', int(NLClass))
	lw.ProceedWithRune('\n', int(NLClass))
	lw.ProceedWithRune('A', int(ALClass))
	if lw.LongestMatch() != 1 {
		t.Fail()
	}
}

func TestLineWrapQU(t *testing.T) {
	SetupUAX14Classes()
	publisher := unicode.NewRunePublisher()
	lw := NewLineWrap()
	lw.InitFor(publisher)
	lw.StartRulesFor('"', int(QUClass))
	lw.ProceedWithRune('"', int(QUClass))
	lw.ProceedWithRune(' ', int(SPClass))
	lw.ProceedWithRune('(', int(OPClass))
	lw.ProceedWithRune(' ', int(SPClass))
}

func TestSegmenterUAX14Init(t *testing.T) {
	SetupUAX14Classes()
	lw := NewLineWrap()
	segm := unicode.NewSegmenter(lw)
	_, _, err := segm.Next()
	fmt.Println(err)
	if err == nil {
		t.Fail()
	}
}

func TestSegmenterUAX14RecognizeRule1(t *testing.T) {
	SetupUAX14Classes()
	lw := NewLineWrap()
	segm := unicode.NewSegmenter(lw)
	segm.Init(strings.NewReader("\" ("))
	_, _, err := segm.Next()
	if err != io.EOF {
		fmt.Println(err)
		t.Fail()
	}
}

func TestSegmenterUAX14Match1(t *testing.T) {
	SetupUAX14Classes()
	lw := NewLineWrap()
	segm := unicode.NewSegmenter(lw)
	segm.Init(strings.NewReader("\" ("))
	match, _, err := segm.Next()
	if err != io.EOF {
		fmt.Println(err)
		t.Fail()
	}
	if match == nil {
		t.Fail()
	}
	fmt.Printf("matched segment = \"%s\"\n", match)
}
