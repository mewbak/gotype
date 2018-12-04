/*
Package for a generator for UAX#29 word breaking classes.

BSD License

Copyright (c) 2017–18, Norbert Pillmayer (norbert@pillmayer.com)

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

Contents

Generator for Unicode UAX#29 word breaking code-point classes.
For more information see http://unicode.org/reports/tr29/

Classes are generated from a UAX#29 companion file: "WordBreakProberty.txt".
This is the definite source for UAX#29 code-point classes. The
generator looks for it in a directory "$GOPATH/etc/".


Usage

The generator has just one option, a "verbose" flag. It should usually
be turned on.

   generator [-v]

This creates a file "uax29classes.go" in the current directory. It is designed
to be called from the "uax29" directory.
*/
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"runtime"
	"text/template"
	"time"

	"os"

	"github.com/emirpasic/gods/lists/arraylist"
	"github.com/npillmayer/gotype/gtcore/unicode/ucd"
)

var logger = log.New(os.Stderr, "UAX#29 generator: ", log.LstdFlags)

// flag: verbose output ?
var verbose bool

var uax29classnames = []string{"ALetter", "CR", "Double_Quote", "Extend",
	"ExtendNumLet", "Format", "Hebrew_Letter", "Katakana", "LF", "MidLetter",
	"MidNum", "MidNumLet", "Newline", "Numeric", "Regional_Indicator",
	"Single_Quote", "WSegSpace", "ZWJ"}

// Load the Unicode UAX#29 definition file: WordBreakProperty.txt
func loadUnicodeLineBreakFile() (map[string][]rune, error) {
	if verbose {
		logger.Printf("reading WordBreakProperty.txt")
	}
	defer timeTrack(time.Now(), "loading WordBreakProperty.txt")
	gopath := os.Getenv("GOPATH")
	f, err := os.Open(gopath + "/etc/WordBreakProperty.txt")
	if err != nil {
		fmt.Printf("ERROR loading " + gopath + "/etc/WordBreakProperty.txt\n")
		return nil, err
	}
	defer f.Close()
	p := ucd.NewUCDParser(f)
	lbcs := make(map[string]*arraylist.List, len(uax29classnames))
	for p.Next() {
		from, to := p.Range(0)
		brclzstr := p.String(1)
		list := lbcs[brclzstr]
		if list == nil {
			list = arraylist.New()
		}
		for r := from; r <= to; r++ {
			list.Add(r)
		}
		lbcs[brclzstr] = list
	}
	err = p.Err()
	if err != nil {
		log.Fatal(err)
	}
	runeranges := make(map[string][]rune)
	for k, v := range lbcs {
		runelist := make([]rune, lbcs[k].Size())
		it := v.Iterator()
		i := 0
		for it.Next() {
			runelist[i] = it.Value().(rune)
			i++
		}
		runeranges[k] = runelist
	}
	return runeranges, err
}

// --- Templates --------------------------------------------------------

var header string = `package uax29

// This file has been generated -- you probably should NOT EDIT IT !
// 
// BSD License, Copyright (c) 2018, Norbert Pillmayer (norbert@pillmayer.com)

import (
    "strconv"
    "unicode"

    "golang.org/x/text/unicode/rangetable"
)
`

var templateClassType string = `
// Type for UAX#29 code-point classes.
// Must be convertable to int.
type UAX29Class int

// Will be initialized in SetupUAX29Classes()
var rangeFromUAX29Class []*unicode.RangeTable
`

var templateRangeTableVars string = `
// Range tables for UAX#29 code-point classes.
// Will be initialized with SetupUAX29Classes().
// Clients can check with unicode.Is(..., rune){{$i:=0}}
var {{range .}}{{$i = inc $i}}{{.}}, {{if modten $i}}
    {{end}}{{end}}unused *unicode.RangeTable
`

var templateClassConsts string = `
// These are all the UAX#29 breaking classes.
const ( {{$i:=0}}
{{range  .}}    {{.}}Class UAX29Class = {{$i}}{{$i = inc $i}}
{{end}}
    Other UAX29Class = 999
    sot   UAX29Class = 1000 // pseudo class "start of text"
    eot   UAX29Class = 1001 // pseudo class "end of text"
)
`

