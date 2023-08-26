package inspector

import (
	"cbsutil/reflext"
	"fmt"
	"reflect"

	"fyne.io/fyne/v2/widget"
)

type Editor interface {
	CreateInspectorGUI(form *widget.Form, label string) error
}

type EditorFactory func(v any) (Editor, error)

var editors map[string]EditorFactory = map[string]EditorFactory{}

func RegisterEditorFactoryByType(typ reflect.Type, f EditorFactory) {
	n := reflext.TypeFullname(typ)
	RegisterEditorFactoryByName(n, f)
}

func RegisterEditorFactoryByName(n string, f EditorFactory) {
	editors[n] = f
}

func GetEditorFactory(n string) EditorFactory {
	return editors[n]
}

func CreateEditor(v any, n string) (Editor, error) {
	if v == nil {
		return nil, nil
	}
	if ed, ok := v.(Editor); ok {
		return ed, nil
	}
	if n == "" {
		typ := reflect.TypeOf(v)
		n = reflext.TypeFullname(typ)
	}
	efac := GetEditorFactory(n)
	if efac == nil {
		return nil, fmt.Errorf("miss Editor '%s'", n)
	}
	return efac(v)
}
