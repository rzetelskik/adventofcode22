package generate

import (
	_ "embed"
	"fmt"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed templates/solution.go.tmpl
var solutionTemplate string

//go:embed templates/solution_test.go.tmpl
var solutionTestTemplate string

const (
	dayPathFormat        = "day%02d"
	solutionFilename     = "solution.go"
	solutionTestFilename = "solution_test.go"
)

type Generator struct {
	path        string
	day         int
	packageName string
	packagePath string
}

func NewGenerator(path string, day int) (*Generator, error) {
	var err error

	g := &Generator{
		path: path,
		day:  day,
	}

	err = g.validate()
	if err != nil {
		return nil, err
	}

	err = g.complete()
	if err != nil {
		return nil, err
	}

	return g, nil
}

func (g *Generator) validate() error {
	return nil
}

func (g *Generator) complete() error {
	g.packagePath = filepath.Join(g.path, fmt.Sprintf(dayPathFormat, g.day))
	g.packageName = fmt.Sprintf(dayPathFormat, g.day)

	return nil
}

func (g *Generator) Run() error {
	var err error

	err = g.createPackage()
	if err != nil {
		return fmt.Errorf("can't create package: %w", err)
	}

	err = g.generateCode(solutionTemplate, solutionFilename)
	if err != nil {
		return fmt.Errorf("can't generate code: %w", err)
	}

	err = g.generateCode(solutionTestTemplate, solutionTestFilename)
	if err != nil {
		return fmt.Errorf("can't generate code: %w", err)
	}

	return nil
}

func (g *Generator) createPackage() error {
	var err error

	err = os.MkdirAll(g.packagePath, 0755)
	if err != nil {
		return fmt.Errorf("can't create directory at %s: %w", g.packagePath, err)
	}

	return nil
}

func (g *Generator) generateCode(templateName string, filename string) error {
	path := filepath.Join(g.packagePath, filename)

	t, err := template.New(g.packageName).Parse(templateName)
	if err != nil {
		return fmt.Errorf("can't parse template: %w", err)
	}

	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("can't create file %s: %w", path, err)
	}

	err = t.Execute(f, nil)
	if err != nil {
		return fmt.Errorf("can't execute template: %w", err)
	}

	return nil
}
