package ipinfo

import (
	"go/ast"
	"log"
	"reflect"
	"regexp"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var ipv4Reg = regexp.MustCompile(`^(\d{1,3})\.\d{1,3}\.\d{1,3}.\d{1,3}$`)

const Doc = `check for hardcoded ip info in code`

var Analyzer = &analysis.Analyzer{
	Name:     "ipinfo",
	Doc:      Doc,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.GenDecl)(nil),
	}

	log.Println("Start...", nodeFilter, "Inspect=", inspect)
	inspect.Nodes(nodeFilter, func(n ast.Node, push bool) (proceed bool) {
		log.Println("Got node:", n)
		genDecl, ok := n.(*ast.GenDecl)
		if !ok {
			return true
		}

		for _, spec := range genDecl.Specs {
			if valueSpec, ok := spec.(*ast.ValueSpec); ok {
				for _, val := range valueSpec.Values {
					lit := val.(*ast.BasicLit)
					log.Println("String value=", lit.Value)
					matches := ipv4Reg.FindAllStringSubmatch(lit.Value, -1)
					if matches == nil {
						return true
					}
					pass.Report(analysis.Diagnostic{
						Pos:     lit.Pos(),
						Message: "Hardcoded ip may become vulnerable to attacks",
						SuggestedFixes: []analysis.SuggestedFix{
							{
								Message: "Put ip in config values",
								TextEdits: []analysis.TextEdit{
									{Pos: lit.Pos(), End: lit.Pos(), NewText: []byte{}},
								},
							},
						},
					})
					log.Println(matches)
				}
			} else {
				log.Println("Cannot cast to value spec:", spec, reflect.TypeOf(spec))
			}
		}
		return true
	})
	return nil, nil
}
