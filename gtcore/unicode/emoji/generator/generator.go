/*
Package for a generator for UTS#51 Emoji character classes.

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

Content

Generator for Unicode Emoji code-point classes. For more information
see http://www.unicode.org/reports/tr51/#Emoji_Properties_and_Data_Files

Classes are generated from a companion file: "emoji-data.txt".
The generator looks for it in a directory "$GOPATH/etc/".


Usage

The generator has just one option, a "verbose" flag. It should usually
be turned on.

   generator [-v]

This creates a file "emojiclasses.go" in the current directory. It is designed
to be called from the "emoji" directory.
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

var logger = log.New(os.Stderr, "UTS#51 generator: ", log.LstdFlags)

// flag: verbose output ?
var verbose bool

var emojiClassnames = []string{
	"Emoji",
	"Emoji_Presentation",
	"Emoji_Modifier",
	"Emoji_Modifier_Base",
	"Emoji_Component",
	"Extended_Pictographic",
}

// Load the Unicode UAX#29 definition file: EmojiBreakProperty.txt
func loadUnicodeEmojiBreakFile() (map[string][]rune, error) {
	if verbose {
		logger.Printf("reading EmojiBreakProperty.txt")
	}
	defer timeTrack(time.Now(), "loading emoji-data.txt")
	gopath := os.Getenv("GOPATH")
	f, err := os.Open(gopath + "/etc/emoji-data.txt")
	if err != nil {
		fmt.Printf("ERROR loading " + gopath + "/etc/emoji-data.txt\n")
		return nil, err
	}
	defer f.Close()
	parser := ucd.NewUCDParser(f)
	gcls := make(map[string]*arraylist.List, len(emojiClassnames))
	for parser.Next() {
		from, to := parser.Range(0)
		clstr := parser.String(1)
		list := gcls[clstr]
		if list == nil {
			list = arraylist.New()
		}
		for r := from; r <= to; r++ {
			list.Add(r)
		}
		gcls[clstr] = list
	}
	err = parser.Err()
	if err != nil {
		log.Fatal(err)
	}
	runeranges := make(map[string][]rune)
	for k, v := range gcls {
		runelist := make([]rune, gcls[k].Size())
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

var header string = `package emoji

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
// Type for UTS#51 emoji code-point classes.
// Must be convertable to int.
type EmojisClass int

// Will be initialized in SetupEmojisClasses()
var rangeFromEmojisClass []*unicode.RangeTable
`

var templateRangeTableVars string = `
// Range tables for emoji code-point classes.
// Will be initialized with SetupEmojisClasses().
// Clients can check with unicode.Is(..., rune){{$i:=0}}
var {{range .}}{{$i = inc $i}}{{.}}, {{if modeight $i}}
    {{end}}{{end}}unused *unicode.RangeTable
`

var templateClassConsts string = `
// These are all the emoji breaking classes.
const ( {{$i:=0}}
{{range  .}}    {{.}}Class EmojisClass = {{$i}}{{$i = inc $i}}
{{end}})
`

//{{range  $k,$v := .}}    {{$k}}Class EmojisClass = {{$v}}

var templateClassStringer string = `
const _EmojisClass_name = "{{range $c,$name := .}}{{$name}}Class{{end}}"

var _EmojisClass_index = [...]uint16{0{{startinxs .}} }

// Stringer for type EmojisClass
func (c EmojisClass) String() string {
    if c < 0 || c >= EmojisClass(len(_EmojisClass_index)-1) {
        return "EmojisClass(" + strconv.FormatInt(int64(c), 10) + ")"
    }
    return _EmojisClass_name[_EmojisClass_index[c]:_EmojisClass_index[c+1]]
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
			if (41+len(out))%80 > 75 {
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
	w.WriteString("\nfunc setupEmojisClasses() {\n")
	w.WriteString("    rangeFromEmojisClass = make([]*unicode.RangeTable, int(Extended_PictographicClass)+1)\n")
	t := makeTemplate("Emoji range", templateRangeForClass)
	for key, codepoints := range codePointLists {
		w.WriteString(fmt.Sprintf("\n    // Range for Emoji class %s\n", key))
		w.WriteString(fmt.Sprintf("    %s = rangetable.New(", key))
		checkFatal(t.Execute(w, codepoints))
		w.WriteString(")\n")
		w.WriteString(fmt.Sprintf("    rangeFromEmojisClass[int(%sClass)] = %s\n", key, key))
	}
	w.WriteString("}\n")
}

func main() {
	doVerbose := flag.Bool("v", false, "verbose output mode")
	flag.Parse()
	verbose = *doVerbose
	codePointLists, err := loadUnicodeEmojiBreakFile()
	checkFatal(err)
	if verbose {
		logger.Printf("loaded %d Emoji breaking classes\n", len(codePointLists))
	}
	f, ioerr := os.Create("emojiclasses.go")
	checkFatal(ioerr)
	defer f.Close()
	w := bufio.NewWriter(f)
	w.WriteString(header)
	w.WriteString(templateClassType)
	t := makeTemplate("Emoji classes", templateClassConsts)
	checkFatal(t.Execute(w, emojiClassnames))
	t = makeTemplate("Emoji range tables", templateRangeTableVars)
	checkFatal(t.Execute(w, emojiClassnames))
	t = makeTemplate("Emoji classes stringer", templateClassStringer)
	checkFatal(t.Execute(w, emojiClassnames))
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
