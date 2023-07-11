package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func main() {
	var filePath string
	flag.StringVar(&filePath, "file", "", "File path to scan ast")
	flag.Parse()

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		fmt.Printf("Filepath=%s err = %s", filePath, err)
	}
	ast.Print(fset, f)
}