var templateClassStringer string = `
const _UAX29Class_name = "{{range $c,$name := .}}{{$name}}Class{{end}}"

var _UAX29Class_index = [...]uint16{0{{startinxs .}} }

// Stringer for type UAX29Class
func (c UAX29Class) String() string {
    if c == sot {
        return "sot"
    } else if c == eot {
        return "eot"
    } else if c == Other {
        return "Other"
    } else if c < 0 || c >= UAX29Class(len(_UAX29Class_index)-1) {
        return "UAX29Class(" + strconv.FormatInt(int64(c), 10) + ")"
    }
    return _UAX29Class_name[_UAX29Class_index[c]:_UAX29Class_index[c+1]]
}
`

var templateRangeForClass string = `{{$i:=0}}{{range .}}{{if notfirst $i}}, {{if modeight $i}}
    {{end}}{{end}}{{$i = inc $i}}{{printf "%+q" .}}{{end}}`

// Helper functions for templates
var funcMap template.FuncMap = template.FuncMap{
	"modten": func(i int) bool {
		return i%10 == 0
	},
	"modeight": func(i int) bool {
		return (i+2)%8 == 0
	},
	"inc": func(i int) int {
		return i + 1
	},
	"notfirst": func(i int) bool {
		return i > 0
	},
	"startinxs": func(str []string) string {
		out := ""
		total := 0
		for _, s := range str {
			l := len(s) + 5
			total += l
			if (38+len(out))%80 > 72 {
				out += fmt.Sprintf(",\n    %d", total)
			} else {
				out += fmt.Sprintf(", %d", total)
			}
		}
		return out
	},
}

func makeTemplate(name string, templString string) *template.Template {
	if verbose {
		logger.Printf("creating %s", name)
	}
	t := template.Must(template.New(name).Funcs(funcMap).Parse(templString))
	return t
}

// --- Main -------------------------------------------------------------

func generateRanges(w *bufio.Writer, codePointLists map[string][]rune) {
	defer timeTrack(time.Now(), "generate range tables")
	w.WriteString("\nfunc setupUAX29Classes() {\n")
	w.WriteString("    rangeFromUAX29Class = make([]*unicode.RangeTable, int(ZWJClass)+1)\n")
	t := makeTemplate("UAX#29 range", templateRangeForClass)
	for key, codepoints := range codePointLists {
		w.WriteString(fmt.Sprintf("\n    // Range for UAX#29 class %s\n", key))
		w.WriteString(fmt.Sprintf("    %s = rangetable.New(", key))
		checkFatal(t.Execute(w, codepoints))
		w.WriteString(")\n")
		w.WriteString(fmt.Sprintf("    rangeFromUAX29Class[int(%sClass)] = %s\n", key, key))
	}
	w.WriteString("}\n")
}

func main() {
	doVerbose := flag.Bool("v", false, "verbose output mode")
	flag.Parse()
	verbose = *doVerbose
	codePointLists, err := loadUnicodeLineBreakFile()
	checkFatal(err)
	if verbose {
		logger.Printf("loaded %d UAX#29 breaking classes\n", len(codePointLists))
	}
	f, ioerr := os.Create("uax29classes.go")
	checkFatal(ioerr)
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString(header)
	w.WriteString(templateClassType)
	t := makeTemplate("UAX#29 classes", templateClassConsts)
	checkFatal(t.Execute(w, uax29classnames))
	t = makeTemplate("UAX#29 range tables", templateRangeTableVars)
	checkFatal(t.Execute(w, uax29classnames))
	t = makeTemplate("UAX#29 classes stringer", templateClassStringer)
	checkFatal(t.Execute(w, uax29classnames))
	generateRanges(w, codePointLists)
	w.Flush()
}

// --- Util -------------------------------------------------------------

// Little helper for testing
func timeTrack(start time.Time, name string) {
	if verbose {
		elapsed := time.Since(start)
		logger.Printf("timing: %s took %s\n", name, elapsed)
	}
}

func checkFatal(err error) {
	_, file, line, _ := runtime.Caller(1)
	if err != nil {
		logger.Fatalln(":", file, ":", line, "-", err)
	}
}
