package inspector

import (
	"cbsutil/executor"
	"log/slog"
	"slices"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type inspItem struct {
	id     int
	label  string
	editor *Editor
}

func (o inspItem) IsNil() bool {
	return o.editor == nil
}

var _ (fyne.Widget) = (*Inspector)(nil)

type Inspector struct {
	widget.BaseWidget

	bindid int
	items  []inspItem

	form       *widget.Form
	toolbar    *fyne.Container
	backButton *widget.Button
	pathText   *widget.Label

	Executor executor.Executor[any, any]
}

func NewInspector() *Inspector {
	o := &Inspector{
		Executor: executor.NewInline[any, any](),
	}
	o.ExtendBaseWidget(o)
	return o
}

func (th *Inspector) onBack() {

}

func (th *Inspector) CreateRenderer() fyne.WidgetRenderer {
	th.backButton = widget.NewButtonWithIcon("", theme.NavigateBackIcon(), th.onBack)
	rbt := widget.NewButtonWithIcon("", theme.ViewRefreshIcon(), func() {
		th.showCurrent()
	})
	th.pathText = widget.NewLabel("")
	th.pathText.Wrapping = fyne.TextTruncate
	c1 := container.NewHBox(th.backButton, rbt)
	c11 := container.NewBorder(nil, nil, c1, nil, th.pathText)
	c2 := container.NewPadded(c11)
	th.toolbar = c2
	th.toolbar.Hide()
	th.form = widget.NewForm()
	co := container.NewVBox(c2, th.form)
	return widget.NewSimpleRenderer(co)
}

func (th *Inspector) current() inspItem {
	if len(th.items) > 0 {
		return th.items[len(th.items)-1]
	}
	return inspItem{}
}

func (th *Inspector) pushItem(item inspItem) {
	th.items = append(th.items, item)
}

func (th *Inspector) popItem() {
	c := len(th.items)
	if c > 0 {
		th.items[c-1].editor.close()
		th.items = slices.Delete(th.items, c-1, c)
	}
}

func (th *Inspector) showCurrent() {
	th.form.Items = nil
	th.form.Refresh()

	item := th.current()
	if item.IsNil() {
		th.toolbar.Hide()
	} else {
		item.editor.Form = th.form
		err := item.editor.builder.BuildEditor(item.editor)
		if err != nil {
			slog.Error("Build editor fail", "editor", item.editor, "data", item.editor.Data, "error", err)
		}
		th.toolbar.Show()
		if len(th.items) > 1 {
			th.backButton.Enable()
		} else {
			th.backButton.Disable()
		}
	}
	th.form.Refresh()
}

func (th *Inspector) Bind(data any, editorType string) (int, error) {
	ed, err := CreateEditor(th, data, editorType)
	if err != nil {
		return 0, err
	}
	th.bindid++
	id := th.bindid
	item := inspItem{
		id:     id,
		editor: ed,
	}
	c := len(th.items)
	for i := 0; i < c; i++ {
		th.popItem()
	}
	th.pushItem(item)
	th.showCurrent()
	return id, nil
}

func (th *Inspector) Unbind(id int) {
	idx := -1
	for i, v := range th.items {
		if v.id == id {
			idx = i
			break
		}
	}
	if idx == -1 {
		return
	}
	for i := len(th.items) - 1; i >= idx; i-- {
		th.popItem()
	}
	th.showCurrent()
}

func (th *Inspector) Push(label string, data any, editorType string) error {
	ed, err := CreateEditor(th, data, editorType)
	if err != nil {
		return err
	}
	th.bindid++
	id := th.bindid
	item := inspItem{
		id:     id,
		label:  label,
		editor: ed,
	}
	th.pushItem(item)
	th.showCurrent()
	return nil
}
