// Pre (1): ANTLR4
// E.g., antlr-4.7.1-complete.jar
// (See go:generate below)

// Pre (2): ANTLR4 Runtime for Go
//$ go get github.com/antlr/antlr4/runtime/Go/antlr
// Optional:
//$ cd $CYGHOME/code/go/src/github.com/antlr/antlr4
//$ git checkout -b antlr-go-runtime tags/4.7.1  // Match antlr-4.7.1-complete.jar -- but unnecessary

//rhu@HZHL4 MINGW64 ~/code/go/src/
//$ go run github.com/rhu1/fgg -v -eval=10 fg/examples/hello/hello.go
//$ go run github.com/rhu1/fgg -v -inline="package main; type A struct {}; func main() { _ = A{} }"
// or
//$ go install
//$ /c/Users/rhu/code/go/bin/fgg.exe ...

// N.B. GoInstall installs to $CYGHOME/code/go/bin (not $WINHOME)

// Assuming "antlr4" alias for (e.g.): java -jar ~/code/java/lib/antlr-4.7.1-complete.jar
//go:generate antlr4 -Dlanguage=Go -o parser/fg parser/FG.g4
//go:generate antlr4 -Dlanguage=Go -o parser/fgg parser/FGG.g4

// FGG gotchas:
// type B(type a Any) struct { f a }; // Any parsed as a TParam -- currently not permitted
// Node(Nat){...} // fgg.FGGNode (Nat) is fgg.TParam, not fgg.TName
// type IA(type ) interface { m1() };  // m1() parsed as a TName (an invalid Spec) -- N.B. ret missing anyway

/* TODO
- WF: repeat type decl

	//b.WriteString("type B struct { f t };\n")  // TODO: unknown type
	//b.WriteString("type B struct { b B };\n")  // TODO: recursive struct
*/

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/rhu1/fgg/base"
	"github.com/rhu1/fgg/fg"
	"github.com/rhu1/fgg/fgg"
)

var _ = reflect.TypeOf
var _ = strconv.Itoa

const (
	EVAL_TO_VAL = -1 // Must be < 0
	NO_EVAL     = -2 // Must be < EVAL_TO_VAL
)

// Command line parameters/flags
var (
	interpFG  bool // parse FG
	interpFGG bool // parse FGG

	monom       bool   // parse FGG and monomorphise FGG source
	monomOutput string // output filename of monomorphised FGG

	useInternalSrc bool   // use internal source
	inlineSrc      string // use content of this as source
	strictParse    bool   // use strict parsing mode

	evalSteps int  // number of steps to evaluate
	verbose   bool // verbose mode
)

func init() {
	// FG or FGG
	flag.BoolVar(&interpFG, "fg", false,
		"interpret input as FG (defaults to true if neither -fg/-fgg set)")
	flag.BoolVar(&interpFGG, "fgg", false,
		"interpret input as FGG")

	// Monomorphise
	flag.BoolVar(&monom, "monom", false,
		"[WIP] monomorphise FGG source using formal notation (ignored if -fgg not set)")
	flag.StringVar(&monomOutput, "compile", "",
		"[WIP] monomorphise FGG source to FG (ignored if -fgg not set)\nspecify '--' to print to stdout")

	// Parsing options
	flag.BoolVar(&useInternalSrc, "internal", false,
		`use "internal" input as source`)
	flag.StringVar(&inlineSrc, "inline", "",
		`-inline="[FG/FGG src]", use inline input as source`)
	flag.BoolVar(&strictParse, "strict", true,
		"strict parsing (don't attempt recovery on parsing errors)")

	flag.IntVar(&evalSteps, "eval", NO_EVAL,
		" N ⇒ evaluate N (≥ 0) steps; or\n-1 ⇒ evaluate to value (or panic)")
	flag.BoolVar(&verbose, "v", false,
		"enable verbose printing")
}

