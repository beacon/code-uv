package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/beacon/code-uv/analyzer/ipinfo"
)

func main() {
	unitchecker.Main(ipinfo.Analyzer)
}
