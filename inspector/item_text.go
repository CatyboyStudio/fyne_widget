package inspector

import (
	"log/slog"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type TextItem struct {
	prop    *Property
	control *widget.Label
	ed      *Editor
}

func NewTextItem(p *Property) *TextItem {
	return &TextItem{
		prop: p,
	}
}

func (th *TextItem) Control() *widget.Label {
	if th.control == nil {
		th.control = widget.NewLabel("")
		th.control.Wrapping = fyne.TextTruncate
		th.Reload()
	}
	return th.control
}

func (th *TextItem) Wrapping(v fyne.TextWrap) *TextItem {
	c := th.Control()
	c.Wrapping = v
	c.Refresh()
	return th
}

func (th *TextItem) Bind(ed *Editor) *TextItem {
	th.ed = ed
	ed.Form.Append(th.prop.Title, th.Control())
	return th
}

func (th *TextItem) Watch() {
	if th.ed == nil {
		panic("not bind editor")
	}
	th.ed.Watch(th.Reload)
}

func (th *TextItem) Reload() {
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
