package main

import (
	"encoding/json"
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
	f, err := parser.ParseFile(fset, "", srcCode, 0)
	if err != nil {
		fmt.Printf("err = %s", err)
	}
	ast.Print(fset, f)
	jBytes, _ := json.MarshalIndent(f, "", "  ")
	fmt.Println(string(jBytes))
}
