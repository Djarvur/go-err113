package err113_test

import (
	"testing"

	"github.com/Djarvur/go-err113"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestErr113(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), err113.NewAnalyzer())
}
