package inspector

import (
	"log/slog"
	"path/filepath"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

type FilePathItem struct {
	prop    *Property
	content fyne.CanvasObject
	control *widget.Entry
	button  *widget.Button
	ed      *Editor
}

func NewFilePathItem(p *Property) *FilePathItem {
	return &FilePathItem{
		prop: p,
	}
}

func (th *FilePathItem) Content() fyne.CanvasObject {
	if th.content == nil {
		th.control = widget.NewEntry()
		th.control.OnChanged = th.onSetValue
		th.button = widget.NewButtonWithIcon("", theme.FileIcon(), th.onFileDialog)
		th.content = container.NewBorder(nil, nil, nil, th.button, th.control)
		th.Reload()
	}
	return th.content
}

func (th *FilePathItem) onSetValue(s string) {
	th.ed.Execute(func() {
		err := th.prop.SetString(s)
		if err != nil {
			slog.Error("SetString fail", "property", th.prop.Title, "error", err)
		} else {
			th.prop.OnUpdate()
		}
	})
}

func (th *FilePathItem) onFileDialog() {
	dlg := dialog.NewFileSavePath(func(uc fyne.URI, err error) {
		if err != nil {
			slog.Error("Filepath FileSaveDialog fail", "error", err)
			return
		}
		if uc == nil {
			return
		}

		fn := uc.String()
		if strings.HasPrefix(fn, "file://") {
			fn = strings.TrimPrefix(fn, "file://")
			th.control.Text = fn
			th.control.Refresh()
			th.onSetValue(fn)
		}
	}, th.ed.Inspector.Window)
	fn, _ := th.prop.GetString()
	if fn != "" {
		bfn := filepath.Base(fn)
		dlg.SetFileName(bfn)
	}
	dlg.SetFilter(storage.NewExtensionFileFilter([]string{".go", ".txt"}))
	dlg.Show()
}

func (th *FilePathItem) Validator(v fyne.StringValidator) *FilePathItem {
	th.Content()
	th.control.Validator = v
	return th
}

func (th *FilePathItem) Bind(ed *Editor) *FilePathItem {
	th.ed = ed
	ed.Form.Append(th.prop.Title, th.Content())
	return th
}

func (th *FilePathItem) Watch() {
	if th.ed == nil {
		panic("not bind editor")
	}
	th.ed.Watch(th.Reload)
}

func (th *FilePathItem) Reload() {
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
