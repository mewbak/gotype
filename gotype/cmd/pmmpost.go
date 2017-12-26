package cmd

/*
BSD License
Copyright (c) 2017, Norbert Pillmayer <norbert@pillmayer.com>

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

----------------------------------------------------------------------

Command line interface for the Poor Man's MetaPost language and
graphical system.

*/

import (
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/chzyer/readline"
	"github.com/mitchellh/colorstring"
	"github.com/npillmayer/gotype/gtbackend/gfx"
	"github.com/npillmayer/gotype/gtbackend/gfx/png"
	"github.com/npillmayer/gotype/gtcore/config"
	"github.com/npillmayer/gotype/gtcore/config/tracing"
	"github.com/npillmayer/gotype/syntax/pmmpost"
	"github.com/spf13/cobra"
)

var T tracing.Trace = tracing.CommandTracer

// global settings
const toolname = "pmmp"
const welcomeMessage = "Welcome to Poor Man's MetaPost [V0.1 experimental]"
const stdprompt = "[green]pmmpost> "

var editmode string = "emacs"

// pmmpostCmd represents the pmmp command
var pmmpostCmd = &cobra.Command{
	Use:   "pmmp",
	Short: "A poor man's MetaPost implementation",
	Long: `pmmp is a drawing language und engine reminiscend of John Hobby's
	MetaPost sytem. Users supply an input program, either as a text file
	or on the command prompt. Output may be generated as PDF, SVG or PNG.`,
	Args: cobra.MaximumNArgs(1),
	Run:  PMMPostCmd,
}

func init() {
	rootCmd.AddCommand(pmmpostCmd)
	// Here you will define your flags and configuration settings.
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pmmpostCmd.PersistentFlags().String("foo", "", "A help for foo")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	pmmpostCmd.Flags().BoolP("vi", "m", false, "Set vi editing mode")
	pmmpostCmd.Flags().StringP("outdir", "o", ".", "Output directory")
}

// We support some sub-commands (not part of the PMMP grammar).
func displayCommands(out io.Writer) {
	io.WriteString(out, welcomeMessage)
	io.WriteString(out, "\n\nThe following commands are available:\n\n")
	io.WriteString(out, "help:               print this message\n")
	io.WriteString(out, "bye:                quit\n")
	io.WriteString(out, "mode [mode]:        display or set current editing mode\n")
	io.WriteString(out, "setprompt [prompt]: set current editing mode [to default],\n")
	io.WriteString(out, "                    supports color strings, e.g. '[blue]myprompt#'\n\n")
}

// Completer-tree for pmmp sub-commands
var replCompleter = readline.NewPrefixCompleter(
	readline.PcItem("help"),
	readline.PcItem("bye"),
	readline.PcItem("mode",
		readline.PcItem("vi"),
		readline.PcItem("emacs"),
	),
	readline.PcItem("setprompt"),
)

// A helper type to instantiate a REPL interpreter.
type PMMPostREPL struct {
	interpreter *pmmpost.PMMPostInterpreter
	readline    *readline.Instance
}

/* Set up a new REPL entity. It contains a readline-instance (for putting
 * out a prompt and read in a line) and a PMMetaPost parser. The REPL will
 * then forward PMMetaPost statements to the parser.
 */
func NewPMMPostREPL() *PMMPostREPL {
	rl, err := readline.NewEx(&readline.Config{
		Prompt:              colorstring.Color(stdprompt),
		HistoryFile:         "/tmp/pmmpost-repl-history.tmp",
		AutoComplete:        replCompleter,
		InterruptPrompt:     "^C",
		EOFPrompt:           "exit",
		HistorySearchFold:   true,
		FuncFilterInputRune: filterReplInput,
	})
	if err != nil {
		panic(err)
	}
	repl := &PMMPostREPL{
		interpreter: pmmpost.NewPMMPostInterpreter(), //  TODO: re-factor name clash...
		readline:    rl,
	}
	gfx.GlobalCanvasFactory = png.NewContextFactory()            // use GG drawing package
	repl.interpreter.SetOutputRoutine(png.NewPNGOutputRoutine()) // will produce PNG format
	return repl
}

