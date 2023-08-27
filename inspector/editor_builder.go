package inspector

import (
	"cbsutil/reflext"
	"reflect"
)

type EditorBuilder interface {
	BuildEditor(ed *Editor) error
}

var editorBuilders map[string]EditorBuilder = map[string]EditorBuilder{}

func RegisterEditorBuilderByType(typ reflect.Type, f EditorBuilder) {
	n := reflext.TypeFullname(typ)
	RegisterEditorByName(n, f)
}

func RegisterEditorByName(n string, f EditorBuilder) {
	editorBuilders[n] = f
}

func GetEditorBuilder(n string) EditorBuilder {
	return editorBuilders[n]
}
