package inspector

import "fyne.io/fyne/v2"

type Editor interface {
	IsPanel() bool

	CreateInspectorGUI() fyne.CanvasObject
}
