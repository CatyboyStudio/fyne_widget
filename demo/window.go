package main

import (
	"fyne_widget"
	goapp_fyne "fyne_widgets"
	"goapp_commons"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type DemoWindow struct {
	win fyne.Window
}

func NewDemoWindow() *DemoWindow {
	o := &DemoWindow{
		win: fyne.CurrentApp().NewWindow(goapp_commons.GetMessage("MainWindow.Title")),
	}

	return o
}

func (this *DemoWindow) Show() {
	this.win.Resize(fyne.NewSize(720, 480))
	this.win.SetContent(this.build_Content())
	goapp_fyne.ShowMaximizeWindow(this.win)
}

func (this *DemoWindow) build_Content() fyne.CanvasObject {

	tabs := container.NewAppTabs(
		container.NewTabItem(fyne_widget.GetI18n("MainWindow.Tab.Layout"), this.build_Tab_Layout()),
		container.NewTabItem(fyne_widget.GetI18n("MainWindow.Tab.Inspector"), this.build_Tab_Inspector()),
	)
	// tabs.SetTabLocation(container.TabLocationLeading)
	tabs.SelectIndex(1)
	return tabs
}

func (this *DemoWindow) build_Tab_Layout() fyne.CanvasObject {
	return widget.NewLabel("Test")
}
