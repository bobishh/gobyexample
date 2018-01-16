package main

import s "strings"
import "fmt"

var p = fmt.Println

func main() {
	p("Contains:", s.Contains("test", "es"))
	p("HasPrefix:", s.HasPrefix("test", "te"))
	p("HasSuffix:", s.HasSuffix("test", "st"))

	p("Count:", s.Count("test", "t"))
	p("Index:", s.Index("test", "e"))

	p("Join:", s.Join([]string{"ass", "hole"}, "-"))

	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))

	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))

	p("Len: ", len("hello"))
	p("Char:", "hello"[1])
}
