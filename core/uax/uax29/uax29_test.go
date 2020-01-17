package uax29_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/npillmayer/gotype/core/config/gtrace"
	"github.com/npillmayer/gotype/core/config/tracing"
	"github.com/npillmayer/gotype/core/config/tracing/gotestingadapter"
	"github.com/npillmayer/gotype/core/uax/segment"
	"github.com/npillmayer/gotype/core/uax/uax29"
	"github.com/npillmayer/gotype/core/uax/ucd"
)

func ExampleWordBreaker() {
	onWords := uax29.NewWordBreaker()
	segmenter := segment.NewSegmenter(onWords)
	segmenter.Init(strings.NewReader("Hello World🇩🇪!"))
	for segmenter.Next() {
		fmt.Printf("'%s'\n", segmenter.Text())
	}
	// Output: 'Hello'
	// ' '
	// 'World'
	// '🇩🇪'
	// '!'
}

func TestWordBreaks1(t *testing.T) {
	onWords := uax29.NewWordBreaker()
	segmenter := segment.NewSegmenter(onWords)
	segmenter.Init(strings.NewReader("Hello World "))
	n := 0
	for segmenter.Next() {
		t.Logf("'%s'\n", segmenter.Text())
		n++
	}
	if n != 4 {
		t.Errorf("Expected # of segments to be 4, is %d", n)
	}
}

func TestWordBreaks2(t *testing.T) {
	onWords := uax29.NewWordBreaker()
	segmenter := segment.NewSegmenter(onWords)
	segmenter.Init(strings.NewReader("lime-tree"))
	n := 0
	for segmenter.Next() {
		t.Logf("'%s'  (p=%v)", segmenter.Text(), segmenter.Penalties())
		n++
	}
	if n != 3 {
		t.Errorf("Expected # of segments to be 3, is %d", n)
	}
}

func TestWordBreakTestFile(t *testing.T) {
	teardown := gotestingadapter.RedirectTracing(t)
	defer teardown()
	gtrace.CoreTracer.SetTraceLevel(tracing.LevelError)
	onWordBreak := uax29.NewWordBreaker()
	seg := segment.NewSegmenter(onWordBreak)
	tf := ucd.OpenTestFile("./WordBreakTest.txt", t)
	defer tf.Close()
	failcnt, i, from, to := 0, 0, 1, 1900
	for tf.Scan() {
		i++
		if i >= from {
			//t.Logf(tf.Comment())
			in, out := ucd.BreakTestInput(tf.Text())
			if !executeSingleTest(t, seg, i, in, out) {
				failcnt++
			}
		}
		if i >= to {
			break
		}
	}
	if err := tf.Err(); err != nil {
		t.Errorf("reading input: %s", err)
	}
	t.Logf("%d TEST CASES OUT of %d FAILED", failcnt, i-from+1)
}

func executeSingleTest(t *testing.T, seg *segment.Segmenter, tno int, in string, out []string) bool {
	seg.Init(strings.NewReader(in))
	i := 0
	ok := true
	for seg.Next() {
		if len(out) <= i {
			t.Errorf("test #%d: number of segments too large: %d > %d", tno, i+1, len(out))
		} else if out[i] != seg.Text() {
			t.Errorf("test #%d: '%+q' should be '%+q'", tno, seg.Bytes(), out[i])
			ok = false
		}
		i++
	}
	return ok
}
