package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/doc"
	"go/parser"
	"go/token"
	"go/types"
)

var path, tp string

func init() {
	flag.StringVar(&path, "p", ".", "path")
	flag.StringVar(&tp, "t", "string", "type")
}

func parseDocs(path, tp string) (b []byte, err error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		return
	}
	var buf bytes.Buffer
	for _, p := range pkgs {
		p := doc.New(p, ".", doc.AllDecls)
		for _, f := range p.Funcs {
			if len(f.Doc) > 0 {
				ps := f.Decl.Type.Params
				if ps.NumFields() == 1 {
					for _, i := range ps.List {
						if types.ExprString(i.Type) == tp {
							buf.WriteString(f.Doc)
							buf.WriteString("\n")
						}
					}
				}
			}
		}
	}
	buf.Truncate(buf.Len() - 1)
	return buf.Bytes(), nil
}

func main() {
	flag.Parse()

	b, _ := parseDocs(path, tp)
	fmt.Printf("%s\n", b)
}
