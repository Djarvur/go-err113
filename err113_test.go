package err113_test

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"

	"github.com/Djarvur/go-err113"
)

func TestErr113(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), err113.NewAnalyzer())
}
