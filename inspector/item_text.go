package inspector

import (
	"log/slog"

	"fyne.io/fyne/v2/widget"
)

type TextItem struct {
	prop *Property
	text *widget.Label
	ed   *Editor
}

func NewTextItem(p *Property) *TextItem {
	return &TextItem{
		prop: p,
	}
}

func (th *TextItem) Bind(ed *Editor) *TextItem {
	th.ed = ed
	if th.text == nil {
		th.text = widget.NewLabel("")
		th.Reload()
	}
	ed.Form.Append(th.prop.Title, th.text)
	return th
}

func (th *TextItem) Watch() {
	if th.ed == nil {
		panic("editor is nil")
	}
	th.ed.Watch(th.Reload)
}

func (th *TextItem) Reload() {
	if th.text == nil {
		return
	}
	s, err := th.prop.GetString()
	if err != nil {
		slog.Error("bind fail", "property", th.prop.Title, "error", err)
		return
	}
	th.text.Text = s
	th.text.Refresh()
}
