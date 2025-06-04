package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("GoApp")

	content := container.NewMax(widget.NewLabel("Home"))

	homeBtn := widget.NewButton("Home", func() {
		content.Objects = []fyne.CanvasObject{widget.NewLabel("Home")}
		content.Refresh()
	})

	settingsBtn := widget.NewButton("Settings", func() {
		content.Objects = []fyne.CanvasObject{widget.NewLabel("Settings")}
		content.Refresh()
	})

	nav := container.NewVBox(homeBtn, settingsBtn)

	split := container.NewHSplit(nav, content)
	split.Offset = 0.2

	w.SetContent(split)
	w.Resize(fyne.NewSize(600, 400))
	w.ShowAndRun()
}
