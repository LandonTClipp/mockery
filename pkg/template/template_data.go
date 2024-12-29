package template

import (
	"fmt"
	"go/types"
	"strings"

	"github.com/vektra/mockery/v2/pkg/registry"
)

// Data is the template data used to render the Moq template.
type Data struct {
	Boilerplate     string
	BuildTags       string
	PkgName         string
	SrcPkgQualifier string
	Imports         []*registry.Package
	Mocks           []MockData
	StubImpl        bool
	SkipEnsure      bool
	WithResets      bool
	TemplateData    map[string]any
}

// MocksSomeMethod returns true of any one of the Mocks has at least 1
// method.
func (d Data) MocksSomeMethod() bool {
	for _, m := range d.Mocks {
		if len(m.Methods) > 0 {
			return true
		}
	}

	return false
}

// MockData is the data used to generate a mock for some interface.
type MockData struct {
	InterfaceName string
	MockName      string
	TypeParams    []TypeParamData
	Methods       []MethodData
	TemplateData  map[string]any
}

// MethodData is the data which represents a method on some interface.
type MethodData struct {
	Name    string
	Params  []ParamData
	Returns []ParamData
}

// ArgList is the string representation of method parameters, ex:
// 's string, n int, foo bar.Baz'.
func (m MethodData) ArgList() string {
	params := make([]string, len(m.Params))
	for i, p := range m.Params {
		params[i] = p.MethodArg()
	}
	return strings.Join(params, ", ")
}

// ArgTypeList returns the argument types in a comma-separated string, ex:
// `string, int, bar.Baz`
func (m MethodData) ArgTypeList() string {
	params := make([]string, len(m.Params))
	for i, p := range m.Params {
		params[i] = p.TypeString()
	}
	return strings.Join(params, ", ")
}

// ArgCallList is the string representation of method call parameters,
// ex: 's, n, foo'. In case of a last variadic parameter, it will be of
// the format 's, n, foos...'
func (m MethodData) ArgCallList() string {
	return m.ArgCallListSlice(0, -1)
}

// ArgCallListSlice is similar to ArgCallList, but it allows specification of
// a slice range to use for the parameter lists. Specifying an integer less than
// 1 for end indicates to slice to the end of the parameters. As with regular
// Go slicing semantics, the end value is a non-inclusive index.
func (m MethodData) ArgCallListSlice(start, end int) string {
	if end < 0 {
		end = len(m.Params)
	}
	paramsSlice := m.Params[start:end]
	params := make([]string, len(paramsSlice))
	for i, p := range paramsSlice {
		params[i] = p.CallName()
	}
	return strings.Join(params, ", ")
}

// ReturnArgTypeList is the string representation of method return
// types, ex: 'bar.Baz', '(string, error)'.
func (m MethodData) ReturnArgTypeList() string {
	params := make([]string, len(m.Returns))
	for i, p := range m.Returns {
		params[i] = p.TypeString()
	}
	if len(m.Returns) > 1 {
		return fmt.Sprintf("(%s)", strings.Join(params, ", "))
	}
	return strings.Join(params, ", ")
}

// ReturnArgNameList is the string representation of values being
// returned from the method, ex: 'foo', 's, err'.
func (m MethodData) ReturnArgNameList() string {
	params := make([]string, len(m.Returns))
	for i, p := range m.Returns {
		params[i] = p.Name()
	}
	return strings.Join(params, ", ")
}

func (m MethodData) IsVariadic() bool {
	return len(m.Params) > 0 && m.Params[len(m.Params)-1].Variadic
}

type TypeParamData struct {
	ParamData
	Constraint types.Type
}

// ParamData is the data which represents a parameter to some method of
// an interface.
type ParamData struct {
	Var      *registry.Var
	Variadic bool
}

// Name returns the name of the parameter.
func (p ParamData) Name() string {
	return p.Var.Name
}

// MethodArg is the representation of the parameter in the function
// signature, ex: 'name a.Type'.
func (p ParamData) MethodArg() string {
	if p.Variadic {
		return fmt.Sprintf("%s ...%s", p.Name(), p.TypeString()[2:])
	}
	return fmt.Sprintf("%s %s", p.Name(), p.TypeString())
}

// CallName returns the string representation of the parameter to be
// used for a method call. For a variadic paramter, it will be of the
// format 'foos...'.
func (p ParamData) CallName() string {
	if p.Variadic {
		return p.Name() + "..."
	}
	return p.Name()
}

// TypeString returns the string representation of the type of the
// parameter.
func (p ParamData) TypeString() string {
	return p.Var.TypeString()
}
