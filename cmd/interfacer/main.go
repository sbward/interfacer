// Command interfaces
package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/rjeczalik/interfaces"
)

var (
	q      = flag.String("for", "", "Type to generate an interface for.")
	as     = flag.String("as", "main.Interface", `Generated interface name.`)
	output = flag.String("o", "-", "Output file.")
)

var tmpl = template.Must(template.New("").Parse(`// Created by interfacer; DO NOT EDIT

package {{.PackageName}}

import (
{{range .Deps}}	"{{.}}"
{{end}})

type {{.InterfaceName}} interface {
{{range .Interface}}	{{.}}
{{end}}}
`))

type Vars struct {
	PackageName   string
	InterfaceName string
	Deps          []string
	Interface     interfaces.Interface
}

func nonil(err ...error) error {
	for _, e := range err {
		if e != nil {
			return e
		}
	}
	return nil
}

func die(v interface{}) {
	fmt.Fprintln(os.Stderr, v)
	os.Exit(1)
}

func main() {
	flag.Parse()
	if *q == "" {
		die("empty -q flag value; see -help for details")
	}
	if *output == "" {
		die("empty -o flag value; see -help for details")
	}
	i, err := interfaces.New(*q)
	if err != nil {
		die(err)
	}
	v := &Vars{
		Deps:      i.Deps(),
		Interface: i,
	}
	if i := strings.IndexRune(*as, '.'); i != -1 {
		v.PackageName = (*as)[:i]
		v.InterfaceName = (*as)[i+1:]
	} else {
		v.InterfaceName = *as
	}
	f := os.Stdout
	if *output != "-" {
		f, err = os.OpenFile(*output, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0755)
		if err != nil {
			die(err)
		}
	}
	err = nonil(tmpl.Execute(f, v), f.Sync(), f.Close())
	if err != nil {
		die(err)
	}
}