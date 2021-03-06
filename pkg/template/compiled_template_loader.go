package template

import (
	"fmt"

	"go.starlark.net/starlark"
)

type CompiledTemplateLoader interface {
	FindCompiledTemplate(string) (*CompiledTemplate, error)
	Load(*starlark.Thread, string) (starlark.StringDict, error)
}

type NoopCompiledTemplateLoader struct {
	tpl *CompiledTemplate
}

func NewNoopCompiledTemplateLoader(tpl *CompiledTemplate) NoopCompiledTemplateLoader {
	return NoopCompiledTemplateLoader{tpl}
}

var _ CompiledTemplateLoader = NoopCompiledTemplateLoader{}

func (l NoopCompiledTemplateLoader) FindCompiledTemplate(_ string) (*CompiledTemplate, error) {
	if l.tpl != nil {
		return l.tpl, nil
	}
	return nil, fmt.Errorf("FindCompiledTemplate is not supported")
}

func (l NoopCompiledTemplateLoader) Load(
	thread *starlark.Thread, module string) (starlark.StringDict, error) {

	return nil, fmt.Errorf("Load is not supported")
}
