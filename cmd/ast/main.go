package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

var srcCode = `
package hello

import "fmt"

func greet() {
    var msg = "Hello World!"
    fmt.Println(msg)
}
`

func main() {
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, "example/ipinfo.go", nil, 0)
	if err != nil {
		fmt.Printf("err = %s", err)
	}
	ast.Print(fset, f)
}
