package interfacer

import (
	"fmt"
	"go/types"
)

// Type is a simple representation of a single parameter type.
type Type struct {
	Package    string `json:"package,omitempty"`    // package name the type is defined in; empty for builtin
	ImportPath string `json:"importPath,omitempty"` // import path of the package
}

func newType(v *types.Var) (typ Type) {
	typ.setFromType(v.Type(), 0, nil)
	return typ
}

type compositeType interface {
	types.Type
	Elem() types.Type
}

func (typ *Type) setFromType(t types.Type, depth int, orig types.Type) {
	if orig == nil {
		orig = t
	}
	if depth > 128 {
		panic("recursive types not supported: " + orig.String())
	}
	switch t := t.(type) {
	case *types.Basic, *types.Interface, *types.Struct, *types.Signature:
	case *types.Named:
		typ.setFromNamed(t)
	case *types.Pointer:
		typ.setFromType(t.Elem(), depth+1, orig)
	case *types.Map:
		typ.setFromComposite(t, depth, orig)
		typ.setFromType(t.Key(), depth+1, orig)
	case compositeType:
		typ.setFromComposite(t, depth, orig)
	default:
		panic(fmt.Sprintf("internal: t=%T, orig=%T", t, orig))
	}
}

func (typ *Type) setFromNamed(t *types.Named) {
	if typ.Package != "" || typ.ImportPath != "" {
		return
	}
	if pkg := t.Obj().Pkg(); pkg != nil {
		typ.Package = pkg.Name()
		typ.ImportPath = pkg.Path()
	}
}

func (typ *Type) setFromComposite(t compositeType, depth int, orig types.Type) {
	typ.setFromType(t.Elem(), depth+1, orig)
}
