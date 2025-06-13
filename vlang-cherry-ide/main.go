package main

import (
	"main/ui"

	"fyne.io/fyne/v2/app"
)

func main() {
	// Crear aplicación
	myApp := app.New()
	myApp.SetIcon(nil) // Puedes agregar un icono más tarde

	// Crear y mostrar ventana principal
	mainWindow := ui.NewMainWindow(myApp)
	mainWindow.ShowAndRun()
}
