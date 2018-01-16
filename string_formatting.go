package main

import (
	"fmt"
	"os"
)

var pf = fmt.Printf
var p = fmt.Println

type point struct {
	x, y int
}

func main() {
	pnt := point{1, 2}
	// Structs:
	// data only
	pf("%v\n", pnt)
	// include field names
	pf("%+v\n", pnt)
	// include implementation
	pf("%#v\n", pnt)
	// print type
	pf("%T\n", pnt)

	// Booleans
	pf("%t\n", true)

	// Digits:
	// default base-10 formatting
	p("Digits, 42 as base: ")
	pf("Default: %d\n", 42)

	pf("Binary: %b\n", 42)
	pf("Hex: %x\n", 42)

	pf("Corresponding character: %c\n", 42)
	pf("Float: %2.2f\n", 42.4242)

	// Scientific notation
	pf("%e\n", 123400000.0)
	pf("%E\n", 123400000.0)

	// Strings
	pf("%s\n", "\"string\"")
	// Double qouted
	pf("%q\n", "\"string\"")
	// Hex
	pf("%x\n", "hex this")

	pf("%p\n", &pnt)

	pf("|%6d|%6d|\n", 12, 345)
	// You can also specify the width of printed floats, though usually you’ll also want to restrict the decimal precision at the same time with the width.precision syntax.

	pf("|%6.2f|%6.2f|\n", 1.2, 3.45)
	// To left-justify, use the - flag.

	pf("|%-6.2f|%-6.2f|\n", 1.2, 3.45)
	// You may also want to control width when formatting strings, especially to ensure that they align in table-like output. For basic right-justified width.

	pf("|%6s|%6s|\n", "foo", "b")
	// To left-justify use the - flag as with numbers.

	pf("|%-6s|%-6s|\n", "foo", "b")
	// So far we’ve seen Printf, which prints the formatted string to os.Stdout. Sprintf formats and returns a string without printing it anywhere.

	s := fmt.Sprintf("a %s", "string")
	p(s)
	// You can format+print to io.Writers other than os.Stdout using Fprintf.

	fmt.Fprintf(os.Stderr, "an %s\n", "error")
}
