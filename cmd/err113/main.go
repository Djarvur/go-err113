package main

import (
	"github.com/Djarvur/go-err113"
	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(err113.NewAnalyzer())
}
