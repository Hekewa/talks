package main

import (
	"github.com/hekewa/talks/gojakarta18072019/example/gocheck"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(gocheck.Analyzer)
}
