package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/doc"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"strings"
)

var path, tp string

func init() {
	flag.StringVar(&path, "p", ".", "path")
	flag.StringVar(&tp, "t", "string", "type")
}

func parseDocs(path string, params []string) (b []byte, err error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		return
	}
	var buf bytes.Buffer
	for _, p := range pkgs {
		p := doc.New(p, ".", doc.AllDecls)
	Loop:
		for _, f := range p.Funcs {
			if len(f.Doc) > 0 {
				ps := f.Decl.Type.Params
				if len(params) != len(ps.List) {
					continue Loop
				}
				for i, v := range ps.List {
					if types.ExprString(v.Type) != params[i] {
						continue Loop
					}
				}
				buf.WriteString(f.Doc)
				buf.WriteString("\n")
			}
		}
	}
	if c := buf.Len(); c > 0 {
		buf.Truncate(c - 1)
	}
	return buf.Bytes(), nil
}

func main() {
	flag.Parse()

	if len(os.Args) == 1 {
		flag.Usage()
		return
	}

	b, _ := parseDocs(path, strings.Split(tp, ","))
	fmt.Printf("%s", b)
}
