package execute

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"plugin"
	"strings"
)

type Executor struct {
	path string
}

func NewExecutor(path string) (*Executor, error) {
	e := &Executor{
		path: path,
	}

	return e, nil
}

func (e *Executor) runSymbol(plugin *plugin.Plugin, symbol string, in io.Reader, out io.Writer) error {
	sym, err := plugin.Lookup(symbol)
	if err != nil {
		return fmt.Errorf("can't lookup symbol %s: %w", symbol, err)
	}

	s, ok := sym.(func(io.Reader, io.Writer) error)
	if !ok {
		return fmt.Errorf("invalid type of symbol")
	}

	return s(in, out)
}

func (e *Executor) Run() error {
	p, err := plugin.Open(e.path)
	if err != nil {
		return fmt.Errorf("can't open plugin: %w", err)
	}

	in, err := io.ReadAll(os.Stdin)
	if err != nil {
		return fmt.Errorf("can't read from stdin: %w", err)
	}

	var b strings.Builder

	sym1 := "SolveFirstPart"
	err = e.runSymbol(p, sym1, bytes.NewReader(in), &b)
	if err != nil {
		return fmt.Errorf("can't run symbol %s: %w", sym1, err)
	}
	fmt.Println(b.String())

	b.Reset()

	sym2 := "SolveSecondPart"
	err = e.runSymbol(p, sym2, bytes.NewReader(in), &b)
	if err != nil {
		return fmt.Errorf("can't run symbol %s: %w", sym2, err)
	}
	fmt.Println(b.String())

	return nil
}
