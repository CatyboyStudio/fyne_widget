package main

import (
	"fyne_widget"
	"goapp_commons"
	"goapp_fyne"
	"os"

	"fyne.io/fyne/v2/app"
)

func main() {
	goapp_commons.Init("config.toml", "log.toml")

	goapp_fyne.InitFont()
	defer os.Unsetenv("FYNE_FONT")

	os.Setenv("FYNE_THEME", "dark")
	myApp := app.NewWithID("FyneDesigner.CatyboyStudio")

	goapp_commons.SupportLangs = append(goapp_commons.SupportLangs,
		// goapp_commons.NewLangInfo("zh-CHS"),
		goapp_commons.NewLangInfo("zh"),
	)
	goapp_commons.InitI18N("i18n")
	fyne_widget.GetI18n = goapp_commons.GetMessage

	mainWindow := NewDemoWindow()
	mainWindow.Show()

	myApp.Run()
}
