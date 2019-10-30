package main

import (
	"fmt"
	"os"

	"github.com/golangq/q"
	"github.com/jpillora/opts"
)

type root struct {
	Refname string `opts:"mode=arg"`
	Oldrev  string `opts:"mode=arg"`
	Newrev  string `opts:"mode=arg"`
}

func main() {
	r := &root{}
	op := opts.New(r).
		Name("update").
		Parse()
	op.RunFatal()
}

func (rt *root) Run() {
	q.Q(rt)
	fmt.Printf("%v\n", rt)
	os.Exit(1)
}