var usage = func() {
	fmt.Fprintf(os.Stderr, `Usage:

	fgg [options] -fg  path/to/file.fg
	fgg [options] -fgg path/to/file.fgg
	fgg [options] -internal
	fgg [options] -inline "package main; type ...; func main() { ... }"

Options:

`)
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()

	// Determine mode
	if !interpFG && !interpFGG {
		interpFG = true
	}

	// Determine source
	var src string
	switch {
	case useInternalSrc: // First priority
		src = internalSrc()
	case inlineSrc != "": // Second priority, i.e. -inline overrules src file
		src = inlineSrc
	default:
		if flag.NArg() < 1 {
			fmt.Fprintln(os.Stderr, "Input error: need a source .go file (or an -inline program)")
			flag.Usage()
		}
		b, err := ioutil.ReadFile(flag.Arg(0))
		if err != nil {
			checkErr(err)
		}
		src = string(b)
	}

	switch {
	case interpFG:
		var a fg.FGAdaptor
		interp(&a, src, strictParse, evalSteps, false, "")
	case interpFGG:
		var a fgg.FGGAdaptor
		interp(&a, src, strictParse, evalSteps, monom, monomOutput)
	}
}

// Pre: monom==true || compile != "" => -fgg is set
func interp(a base.Adaptor, src string, strict bool, steps int, monom bool,
	compile string) {
	vPrintln("\nParsing AST:")
	prog := a.Parse(strict, src) // AST (FGProgram root)
	vPrintln(prog.String())

	vPrintln("\nChecking source program OK:")
	allowStupid := false
	prog.Ok(allowStupid)

	if steps > NO_EVAL {
		eval(prog, steps)
	}

	if monom || compile != "" {
		p_mono := fgg.Monomorph(prog.(fgg.FGGProgram)) // TODO: reformat (e.g., "<...>") to make an actual FG program
		if monom {
			vPrintln("\nMonomorphising, formal notation: [Warning] WIP [Warning]")
			vPrintln(p_mono.String())
		}
		if compile != "" {
			vPrintln("\nMonomorphising, FG output: [Warning] WIP [Warning]")
			out := p_mono.String()
			out = strings.Replace(out, ",,", "", -1)
			out = strings.Replace(out, "<", "", -1)
			out = strings.Replace(out, ">", "", -1)
			if compile == "--" {
				vPrintln(out)
			} else {
				vPrintln("Writing output to: " + compile)
				d1 := []byte(out)
				err := ioutil.WriteFile(compile, d1, 0644)
				checkErr(err)
			}
		}
	}
}

// N.B. currently FG panic comes out implicitly as an underlying run-time panic
// TODO: add explicit FG panics
// If steps == EVAL_TO_VAL, then eval to value
func eval(p base.Program, steps int) {
	allowStupid := true
	vPrintln("\nEntering Eval loop:")
	vPrintln("Decls:")
	for _, v := range p.GetDecls() {
		vPrintln("\t" + v.String() + ";")
	}
	vPrintln("Eval steps:")
	vPrintln(fmt.Sprintf("%6d: %8s %v", 0, "", p.GetExpr())) // Initial prog OK already checked

	done := steps > EVAL_TO_VAL || // Ignore 'done' if num steps fixed (set true, for ||!done below)
		p.GetExpr().IsValue() // O/w evaluate until a val -- here, check if init expr is already a val
	var rule string
	for i := 1; i <= steps || !done; i++ {
		p, rule = p.Eval()
		vPrintln(fmt.Sprintf("%6d: %8s %v", i, "["+rule+"]", p.GetExpr()))
		vPrintln("Checking OK:") // TODO: maybe disable by default, enable by flag
		p.Ok(allowStupid)
		if !done && p.GetExpr().IsValue() {
			done = true
		}
	}
	fmt.Println(p.GetExpr().String()) // Final result
}

// For convenient quick testing -- via flag "-internal=true"
func internalSrc() string {
	Any := "type Any interface {}"
	ToAny := "type ToAny struct { any Any }"
	e := "ToAny{1}"
	return fg.MakeFgProgram(Any, ToAny, e)
}

/* Helpers */

// ECheckErr
func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func vPrintln(x string) {
	if verbose {
		fmt.Println(x)
	}
}
