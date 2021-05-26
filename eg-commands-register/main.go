package main

import (
	"fmt"
	"os"

	"github.com/jpillora/opts"
	"github.com/jpillora/opts-examples/eg-commands-register/foo"
)

type cmd struct{}

func main() {
	c := cmd{}
	//default name for the root command (package main) is the binary name
	parsedOpts, err := opts.New(&c).
		AddCommand(foo.New()).
		Parse().
		Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, parsedOpts.Help())
		fmt.Fprintf(os.Stderr, "Error:\n\t%s\n", err)
	}

}
