package execute

import (
	"fmt"
	"io"
	"os"
	"plugin"
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

func (e *Executor) runSymbol(plugin *plugin.Plugin, symbol string) (int, error) {
	sym, err := plugin.Lookup(symbol)
	if err != nil {
		return 0, fmt.Errorf("can't lookup symbol %s: %w", symbol, err)
	}

	s, ok := sym.(func(io.Reader, io.Writer) (int, error))
	if !ok {
		return 0, fmt.Errorf("invalid type of symbol")
	}

	return s(os.Stdin, os.Stdout)
}

func (e *Executor) Run() error {
	p, err := plugin.Open(e.path)
	if err != nil {
		return fmt.Errorf("can't open plugin: %w", err)
	}

	sym1 := "SolveFirstPart"
	res1, err := e.runSymbol(p, sym1)
	if err != nil {
		return fmt.Errorf("can't run symbol %s: %w", sym1, err)
	}
	fmt.Println(res1)

	sym2 := "SolveSecondPart"
	res2, err := e.runSymbol(p, sym1)
	if err != nil {
		return fmt.Errorf("can't run symbol %s: %w", sym2, err)
	}
	fmt.Println(res2)

	return nil
}
