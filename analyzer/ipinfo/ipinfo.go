package ipinfo

import (
	"go/ast"
	"go/token"
	"log"
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
		(*ast.BasicLit)(nil),
	}

	//log.Println("Start...", nodeFilter)
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		literal := n.(*ast.BasicLit)
		if literal.Kind != token.STRING {
			return
		}

		matches := ipv4Reg.FindAllStringSubmatch(literal.Value, -1)
		if matches == nil {
			return
		}

		pass.Report(analysis.Diagnostic{
			Pos:     literal.Pos(),
			Message: "Hardcoded ip may become vulnerable to attacks",
			SuggestedFixes: []analysis.SuggestedFix{
				{
					Message: "Put ip in config values",
					TextEdits: []analysis.TextEdit{
						{Pos: literal.Pos(), End: literal.Pos(), NewText: []byte{}},
					},
				},
			},
		})
		log.Println(matches)
	})
	return nil, nil
}
