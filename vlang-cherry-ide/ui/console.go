package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Console struct {
	Container fyne.CanvasObject // CORREGIDO: Cambié *container.Border por fyne.CanvasObject
	Output    *widget.Entry
}

func NewConsole() *Console {
	// Crear área de salida
	output := widget.NewMultiLineEntry()
	output.SetPlaceHolder("La salida del programa aparecerá aquí...")
	output.Disable() // Solo lectura

	// Crear etiqueta
	label := widget.NewLabel("Consola")

	console := &Console{
		Output: output,
	}

	// Crear layout
	scrollOutput := container.NewScroll(output)
	console.Container = container.NewBorder(
		label,        // top
		nil,          // bottom
		nil,          // left
		nil,          // right
		scrollOutput, // center
	)

	return console
}

func (c *Console) AddOutput(text string) {
	current := c.Output.Text
	c.Output.SetText(current + text)
}

func (c *Console) ClearOutput() {
	c.Output.SetText("")
}
