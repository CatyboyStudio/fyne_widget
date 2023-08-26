package main

import (
	"fyne_widget"
	"fyne_widget/inspector"
	"goapp_fyne"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func (this *DemoWindow) build_Tab_Inspector() fyne.CanvasObject {
	insp := inspector.NewInspector()
	co := this.build_Tab_Inspector_Content(insp)
	split := container.NewHSplit(co, insp)
	split.Offset = 0.7
	return split
}

type Test1 struct {
	A int
	B string
	C bool
}

func (this *DemoWindow) build_Tab_Inspector_Content(insp *inspector.Inspector) fyne.CanvasObject {
	l := container.NewGridWrap(fyne.NewSize(200, 100))
	c1 := canvas.NewRectangle(goapp_fyne.StrToColor("red"))
	o1 := fyne_widget.NewTappedWith(c1, func() {
		insp.Bind(&Test1{}, "")
	})
	o2 := fyne_widget.NewTappedWith(canvas.NewRectangle(goapp_fyne.StrToColor("green")), func() {
		// insp.Bind(Test2{})
	})
	olist := []fyne.CanvasObject{
		o1, o2,
	}
	for _, o := range olist {
		l.Add(container.NewMax(container.NewPadded(o)))
	}
	return l
}
