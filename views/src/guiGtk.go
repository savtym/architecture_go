package src

import (
	"../ui"
)

var HandEditText *ui.Entry
var DecodeBtn *ui.Button
var HandOutputTextView *ui.Label
var ChooseFile *ui.Button
var FileOutputTextView *ui.Label

func ShowWindow() {
	err := ui.Main(func() {
		mainBox := ui.NewHorizontalBox()
		mainBox.SetPadded(true)
		mainBox.Append(makeHandInputBlock(), true)
		mainBox.Append(makeFromFileBlock(), true)

		window := ui.NewWindow("Morse decoder", 700, 300, false)
		window.SetChild(mainBox)
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}

func makeHandInputBlock() *ui.Group {
	help := ui.NewLabel("Please, enter the morse code in field below")
	HandEditText = ui.NewEntry()
	HandOutputTextView = ui.NewLabel("Here will be the decoded text.")
	DecodeBtn = ui.NewButton("Decode")

	leftBox := ui.NewVerticalBox()
	leftBox.SetPadded(true)
	leftBox.Append(help, false)
	leftBox.Append(HandEditText, false)
	leftBox.Append(DecodeBtn, false)
	leftBox.Append(ui.NewHorizontalSeparator(), false)
	leftBox.Append(HandOutputTextView, true)

	leftGroup := ui.NewGroup("Hand HandEditText")
	leftGroup.SetMargined(true)
	leftGroup.SetChild(leftBox)
	return  leftGroup
}

func makeFromFileBlock() *ui.Group {
	ChooseFile = ui.NewButton("Choose file")
	help := ui.NewLabel("Please, choose the .txt file with morse code")
	FileOutputTextView = ui.NewLabel("Here will be the decoded text.")

	rightBox := ui.NewVerticalBox()
	rightBox.SetPadded(true)
	rightBox.Append(help, false)
	rightBox.Append(ChooseFile, false)
	rightBox.Append(ui.NewHorizontalSeparator(), false)
	rightBox.Append(FileOutputTextView, true)

	rightGroup := ui.NewGroup("Open from file")
	rightGroup.SetMargined(true)
	rightGroup.SetChild(rightBox)
	return rightGroup
}
