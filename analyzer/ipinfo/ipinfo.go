package ipinfo

import (
	"go/ast"
	"regexp"

	"github.com/beacon/code-uv/analyzer/ignore"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var ipv4Reg = regexp.MustCompile(`^"(\d{1,3})\.\d{1,3}\.\d{1,3}.\d{1,3}"$`)

const Doc = `check for hardcoded ip info in code`

var Analyzer = &analysis.Analyzer{
	Name:     "ipinfo",
	Doc:      Doc,
	Requires: []*analysis.Analyzer{inspect.Analyzer, ignore.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	ignoreResult := pass.ResultOf[ignore.Analyzer].(*ignore.IgnoreResult)

	nodeFilter := []ast.Node{
		(*ast.BasicLit)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		lit := n.(*ast.BasicLit)
		matches := ipv4Reg.FindStringSubmatch(lit.Value)
		if matches == nil {
			return
		}

		if ignoreResult.IsIgnored(pass, lit.Pos()) {
			return
		}

		// Ignore loopback and masks, this won't cause security issues.
		if matches[1] == "127" || matches[1] == "255" {
			return
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
	})
	return nil, nil
}
