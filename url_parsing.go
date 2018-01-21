package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	puts := fmt.Println
	s := "postgres://user:pass@host.com:5432/path?k=v#f"
	u, err := url.Parse(s)

	if err != nil {
		panic(err)
	}

	puts(u.Scheme)
	puts(u.User)
	puts(u.User.Username())
	p, _ := u.User.Password()
	puts(p)

	puts(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	puts(host)
	puts(port)

	puts(u.Path)
	puts(u.Fragment)

	// To get query params in a string of k=v format, use RawQuery. You can also parse query params into a map. The parsed query param maps are from strings to slices of strings, so index into [0] if you only want the first value.
	puts(u.RawQuery)

	m, _ := url.ParseQuery(u.RawQuery)
	puts(m)
	puts(m["k"][0])
}
