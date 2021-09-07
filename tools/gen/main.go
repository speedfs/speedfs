package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var typeMap = map[string]bool{
	"byte":   true,
	"bool":   true,
	"int8":   true,
	"int16":  true,
	"int32":  true,
	"int64":  true,
	"uint8":  true,
	"uint16": true,
	"uint32": true,
	"uint64": true,
	"string": true,
}

var (
	input  string
	output string
)

func main() {
	flag.StringVar(&input, "input", "", "input file")
	flag.StringVar(&output, "output", "", "output file")
	flag.Parse()

	if _, err := ioutil.ReadFile(input); err != nil {
		log.Fatal(err)
	}

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, input, nil, 0)
	if err != nil {
		log.Fatal(err)
	}

	var b strings.Builder

	// package
	b.WriteString("// Code generated by speedfs-gen. DO NOT EDIT.\n\n")
	b.WriteString(fmt.Sprintf("package %s\n\n", f.Name))
	b.WriteString("import (\n\t\"github.com/speedfs/speedfs/internal/rpc\"\n)\n\n")

	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}
		if genDecl.Tok != token.TYPE {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			if !token.IsExported(typeSpec.Name.Name) {
				continue
			}
			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			// func EncodeTo
			b.WriteString(fmt.Sprintf("func (x *%s) EncodeTo(enc *rpc.Encoder) {\n", typeSpec.Name.Name))
			// fields
			for _, field := range structType.Fields.List {
				switch typ := field.Type.(type) {
				case *ast.Ident:
					if len(field.Names) == 0 {
						break
					}

					if _, ok := typeMap[typ.Name]; ok {
						b.WriteString(fmt.Sprintf("\tenc.Encode%s(x.%s)\n", strings.Title(typ.Name), field.Names[0].Name))
					} else {
						b.WriteString(fmt.Sprintf("\tx.%s.EncodeTo(enc)\n", field.Names[0].Name))
					}
				case *ast.ArrayType:
					identType, ok := typ.Elt.(*ast.Ident)
					if !ok {
						break
					}
					if identType.Name == "byte" {
						b.WriteString(fmt.Sprintf("\tenc.EncodeBytes(x.%s[:])\n", field.Names[0].Name))
					}

				}
			}
			b.WriteString("}\n\n")
		}
	}

	buf, err := format.Source([]byte(b.String()))
	if err != nil {
		log.Fatal(err)
	}

	w, err := os.OpenFile(output, os.O_CREATE|os.O_TRUNC|os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := w.Write(buf); err != nil {
		log.Fatal(err)
	}
}