/* Interactive version of the PMMpost command.
 */
func PMMPostCmd(cmd *cobra.Command, args []string) {
	fmt.Println(welcomeMessage)
	config.Initialize()
	var inputfilename string
	if len(args) > 0 {
		T.Infof("input file is %s", args[0])
		inputfilename = args[0]
	}
	tracing.Tracefile = tracing.ConfigTracing(inputfilename)
	defer tracing.Tracefile.Close()
	startInput(inputfilename)
}

func startInput(inputfilename string) {
	if inputfilename == "" {
		config.IsInteractive = true
		repl := NewPMMPostREPL() // go into interactive mode
		log.SetOutput(repl.readline.Stderr())
		defer repl.readline.Close()
		repl.doLoop()
	} else {
		input, err := antlr.NewFileStream(inputfilename) // TODO refactor to get rid of ANTLR
		if err != nil {
			T.Error("cannot open input file")
		} else {
			config.IsInteractive = false
			/*
				defer func() {
					if r := recover(); r != nil {
						T.Error("error executing PMMPost statement!")
					}
				}()
			*/
			interpreter := pmmpost.NewPMMPostInterpreter()
			interpreter.ParseStatements(input)
		}
	}
}

/* Enter a REPL and execute each command.
 * Commands are either tool-commands (setprompt, help, etc.)
 * or PMMetaPost statements.
 */
func (repl *PMMPostREPL) doLoop() {
	for {
		line, err := repl.readline.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}
		line = strings.TrimSpace(line)
		words := strings.Fields(line)
		command := "<no command>"
		if len(words) > 0 {
			command = words[0]
		}
		//log.Println("Executing command", command)
		//log.Println("   Arguments are:", words)
		if doExit := repl.executeCommand(command, words, line); doExit {
			break
		}
	}
}

/* Central dispatcher function to execute internal commands and PMMetaPost
 * statements. It receives the command (i.e. the first word of the line),
 * a list of words (args) including the command, and the complete line of text.
 * If it returns true, the REPL should terminate.
 */
func (repl *PMMPostREPL) executeCommand(cmd string, args []string, line string) bool {
	switch {
	case cmd == "help":
		displayCommands(repl.readline.Stderr())
	case cmd == "bye":
		println("> goodbye!")
		return true
	case cmd == "mode":
		if len(args) > 1 {
			switch args[1] {
			case "vi":
				repl.readline.SetVimMode(true)
				editmode = "vi"
				return false
			case "emacs":
				repl.readline.SetVimMode(false)
				editmode = "emacs"
				return false
			}
		}
		io.WriteString(repl.readline.Stderr(),
			fmt.Sprintf("> current input mode: %s\n", editmode))
	case cmd == "setprompt":
		var prmpt string
		if len(line) <= 10 {
			prmpt = stdprompt
		} else {
			prmpt = line[10:] + " "
		}
		repl.readline.SetPrompt(colorstring.Color(prmpt))
	case cmd == "":
	default:
		T.Debugf("call PMMPost parser on: '%s'", line)
		repl.callPMMPostInterpreter(line)
	}
	return false // do not exit
}

func (repl *PMMPostREPL) callPMMPostInterpreter(line string) {
	input := antlr.NewInputStream(line) // TODO refactor to get rid of ANTLR
	/*
		defer func() {
			if r := recover(); r != nil {
				io.WriteString(repl.readline.Stderr(), "> error executing PMMPost statements!\n")
				io.WriteString(repl.readline.Stderr(), fmt.Sprintf("> %v\n", r)) // TODO: get ERROR and print
			}
		}()
	*/
	repl.interpreter.ParseStatements(input)
}

/* Input filter for REPL. Blocks ctrl-z.
 */
func filterReplInput(r rune) (rune, bool) {
	switch r {
	// block CtrlZ feature
	case readline.CharCtrlZ:
		return r, false
	}
	return r, true
}