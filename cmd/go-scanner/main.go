package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/beacon/code-uv/analyzer/ignore"
	"github.com/beacon/code-uv/analyzer/ipinfo"
)

func main() {
	unitchecker.Main(
		ignore.Analyzer,
		ipinfo.Analyzer,
	)
}
