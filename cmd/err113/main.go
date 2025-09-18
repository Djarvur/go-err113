package main

import (
	"golang.org/x/tools/go/analysis/singlechecker"

	"github.com/Djarvur/go-err113"
)

func main() {
	singlechecker.Main(err113.NewAnalyzer())
}
