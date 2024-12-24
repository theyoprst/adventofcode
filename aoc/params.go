package aoc

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/theyoprst/adventofcode/must"
	"gopkg.in/yaml.v3"
)

type (
	paramsCtxKeyType string
	Params           map[string]any
)

const paramsContextKey = paramsCtxKeyType("params")

// contextWithParams returns a new context with the params.
func contextWithParams(ctx context.Context, params Params) context.Context {
	return context.WithValue(ctx, paramsContextKey, params)
}

// GetParams returns the params from the context. Panics if no params are found.
func GetParams(ctx context.Context) Params {
	if params, ok := ctx.Value(paramsContextKey).(Params); ok {
		return params
	}
	panic("params not found in context")
}

// paramsForInput returns the params for the input file.
// It finds it in the tests.yaml file in the same directory as the input file.
func paramsForInput(path string) Params {
	dir := filepath.Dir(path)
	base := filepath.Base(path)

	testsYAMLPath := filepath.Join(dir, "tests.yaml")
	data, err := os.ReadFile(testsYAMLPath)
	must.NoError(err)

	var tests Tests
	must.NoError(yaml.Unmarshal(data, &tests))

	for _, input := range tests.Inputs {
		if input.Path == base {
			return input.Params
		}
	}

	return nil
}

func (p Params) get(key string) any {
	if v, ok := p[key]; ok {
		return v
	}
	panic(fmt.Errorf("param %q not found", key))
}

// Int returns the value of the key as an int.
func (p Params) Int(key string) int {
	if n, ok := p.get(key).(int); ok {
		return n
	}
	panic(fmt.Errorf("param %q type mismatch: got %T, want int", key, p[key]))
}

// Bool returns the value of the key as a bool.
func (p Params) Bool(key string) bool {
	if x, ok := p.get(key).(bool); ok {
		return x
	}
	panic(fmt.Errorf("param %q type mismatch: got %T, want bool", key, p[key]))
}
