package inspector

import (
	"fmt"
	"goapp_commons"
	"reflect"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
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

type Property struct {
	object *Object
	parent *Property
	name   string
	index  int

	displayName string
	editorType  string
	propKind    PropertyKind
	propHandler PropertyHandler
	s2v         func(string) (any, error)
	v2s         func(any) (string, error)
	dataKind    DataKind
	dataHandler any

	editor  Editor
	content fyne.CanvasObject
}

func (this *Property) String() string {
	return fmt.Sprintf("Property[%v, %s]", this.object.target, this.GetPath())
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
func (this *Property) GetPropKind() PropertyKind {
	return this.propKind
}
func (this *Property) GetPropertyHandler() PropertyHandler {
	return this.propHandler
}
func (this *Property) GetDataKind() DataKind {
	return this.dataKind
}
func (this *Property) GetDataHandler() any {
	return this.dataHandler
}
func (this *Property) CreateProperty(typeName string) (*Property, error) {
	return CreateProperty(this.object, this, typeName)
}
func (this *Property) S2V() func(string) (any, error) {
	return this.s2v
}
func (this *Property) V2S() func(any) (string, error) {
	return this.v2s
}
func (this *Property) GetEditor() Editor {
	return this.editor
}
func (this *Property) GetEditorContent() fyne.CanvasObject {
	return this.content
}

func (this *Property) FindProperty(n string) (*Property, error) {
	switch this.dataKind {
	case DataKind_Value:
		return nil, nil
	case DataKind_Struct:
		h := this.dataHandler.(DataStructHandler)
		return h.GetProperty(this, n)
	case DataKind_List:
		idx, err := strconv.Atoi(n)
		if err != nil {
			return nil, err
		}
		h := this.dataHandler.(DataListHandler)
		return h.GetElementAtIndex(this, idx)
	default:
		panic(fmt.Sprintf("unknow Kind %d", this.dataKind))
	}
}

func (this *Property) GetValue() (any, error) {
	if this.parent == nil {
		return this.object.target, nil
	} else {
		v, err := this.parent.GetValue()
		if err != nil {
			return nil, err
		}
		return this.propHandler.GetValue(this, v)
	}
}

func (this *Property) SetValue(v any) error {
	if this.parent == nil {
		// can't change root
		return nil
	} else {
		t, err := this.parent.GetValue()
		if err != nil {
			return err
		}
		return this.propHandler.SetValue(this, t, v)
	}
}

func (this *Property) GetText() (string, error) {
	v, err := this.GetValue()
	if err != nil {
		return "", err
	}
	if this.v2s != nil {
		return this.v2s(v)
	}
	return V2S_Any(v)
}

func (this *Property) SetText(s string) error {
	if this.parent == nil {
		// can't change root
		return nil
	} else {
		t, err := this.parent.GetValue()
		if err != nil {
			return err
		}
		var v any
		if this.v2s == nil {
			v = s
		} else {
			v, err = this.v2s(s)
			if err != nil {
				return err
			}
		}
		return this.propHandler.SetValue(this, t, v)
	}
}

type PropertyKind int

const (
	PropertyKind_Root PropertyKind = iota
	PropertyKind_Field
	PropertyKind_Element
)

type PropertyHandler interface {
	GetValue(self *Property, target any) (any, error)

	SetValue(self *Property, target any, value any) error
}

type DataKind int

const (
	DataKind_Value DataKind = iota
	DataKind_Struct
	DataKind_List
)

type DataStructHandler interface {
	GetPropertyCount(self *Property) (int, error)

	GetProperty(self *Property, name string) (*Property, error)

	Properties(self *Property) ([]*Property, error)
}

type DataListHandler interface {
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

func (this *PropertyBuilder) WithPropHandler(k PropertyKind, h PropertyHandler) *PropertyBuilder {
	this.prop.propKind = k
	this.prop.propHandler = h
	return this
}

func (this *PropertyBuilder) WithS2V(h func(string) (any, error)) *PropertyBuilder {
	this.prop.s2v = h
	return this
}
func (this *PropertyBuilder) WithV2S(h func(any) (string, error)) *PropertyBuilder {
	this.prop.v2s = h
	return this
}

func (this *PropertyBuilder) BeValue() *PropertyBuilder {
	this.prop.dataKind = DataKind_Value
	this.prop.dataHandler = nil
	return this
}

func (this *PropertyBuilder) BeStruct(h DataStructHandler) *PropertyBuilder {
	this.prop.dataKind = DataKind_Struct
	this.prop.dataHandler = h
	return this
}

func (this *PropertyBuilder) BeList(h DataListHandler) *PropertyBuilder {
	this.prop.dataKind = DataKind_List
	this.prop.dataHandler = h
	return this
}

func (this *PropertyBuilder) Build() error {
	return nil
}
