// Static analytic service. Include static analytic packages:
// - golang.org/x/tools/go/analysis/passes
// - all SA classes staticcheck.io
// - Go-critic and nilerr linters
// - OsExitAnalyzer to check os.Exit calls in main packages
//
// How to run:
// ./cmd/staticlint/main ./...

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/multichecker"
	"golang.org/x/tools/go/analysis/passes/httpresponse"
	"golang.org/x/tools/go/analysis/passes/printf"
	"golang.org/x/tools/go/analysis/passes/shadow"
	"golang.org/x/tools/go/analysis/passes/structtag"
	"honnef.co/go/tools/staticcheck"
	"honnef.co/go/tools/stylecheck"

	goc "github.com/go-critic/go-critic/checkers/analyzer"
	"github.com/gostaticanalysis/nilerr"
	"github.com/lipandr/yandex_practicum_url_shortener/cmd/staticlint/analyzer"
)

const Config = `config/config.json`

// ConfigData структура описывающая проверки кода
type ConfigData struct {
	StaticCheck []string
	StyleCheck  []string
}

func main() {
	appFile, err := os.Executable()
	if err != nil {
		panic(err)
	}

	data, err := os.ReadFile(filepath.Join(filepath.Dir(appFile), Config))
	if err != nil {
		panic(err)
	}

	var cfg ConfigData
	if err = json.Unmarshal(data, &cfg); err != nil {
		panic(err)
	}

	analyzers := []*analysis.Analyzer{
		analyzer.OsExitAnalyzer,
		printf.Analyzer,
		shadow.Analyzer,
		structtag.Analyzer,
		httpresponse.Analyzer,
		goc.Analyzer,
		nilerr.Analyzer,
	}

	for _, v := range staticcheck.Analyzers {
		for _, sc := range cfg.StaticCheck {
			if strings.HasPrefix(v.Analyzer.Name, sc) {
				analyzers = append(analyzers, v.Analyzer)
			}
		}
	}
	for _, v := range stylecheck.Analyzers {
		for _, sc := range cfg.StyleCheck {
			if strings.HasPrefix(v.Analyzer.Name, sc) {
				analyzers = append(analyzers, v.Analyzer)
			}
		}
	}

	fmt.Println("Multirunner checks:\n", analyzers)

	multichecker.Main(analyzers...)
}
