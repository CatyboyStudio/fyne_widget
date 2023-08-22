package inspector

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

var _ (fyne.Widget) = (*Inspector)(nil)

type Inspector struct {
	widget.BaseWidget

	data any
}

func NewInspector(data any) *Inspector {
	o := &Inspector{
		data: data,
	}
	o.ExtendBaseWidget(o)
	return o
}

func (*Inspector) CreateRenderer() fyne.WidgetRenderer {
	co := widget.NewForm()
	return widget.NewSimpleRenderer(co)
}

func (this *Inspector) Bind(data any) error {
	return nil
}
