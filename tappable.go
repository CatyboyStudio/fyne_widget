package fyne_widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
)

var _ fyne.Tappable = (*TappedWith)(nil)
var _ fyne.Widget = (*TappedWith)(nil)
var _ desktop.Cursorable = (*TappedWith)(nil)

type TappedWith struct {
	widget.BaseWidget
	content  fyne.CanvasObject
	OnTapped func()
}

func NewTappedWith(co fyne.CanvasObject, tap func()) *TappedWith {
	o := &TappedWith{
		content:  co,
		OnTapped: tap,
	}
	return o
}

// Cursor implements desktop.Cursorable.
func (*TappedWith) Cursor() desktop.Cursor {
	return desktop.PointerCursor
}

// CreateRenderer implements fyne.Widget.
func (this *TappedWith) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(this.content)
}

// Tapped implements fyne.Tappable.
func (this *TappedWith) Tapped(ev *fyne.PointEvent) {
	if this.OnTapped != nil {
		this.OnTapped()
	}
}
