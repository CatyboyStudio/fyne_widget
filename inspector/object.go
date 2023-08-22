package inspector

import (
	"fmt"
	"goapp_commons"
	"reflect"
	"strconv"
	"strings"
)

var mapPropFactory map[string]func(*PropertyBuilder) error = map[string]func(*PropertyBuilder) error{}

func RegisterPropertyFactory(n string, f func(*PropertyBuilder) error) {
	mapPropFactory[n] = f
}

type Object struct {
	target any
	root   *Property
}

func NewInspectableObject(data any) (*Object, error) {
	if data == nil {
		return nil, nil
	}
	typ := reflect.TypeOf(data)
	n := goapp_commons.TypeFullname(typ)
	o := &Object{
		target: data,
	}
	prop, err := CreateProperty(o, nil, n)
	if err != nil {
		return nil, err
	}
	o.root = prop
	return o, nil
}

func CreateProperty(obj *Object, parent *Property, typeName string) (*Property, error) {
	if fac, ok := mapPropFactory[typeName]; ok {
		prop := &Property{
			object: obj,
			parent: parent,
		}
		pb := &PropertyBuilder{prop}
		err := fac(pb)
		if err != nil {
			return nil, err
		}
		return prop, nil
	}
	return nil, fmt.Errorf("unknow property '%s'", typeName)
}

func (this *Object) GetTarget() any {
	return this.target
}

func (this *Object) GetRootProperty() *Property {
	return this.root
}

func (this *Object) FindProperty(propertyPath string) (*Property, error) {
	plist := strings.Split(propertyPath, ".")
	prop := this.root
	for i := 0; i < len(plist); i++ {
		n := plist[i]
		sprop, err := prop.FindProperty(n)
		if err != nil {
			return nil, err
		}
		prop = sprop
	}
	return prop, nil
}

type PropertyKind int

const (
	PropertyKind_Value = iota
	PropertyKind_Struct
	PropertyKind_List
)

type Property struct {
	object *Object
	parent *Property
	name   string
	index  int

	displayName string
	editorType  string
	kind        PropertyKind
	handler     any
}

func (this *Property) GetObject() *Object {
	return this.object
}

func (this *Property) GetParent() *Property {
	return this.parent
}
func (this *Property) GetName() string {
	return this.name
}
func (this *Property) GetIndex() int {
	return this.index
}
func (this *Property) UpdateIndex(idx int) {
	this.name = strconv.FormatInt(int64(idx), 10)
	this.index = idx
}

func (this *Property) GetPath() string {
	if this.parent == nil {
		return ""
	}
	p := this.parent.GetPath()
	if p == "" {
		return this.name
	} else {
		return fmt.Sprintf("%s.%s", p, this.name)
	}
}
func (this *Property) GetDisplayName() string {
	return this.displayName
}
func (this *Property) GetEditorType() string {
	return this.editorType
}
func (this *Property) GetKind() PropertyKind {
	return this.kind
}
func (this *Property) GetHandler() any {
	return this.handler
}

func (this *Property) FindProperty(n string) (*Property, error) {
	switch this.kind {
	case PropertyKind_Value:
		return nil, nil
	case PropertyKind_Struct:
		h := this.handler.(PropertyStructHandler)
		return h.GetProperty(this, n)
	case PropertyKind_List:
		idx, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		h := this.handler.(PropertyListHandler)
		return h.GetElementAtIndex(this, idx)
	default:
		panic(fmt.Sprintf("unknow Kind %d", this.kind))
	}
}

type PropertyValueHandler interface {
	SetValue(self *Property, value any) error

	GetValue(self *Property) (any, error)
}

type PropertyStructHandler interface {
	GetPropertyCount(self *Property) (int, error)

	GetProperty(self *Property, name string) (*Property, error)

	Properties(self *Property) ([]*Property, error)
}

type PropertyListHandler interface {
	GetElementCount(self *Property) (int, error)

	InsertElementAtIndex(self *Property, idx int) (*Property, error)
	MoveElement(self *Property, src, dest int) error
	DeleteElementAtIndex(self *Property, idx int) error
	ClearElements(self *Property) error
	GetElementAtIndex(self *Property, idx int) (*Property, error)
}

type PropertyBuilder struct {
	prop *Property
}

func (this *PropertyBuilder) WithDisplayName(n string) *PropertyBuilder {
	this.prop.displayName = n
	return this
}

func (this *PropertyBuilder) WithEditor(editorType string) *PropertyBuilder {
	this.prop.editorType = editorType
	return this
}

func (this *PropertyBuilder) WithName(n string) *PropertyBuilder {
	this.prop.name = n
	this.prop.index = 0
	return this
}

func (this *PropertyBuilder) WithIndex(idx int) *PropertyBuilder {
	this.prop.UpdateIndex(idx)
	return this
}

func (this *PropertyBuilder) BeValue(h PropertyValueHandler) *PropertyBuilder {
	this.prop.kind = PropertyKind_Value
	this.prop.handler = h
	return this
}

func (this *PropertyBuilder) BeStruct(h PropertyStructHandler) *PropertyBuilder {
	this.prop.kind = PropertyKind_Struct
	this.prop.handler = h
	return this
}

func (this *PropertyBuilder) BeList(h PropertyListHandler) *PropertyBuilder {
	this.prop.kind = PropertyKind_List
	this.prop.handler = h
	return this
}

func (this *PropertyBuilder) Build() {
}
