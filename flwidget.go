package fyne_widget

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type FlWidget interface {
	fyne.Widget

	PerformBuild() fyne.CanvasObject

	PerformLayout(box BoxConstraints) fyne.Size
}

type BaseFlWidget struct {
	widget.BaseWidget

	content fyne.CanvasObject
	impl    FlWidget
}

func (this *BaseFlWidget) ExtendBaseFlWidget(w FlWidget) {
	this.BaseWidget.ExtendBaseWidget(w)
	if this.impl != nil {
		return
	}
	widget.NewToolbarSpacer()
	this.impl = w
}

func (this *BaseFlWidget) PerformLayout(box BoxConstraints) fyne.Size {
	if this.content != nil {
		var sz fyne.Size
		if flw, ok := this.content.(FlWidget); ok {
			sz = flw.PerformLayout(box)
		} else {
			if box.HasBoundedWidth() && box.HasBoundedHeight() {
				sz = box.MaxSize()
			} else {
				sz = box.MinSize()
			}
		}
		this.content.Resize(sz)
		return sz
	}
	return fyne.NewSize(0, 0)
}

// CreateRenderer implements fyne.Widget.
func (this *BaseFlWidget) CreateRenderer() fyne.WidgetRenderer {
	o := this.impl
	if this.content == nil {
		this.content = o.PerformBuild()
	}
	return newFlWidgetRenderer(o, this.content)
}

var _ fyne.WidgetRenderer = (*FlWidgetRenderer)(nil)

type FlWidgetRenderer struct {
	self    FlWidget
	objects []fyne.CanvasObject
}

func newFlWidgetRenderer(self FlWidget, o fyne.CanvasObject) *FlWidgetRenderer {
	return &FlWidgetRenderer{
		self:    self,
		objects: []fyne.CanvasObject{o},
	}
}

func (r *FlWidgetRenderer) Destroy() {
}

func (r *FlWidgetRenderer) Layout(s fyne.Size) {
	ns := r.self.PerformLayout(LooseBoxConstraints(s))
	r.objects[0].Resize(ns)
}

func (r *FlWidgetRenderer) MinSize() fyne.Size {
	return r.objects[0].MinSize()
}

func (r *FlWidgetRenderer) Objects() []fyne.CanvasObject {
	return r.objects
}

func (r *FlWidgetRenderer) Refresh() {
	r.objects[0].Refresh()
}
