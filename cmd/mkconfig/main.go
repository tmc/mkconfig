// mkconfig renders templates and potentially execs
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tmc/mkconfig"
	"github.com/tmc/mkconfig/sources"
)

var (
	source   = flag.String("source", "source of configuration data", "path to templates for yml file rendering")
	template = flag.String("template", "", "path to templates for yml file rendering")
	out      = flag.String("o", "", "output path")
	dryRun   = flag.Bool("dry", false, "allow writing some configuration files if other fail")
)

func main() {
	flag.Parse()
	if *source == "" || *template == "" {
		flag.Usage()
		os.Exit(1)
	}
	source, err := sources.Get(*source)
	if err != nil {
		fmt.Fprintln(os.Stderr, "mkconfig: bad source:", err)
		os.Exit(1)
	}
	if err := mkconfig.Render(source, *template, *out, *dryRun); err != nil {
		fmt.Fprintln(os.Stderr, "mkconfig: error rendering config:", err)
		os.Exit(1)
	}
}
