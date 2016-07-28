package main

import (
	"flag"
	"fmt"
	"go/doc"
	"go/parser"
	"go/token"
	"go/types"
	"log"
)

var path string

func init() {
	flag.StringVar(&path, "p", ".", "path")
}

func main() {
	flag.Parse()

	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		log.Fatalln(err)
	}
	for _, p := range pkgs {
		p := doc.New(p, ".", doc.AllDecls)
		for _, f := range p.Funcs {
			if len(f.Doc) > 0 {
				ps := f.Decl.Type.Params
				if ps.NumFields() == 1 {
					fmt.Println(f.Doc)
					for _, i := range ps.List {
						if types.ExprString(i.Type) == "*iris.Context" {
							fmt.Println(f.Doc) // TODO:
						}
					}
				}
			}
		}
	}
}
