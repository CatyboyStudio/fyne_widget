package inspector

import (
	"cbsutil/reflext"
	"fmt"
	"reflect"
	"time"

	"fyne.io/fyne/v2/widget"
)

type Editor struct {
	Data      any
	Inspector *Inspector
	Form      *widget.Form
	closeC    chan bool
	builder   EditorBuilder
	watchers  []func()
}

func (ed *Editor) CloseC() <-chan bool {
	return ed.closeC
}

func (ed *Editor) close() {
	close(ed.closeC)
	ed.Inspector = nil
	ed.Form = nil
	ed.builder = nil
	ed.Data = nil
	clear(ed.watchers)
}

func CreateEditor(ins *Inspector, v any, editorType string) (*Editor, error) {
	if v == nil {
		return nil, nil
	}
	var eb EditorBuilder
	if o, ok := v.(EditorBuilder); ok {
		eb = o
	} else {
		if editorType == "" {
			typ := reflect.TypeOf(v)
			editorType = reflext.TypeFullname(typ)
		}
		eb = GetEditorBuilder(editorType)
	}
	if eb == nil {
		return nil, fmt.Errorf("miss Editor '%s'", editorType)
	}

	ed := &Editor{
		Inspector: ins,
		Data:      v,
		closeC:    make(chan bool),
		builder:   eb,
	}
	return ed, nil
}

const tickTime = time.Millisecond * 200

func (ed *Editor) Watch(f func()) {
	ed.watchers = append(ed.watchers, f)
	if len(ed.watchers) == 1 {
		go func() {
			tm := time.NewTimer(tickTime)
			defer tm.Stop()
			for {
				select {
				case <-tm.C:
					if ed.Inspector != nil {
						ed.Inspector.Executor.Process(nil, func(a any) (any, error) {
							for _, w := range ed.watchers {
								w()
							}
							return nil, nil
						})
						tm.Reset(tickTime)
					}
				case <-ed.closeC:
					return
				}
			}
		}()
	}
}
