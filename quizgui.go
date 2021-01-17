package main

import (
	"fmt"
	"image/color"
	"time"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"fyne.io/fyne/layout"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Canvas")
	myWindow.SetContent(widget.NewLabel("Hello"))

	//newwindow := widget.NewLabel("Open canvas window")
	nwWidget := widget.NewVBox(
		widget.NewButton("Open Canvas", func() {
			showAnother(myApp)
		}),
	)

	//newcontainer := widget.NewLabel("Open Container window")
	ncWidget := widget.NewVBox(
		widget.NewButton("Open Container", func() {
			showContainer(myApp)
		}),
	)

	nameLabel := widget.NewLabel("Name: ")
	entryWidget := widget.NewEntry()

	container := fyne.NewContainerWithLayout(layout.NewGridLayout(2),
		nwWidget, ncWidget, nameLabel, entryWidget)

	myWindow.SetContent(container)

	myWindow.ShowAndRun()
	tidyUp()
}

func tidyUp() {
	fmt.Println("Exited")
}

func showAnother(a fyne.App) {

	win := a.NewWindow("Second Window")
	win.SetContent(widget.NewLabel("New Window"))
	myCanvas := win.Canvas()

	text := canvas.NewText("Canvas INC!", color.White)
	text.TextStyle.Bold = true
	myCanvas.SetContent(text)
	go changeContent(myCanvas)

	win.Resize(fyne.NewSize(200, 200))
	win.Show()
}

func changeContent(c fyne.Canvas) {
	time.Sleep(time.Second * 2)

	c.SetContent(canvas.NewRectangle(color.Black))

	time.Sleep(time.Second * 2)
	c.SetContent(canvas.NewLine(color.Gray{0x66}))

	time.Sleep(time.Second * 2)
	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 4
	circle.StrokeColor = color.RGBA{0xff, 0x33, 0x33, 0xff}
	c.SetContent(circle)

	time.Sleep(time.Second * 2)
	c.SetContent(canvas.NewImageFromResource(theme.FyneLogo()))
}

func showContainer(a fyne.App) {

	myWindow := a.NewWindow("Container Demo")

	text1 := canvas.NewText("Hello", color.White)
	text2 := canvas.NewText("There", color.White)
	text2.Move(fyne.NewPos(20,20))
	text3 := canvas.NewText("World", color.White)
	container := fyne.NewContainerWithLayout(layout.NewGridLayout(2), text1, text2, text3)

	myWindow.SetContent(container)

	myWindow.Show()
}
