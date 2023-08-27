package inspector

import (
	"log/slog"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type EntryItem struct {
	prop    *Property
	control *widget.Entry
	ed      *Editor
}

func NewEntryItem(p *Property) *EntryItem {
	return &EntryItem{
		prop: p,
	}
}

func (th *EntryItem) Control() *widget.Entry {
	if th.control == nil {
		th.control = widget.NewEntry()
		th.control.OnChanged = func(s string) {
			th.ed.Execute(func() {
				err := th.prop.SetString(s)
				if err != nil {
					slog.Error("SetString fail", "property", th.prop.Title, "error", err)
				} else {
					th.prop.OnUpdate()
				}
			})
		}
		th.Reload()
	}
	return th.control
}

func (th *EntryItem) Validator(v fyne.StringValidator) *EntryItem {
	c := th.Control()
	c.Validator = v
	return th
}

func (th *EntryItem) Bind(ed *Editor) *EntryItem {
	th.ed = ed
	ed.Form.Append(th.prop.Title, th.Control())
	return th
}

func (th *EntryItem) Watch() {
	if th.ed == nil {
		panic("not bind editor")
	}
	th.ed.Watch(th.Reload)
}

func (th *EntryItem) Reload() {
	if th.control == nil {
		return
	}
	s, err := th.prop.GetString()
	if err != nil {
		slog.Error("reload fail", "property", th.prop.Title, "error", err)
	}
	th.control.Text = s
	th.control.Refresh()
}
