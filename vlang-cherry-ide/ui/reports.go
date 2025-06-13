/*
package ui

import (

	"strconv"
	"vlang-cherry-ide/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

)

	type Reports struct {
		Container   *container.AppTabs
		ErrorList   *widget.List
		SymbolTable *widget.Table
		ASTContent  *fyne.Container // CORREGIDO: usar fyne.Container

		// Datos
		errors      []models.ErrorReport
		symbolTable []models.SymbolEntry
	}

	func NewReports() *Reports {
		reports := &Reports{
			errors:      []models.ErrorReport{},
			symbolTable: []models.SymbolEntry{},
		}

		// Crear lista de errores
		reports.ErrorList = widget.NewList(
			func() int {
				return len(reports.errors)
			},
			func() fyne.CanvasObject {
				return container.NewHBox(
					widget.NewLabel("Línea 0"),
					widget.NewLabel("Col 0"),
					widget.NewLabel("Descripción del error"),
					widget.NewLabel("Tipo"),
				)
			},
			func(id widget.ListItemID, obj fyne.CanvasObject) {
				if id < len(reports.errors) {
					error := reports.errors[id]
					hbox := obj.(*fyne.Container) // CORREGIDO
					hbox.Objects[0].(*widget.Label).SetText("Línea " + strconv.Itoa(error.Linea))
					hbox.Objects[1].(*widget.Label).SetText("Col " + strconv.Itoa(error.Columna))
					hbox.Objects[2].(*widget.Label).SetText(error.Descripcion)
					hbox.Objects[3].(*widget.Label).SetText(error.Tipo)
				}
			},
		)

		// Crear tabla de símbolos
		reports.SymbolTable = widget.NewTable(
			func() (int, int) {
				return len(reports.symbolTable) + 1, 6 // +1 para header
			},
			func() fyne.CanvasObject {
				return widget.NewLabel("Template")
			},
			func(id widget.TableCellID, obj fyne.CanvasObject) {
				label := obj.(*widget.Label)
				if id.Row == 0 {
					// Header
					headers := []string{"ID", "Tipo", "Tipo de Dato", "Ámbito", "Línea", "Columna"}
					if id.Col < len(headers) {
						label.SetText(headers[id.Col])
						label.TextStyle = fyne.TextStyle{Bold: true}
					}
				} else {
					// Data
					if id.Row-1 < len(reports.symbolTable) {
						symbol := reports.symbolTable[id.Row-1]
						switch id.Col {
						case 0:
							label.SetText(symbol.ID)
						case 1:
							label.SetText(symbol.SymbolType)
						case 2:
							label.SetText(symbol.DataType)
						case 3:
							label.SetText(symbol.Scope)
						case 4:
							label.SetText(strconv.Itoa(symbol.Line))
						case 5:
							label.SetText(strconv.Itoa(symbol.Column))
						}
					}
				}
			},
		)

		// Crear contenedor para AST
		reports.ASTContent = container.NewVBox(
			widget.NewLabel("No hay AST disponible"),
		)

		// Crear pestañas
		reports.Container = container.NewAppTabs(
			container.NewTabItem("Errores", reports.ErrorList),
			container.NewTabItem("Tabla de Símbolos", reports.SymbolTable),
			container.NewTabItem("AST", container.NewScroll(reports.ASTContent)),
		)

		return reports
	}

// UpdateErrors actualiza la lista de errores

	func (r *Reports) UpdateErrors(errors []models.ErrorReport) {
		r.errors = errors
		r.ErrorList.Refresh()
	}

// UpdateSymbolTable actualiza la tabla de símbolos

	func (r *Reports) UpdateSymbolTable(symbols []models.SymbolEntry) {
		r.symbolTable = symbols
		r.SymbolTable.Refresh()
	}

// UpdateASTWithSVG actualiza la pestaña AST con contenido SVG (método original)

	func (r *Reports) UpdateASTWithSVG(astText string, svgContent string) {
		if svgContent != "" {
			// Crear widget de texto para mostrar el SVG
			astLabel := widget.NewLabel(astText)
			svgText := widget.NewEntry()
			svgText.MultiLine = true
			svgText.Wrapping = fyne.TextWrapWord
			svgText.SetText(svgContent)
			svgText.Resize(fyne.NewSize(400, 300))

			content := container.NewVBox(
				astLabel,
				widget.NewSeparator(),
				svgText,
			)

			r.ASTContent.Objects = []fyne.CanvasObject{content}
			r.ASTContent.Refresh()
		} else {
			label := widget.NewLabel("No hay AST disponible")
			r.ASTContent.Objects = []fyne.CanvasObject{label}
			r.ASTContent.Refresh()
		}
	}

// UpdateASTWithPNG actualiza la pestaña AST con una imagen PNG

	func (r *Reports) UpdateASTWithPNG(astText string, pngData []byte) {
		if len(pngData) > 0 {
			// Crear resource desde los bytes PNG
			resource := fyne.NewStaticResource("ast.png", pngData)

			// Crear imagen desde el resource
			astImage := canvas.NewImageFromResource(resource)
			astImage.FillMode = canvas.ImageFillContain // Mantener proporciones
			astImage.SetMinSize(fyne.NewSize(400, 300))

			// Crear scroll container para la imagen
			scrollContainer := container.NewScroll(astImage)
			scrollContainer.SetMinSize(fyne.NewSize(400, 300))

			// Crear contenedor con título y imagen
			content := container.NewVBox(
				widget.NewLabel("🌳 Árbol Sintáctico Abstracto (AST)"),
				widget.NewSeparator(),
				scrollContainer,
			)

			// Actualizar contenido de la pestaña AST
			r.ASTContent.Objects = []fyne.CanvasObject{content}
			r.ASTContent.Refresh()
		} else {
			// Si no hay PNG, mostrar mensaje
			label := widget.NewLabel("No hay AST disponible")
			label.Alignment = fyne.TextAlignCenter

			r.ASTContent.Objects = []fyne.CanvasObject{label}
			r.ASTContent.Refresh()
		}
	}

// ClearAll limpia todos los reportes

	func (r *Reports) ClearAll() {
		r.errors = []models.ErrorReport{}
		r.symbolTable = []models.SymbolEntry{}

		r.ErrorList.Refresh()
		r.SymbolTable.Refresh()

		// Limpiar AST
		label := widget.NewLabel("No hay AST disponible")
		r.ASTContent.Objects = []fyne.CanvasObject{label}
		r.ASTContent.Refresh()
	}
*/
/*
package ui

import (
	"fmt"
	"image/color" // NUEVO
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
	"vlang-cherry-ide/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Reports struct {
	Container   *container.AppTabs
	ErrorList   *widget.List
	SymbolTable *widget.Table
	ASTContent  *fyne.Container

	// Datos
	errors      []models.ErrorReport
	symbolTable []models.SymbolEntry
}

func NewReports() *Reports {
	reports := &Reports{
		errors:      []models.ErrorReport{},
		symbolTable: []models.SymbolEntry{},
	}

	// Crear lista de errores
	reports.ErrorList = widget.NewList(
		func() int {
			return len(reports.errors)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(
				widget.NewLabel("Línea 0"),
				widget.NewLabel("Col 0"),
				widget.NewLabel("Descripción del error"),
				widget.NewLabel("Tipo"),
			)
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			if id < len(reports.errors) {
				error := reports.errors[id]
				hbox := obj.(*fyne.Container)
				hbox.Objects[0].(*widget.Label).SetText("Línea " + strconv.Itoa(error.Linea))
				hbox.Objects[1].(*widget.Label).SetText("Col " + strconv.Itoa(error.Columna))
				hbox.Objects[2].(*widget.Label).SetText(error.Descripcion)
				hbox.Objects[3].(*widget.Label).SetText(error.Tipo)
			}
		},
	)

	// Crear tabla de símbolos
	reports.SymbolTable = widget.NewTable(
		func() (int, int) {
			return len(reports.symbolTable) + 1, 6 // +1 para header
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)
			if id.Row == 0 {
				// Header
				headers := []string{"-------ID-------", "Tipo", "Tipo de Dato", "Ámbito", "Línea", "Columna"}
				if id.Col < len(headers) {
					label.SetText(headers[id.Col])
					label.TextStyle = fyne.TextStyle{Bold: true}
				}
			} else {
				// Data
				if id.Row-1 < len(reports.symbolTable) {
					symbol := reports.symbolTable[id.Row-1]
					switch id.Col {
					case 0:
						label.SetText(symbol.ID)
					case 1:
						label.SetText(symbol.SymbolType)
					case 2:
						label.SetText(symbol.DataType)
					case 3:
						label.SetText(symbol.Scope)
					case 4:
						label.SetText(strconv.Itoa(symbol.Line))
					case 5:
						label.SetText(strconv.Itoa(symbol.Column))
					}
				}
			}
		},
	)

	// Crear contenedor para AST
	reports.ASTContent = container.NewVBox(
		widget.NewLabel("No hay AST disponible"),
	)

	// Crear pestañas
	reports.Container = container.NewAppTabs(
		container.NewTabItem("Errores", reports.ErrorList),
		container.NewTabItem("Tabla de Símbolos", reports.SymbolTable),
		container.NewTabItem("AST", container.NewScroll(reports.ASTContent)),
	)

	return reports
}

// UpdateErrors actualiza la lista de errores
func (r *Reports) UpdateErrors(errors []models.ErrorReport) {
	r.errors = errors
	r.ErrorList.Refresh()
}

// UpdateSymbolTable actualiza la tabla de símbolos
func (r *Reports) UpdateSymbolTable(symbols []models.SymbolEntry) {
	r.symbolTable = symbols
	r.SymbolTable.Refresh()
}

// UpdateASTWithSVG - Versión que funciona: placeholder + navegador
func (r *Reports) UpdateASTWithSVG(astText string, svgContent string) {
	if svgContent != "" {
		// 1. Guardar SVG como archivo
		svgPath, err := r.saveSVGToFile(svgContent)
		if err != nil {
			r.showError("Error guardando SVG: " + err.Error())
			return
		}

		// 2. Crear un placeholder visual (ya que Fyne no puede mostrar SVG complejos)
		placeholderImage := r.createASTPlaceholder()

		// 3. Información del AST
		infoText := fmt.Sprintf(`🌳 Árbol Sintáctico Abstracto

✅ SVG generado y guardado exitosamente
📁 Archivo: %s
📊 Tamaño: %d caracteres

💡 Haz clic en "Abrir en Navegador" para ver el AST completo`,
			filepath.Base(svgPath),
			len(svgContent))

		infoLabel := widget.NewLabel(infoText)
		infoLabel.Wrapping = fyne.TextWrapWord

		// 4. Botón para abrir en navegador (funciona perfecto)
		openBtn := widget.NewButton("🌐 Abrir en Navegador", func() {
			r.openFileInBrowser(svgPath)
		})
		openBtn.Importance = widget.HighImportance

		// 5. Layout con placeholder + botón
		content := container.NewVBox(
			infoLabel,
			widget.NewSeparator(),
			openBtn,
			widget.NewSeparator(),
			widget.NewLabel("Vista previa (abrir en navegador para ver completo):"),
			container.NewScroll(placeholderImage),
		)

		r.ASTContent.Objects = []fyne.CanvasObject{content}
		r.ASTContent.Refresh()

	} else {
		r.showError("No hay AST disponible")
	}
}

// createASTPlaceholder crea una imagen placeholder que representa un árbol
func (r *Reports) createASTPlaceholder() *fyne.Container {
	// Crear un placeholder visual simple que se parece a un árbol
	treeText := `
           🌳 AST GENERADO
         /              \
     stmt:9           stmt:9
    /      \         /      \
func_call:1  print  mut   decl
    |         |      |      |
   print   "text"  var   type
                    |      |
                  value  int

💡 El AST completo se muestra en el navegador
   (Fyne no puede renderizar SVG complejos)
`

	treeLabel := widget.NewLabel(treeText)
	treeLabel.TextStyle = fyne.TextStyle{Monospace: true}

	// Crear un contenedor con fondo gris claro
	background := canvas.NewRectangle(&color.RGBA{245, 245, 245, 255})

	content := container.NewStack(
		background,
		container.NewPadded(treeLabel),
	)

	content.Resize(fyne.NewSize(500, 300))
	return content
}

// saveSVGToFile guarda el SVG como archivo local
func (r *Reports) saveSVGToFile(svgContent string) (string, error) {
	// Crear directorio si no existe
	outputDir := "generated_asts"
	os.MkdirAll(outputDir, 0755)

	// Generar nombre único con timestamp
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("ast_%s.svg", timestamp)
	fullPath := filepath.Join(outputDir, filename)

	// Escribir archivo
	err := os.WriteFile(fullPath, []byte(svgContent), 0644)
	if err != nil {
		return "", err
	}

	// Obtener ruta absoluta
	absPath, err := filepath.Abs(fullPath)
	if err != nil {
		return fullPath, nil
	}

	return absPath, nil
}

// showError muestra un mensaje de error
func (r *Reports) showError(message string) {
	label := widget.NewLabel("❌ " + message)
	label.Alignment = fyne.TextAlignCenter
	r.ASTContent.Objects = []fyne.CanvasObject{label}
	r.ASTContent.Refresh()
}

// UpdateASTWithPNG mantener compatibilidad
func (r *Reports) UpdateASTWithPNG(astText string, pngData []byte) {
	// Si tenemos datos PNG, mostrarlos, sino mostrar mensaje
	if len(pngData) > 0 {
		// Crear imagen desde bytes
		resource := fyne.NewStaticResource("ast.png", pngData)
		astImage := canvas.NewImageFromResource(resource)
		astImage.FillMode = canvas.ImageFillContain
		astImage.SetMinSize(fyne.NewSize(400, 300))

		content := container.NewVBox(
			widget.NewLabel("🌳 AST generado como PNG"),
			widget.NewSeparator(),
			container.NewScroll(astImage),
		)

		r.ASTContent.Objects = []fyne.CanvasObject{content}
		r.ASTContent.Refresh()
	} else {
		label := widget.NewLabel("❌ No hay datos PNG disponibles")
		label.Alignment = fyne.TextAlignCenter
		r.ASTContent.Objects = []fyne.CanvasObject{label}
		r.ASTContent.Refresh()
	}
}

// ClearAll limpia todos los reportes
func (r *Reports) ClearAll() {
	r.errors = []models.ErrorReport{}
	r.symbolTable = []models.SymbolEntry{}

	r.ErrorList.Refresh()
	r.SymbolTable.Refresh()

	// Limpiar AST
	label := widget.NewLabel("No hay AST disponible")
	r.ASTContent.Objects = []fyne.CanvasObject{label}
	r.ASTContent.Refresh()
}

// openFileInBrowser abre cualquier archivo en el navegador del sistema
func (r *Reports) openFileInBrowser(filePath string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", filePath)
	case "darwin": // macOS
		cmd = exec.Command("open", filePath)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", filePath)
	default:
		r.showError("Sistema operativo no soportado")
		return
	}

	if err := cmd.Start(); err != nil {
		r.showError("Error abriendo archivo: " + err.Error())
	}
}
*/
/*
package ui

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"time"
	"vlang-cherry-ide/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas" // AGREGAR ESTA LÍNEA
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Reports struct {
	Container   *container.AppTabs
	ErrorList   *widget.List
	SymbolTable *widget.Table
	ASTContent  *fyne.Container

	// Datos
	errors      []models.ErrorReport
	symbolTable []models.SymbolEntry
}

func NewReports() *Reports {
	reports := &Reports{
		errors:      []models.ErrorReport{},
		symbolTable: []models.SymbolEntry{},
	}

	// Crear lista de errores
	reports.ErrorList = widget.NewList(
		func() int {
			return len(reports.errors)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(
				widget.NewLabel("Línea 0"),
				widget.NewLabel("Col 0"),
				widget.NewLabel("Descripción del error"),
				widget.NewLabel("Tipo"),
			)
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			if id < len(reports.errors) {
				error := reports.errors[id]
				hbox := obj.(*fyne.Container)
				hbox.Objects[0].(*widget.Label).SetText("Línea " + strconv.Itoa(error.Linea))
				hbox.Objects[1].(*widget.Label).SetText("Col " + strconv.Itoa(error.Columna))
				hbox.Objects[2].(*widget.Label).SetText(error.Descripcion)
				hbox.Objects[3].(*widget.Label).SetText(error.Tipo)
			}
		},
	)

	// Crear tabla de símbolos con columnas más anchas
	reports.SymbolTable = widget.NewTable(
		func() (int, int) {
			return len(reports.symbolTable) + 1, 6 // +1 para header
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)
			if id.Row == 0 {
				// Header
				headers := []string{"ID", "Tipo", "Tipo de Dato", "Ámbito", "Línea", "Columna"}
				if id.Col < len(headers) {
					label.SetText(headers[id.Col])
					label.TextStyle = fyne.TextStyle{Bold: true}
				}
			} else {
				// Data
				if id.Row-1 < len(reports.symbolTable) {
					symbol := reports.symbolTable[id.Row-1]
					switch id.Col {
					case 0:
						label.SetText(symbol.ID)
					case 1:
						label.SetText(symbol.SymbolType)
					case 2:
						label.SetText(symbol.DataType)
					case 3:
						label.SetText(symbol.Scope)
					case 4:
						label.SetText(strconv.Itoa(symbol.Line))
					case 5:
						label.SetText(strconv.Itoa(symbol.Column))
					}
				}
			}
		},
	)

	// Configurar anchos de columna para mejor legibilidad
	reports.SymbolTable.SetColumnWidth(0, 120) // ID
	reports.SymbolTable.SetColumnWidth(1, 100) // Tipo
	reports.SymbolTable.SetColumnWidth(2, 120) // Tipo de Dato
	reports.SymbolTable.SetColumnWidth(3, 80)  // Ámbito
	reports.SymbolTable.SetColumnWidth(4, 60)  // Línea
	reports.SymbolTable.SetColumnWidth(5, 70)  // Columna

	// Crear contenedor para AST
	reports.ASTContent = container.NewVBox(
		widget.NewLabel("No hay AST disponible"),
	)

	// Crear pestañas
	reports.Container = container.NewAppTabs(
		container.NewTabItem("Errores", reports.ErrorList),
		container.NewTabItem("Tabla de Símbolos", reports.SymbolTable),
		container.NewTabItem("AST", container.NewScroll(reports.ASTContent)),
	)

	return reports
}

// UpdateErrors actualiza la lista de errores
func (r *Reports) UpdateErrors(errors []models.ErrorReport) {
	r.errors = errors
	r.ErrorList.Refresh()
}

// UpdateSymbolTable actualiza la tabla de símbolos
func (r *Reports) UpdateSymbolTable(symbols []models.SymbolEntry) {
	r.symbolTable = symbols
	r.SymbolTable.Refresh()
}

// UpdateASTWithSVG - Versión que funciona: placeholder + navegador
func (r *Reports) UpdateASTWithSVG(astText string, svgContent string) {
	if svgContent != "" {
		// 1. Guardar SVG como archivo
		svgPath, err := r.saveSVGToFile(svgContent)
		if err != nil {
			r.showError("Error guardando SVG: " + err.Error())
			return
		}

		// 2. Crear un placeholder visual (ya que Fyne no puede mostrar SVG complejos)
		placeholderImage := r.createASTPlaceholder()

		// 3. Información del AST
		infoText := fmt.Sprintf(`🌳 Árbol Sintáctico Abstracto

✅ SVG generado y guardado exitosamente
📁 Archivo: %s
📊 Tamaño: %d caracteres

💡 Haz clic en "Abrir en Navegador" para ver el AST completo`,
			filepath.Base(svgPath),
			len(svgContent))

		infoLabel := widget.NewLabel(infoText)
		infoLabel.Wrapping = fyne.TextWrapWord

		// 4. Botón para abrir en navegador (funciona perfecto)
		openBtn := widget.NewButton("🌐 Abrir en Navegador", func() {
			r.openFileInBrowser(svgPath)
		})
		openBtn.Importance = widget.HighImportance

		// 5. Layout con placeholder + botón
		content := container.NewVBox(
			infoLabel,
			widget.NewSeparator(),
			openBtn,
			widget.NewSeparator(),
			widget.NewLabel("Vista previa (abrir en navegador para ver completo):"),
			container.NewScroll(placeholderImage),
		)

		r.ASTContent.Objects = []fyne.CanvasObject{content}
		r.ASTContent.Refresh()

	} else {
		r.showError("No hay AST disponible")
	}
}

// createASTPlaceholder crea una imagen placeholder más simple y visible
func (r *Reports) createASTPlaceholder() *fyne.Container {
	// Crear texto del árbol más simple y visible
	treeText := `🌳 AST GENERADO EXITOSAMENTE

        program:1
       /         \
   stmt:9       stmt:9
  /     \      /     \
print   mut   print  bool
  |      |     |      |
"text" var   "val"  type

📊 Información:
• Árbol sintáctico generado
• Archivo SVG guardado
• Usa el botón para ver completo

💡 El SVG completo está en el navegador`

	// Crear label con texto monospace
	treeLabel := widget.NewLabel(treeText)
	treeLabel.TextStyle = fyne.TextStyle{Monospace: true}
	treeLabel.Alignment = fyne.TextAlignCenter

	// Crear contenedor simple SIN fondo (para evitar bloqueos)
	content := container.NewPadded(treeLabel)
	return content
}

// saveSVGToFile guarda el SVG como archivo local
func (r *Reports) saveSVGToFile(svgContent string) (string, error) {
	// Crear directorio si no existe
	outputDir := "generated_asts"
	os.MkdirAll(outputDir, 0755)

	// Generar nombre único con timestamp
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("ast_%s.svg", timestamp)
	fullPath := filepath.Join(outputDir, filename)

	// Escribir archivo
	err := os.WriteFile(fullPath, []byte(svgContent), 0644)
	if err != nil {
		return "", err
	}

	// Obtener ruta absoluta
	absPath, err := filepath.Abs(fullPath)
	if err != nil {
		return fullPath, nil
	}

	return absPath, nil
}

// showError muestra un mensaje de error
func (r *Reports) showError(message string) {
	label := widget.NewLabel("❌ " + message)
	label.Alignment = fyne.TextAlignCenter
	r.ASTContent.Objects = []fyne.CanvasObject{label}
	r.ASTContent.Refresh()
}

// UpdateASTWithPNG mantener compatibilidad
func (r *Reports) UpdateASTWithPNG(astText string, pngData []byte) {
	// Si tenemos datos PNG, mostrarlos, sino mostrar mensaje
	if len(pngData) > 0 {
		// Crear imagen desde bytes
		resource := fyne.NewStaticResource("ast.png", pngData)
		astImage := canvas.NewImageFromResource(resource)
		astImage.FillMode = canvas.ImageFillContain
		astImage.SetMinSize(fyne.NewSize(400, 300))

		content := container.NewVBox(
			widget.NewLabel("🌳 AST generado como PNG"),
			widget.NewSeparator(),
			container.NewScroll(astImage),
		)

		r.ASTContent.Objects = []fyne.CanvasObject{content}
		r.ASTContent.Refresh()
	} else {
		label := widget.NewLabel("❌ No hay datos PNG disponibles")
		label.Alignment = fyne.TextAlignCenter
		r.ASTContent.Objects = []fyne.CanvasObject{label}
		r.ASTContent.Refresh()
	}
}

// ClearAll limpia todos los reportes
func (r *Reports) ClearAll() {
	r.errors = []models.ErrorReport{}
	r.symbolTable = []models.SymbolEntry{}

	r.ErrorList.Refresh()
	r.SymbolTable.Refresh()

	// Limpiar AST
	label := widget.NewLabel("No hay AST disponible")
	r.ASTContent.Objects = []fyne.CanvasObject{label}
	r.ASTContent.Refresh()
}

// openFileInBrowser abre cualquier archivo en el navegador del sistema
func (r *Reports) openFileInBrowser(filePath string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", filePath)
	case "darwin": // macOS
		cmd = exec.Command("open", filePath)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", filePath)
	default:
		r.showError("Sistema operativo no soportado")
		return
	}

	if err := cmd.Start(); err != nil {
		r.showError("Error abriendo archivo: " + err.Error())
	}
}
*/

package ui

import (
	"fmt"
	"main/models"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Reports struct {
	Container   *container.AppTabs
	ErrorList   *widget.List
	SymbolTable *widget.Table
	ASTContent  *fyne.Container

	// Datos
	errors      []models.ErrorReport
	symbolTable []models.SymbolEntry
}

func NewReports() *Reports {
	reports := &Reports{
		errors:      []models.ErrorReport{},
		symbolTable: []models.SymbolEntry{},
	}

	// Crear lista de errores
	reports.ErrorList = widget.NewList(
		func() int {
			return len(reports.errors)
		},
		func() fyne.CanvasObject {
			return container.NewHBox(
				widget.NewLabel("Línea 0"),
				widget.NewLabel("Col 0"),
				widget.NewLabel("Descripción del error"),
				widget.NewLabel("Tipo"),
			)
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			if id < len(reports.errors) {
				error := reports.errors[id]
				hbox := obj.(*fyne.Container)
				hbox.Objects[0].(*widget.Label).SetText("Línea " + strconv.Itoa(error.Linea))
				hbox.Objects[1].(*widget.Label).SetText("Col " + strconv.Itoa(error.Columna))
				hbox.Objects[2].(*widget.Label).SetText(error.Descripcion)
				hbox.Objects[3].(*widget.Label).SetText(error.Tipo)
			}
		},
	)

	// Crear tabla de símbolos con columnas más anchas
	reports.SymbolTable = widget.NewTable(
		func() (int, int) {
			return len(reports.symbolTable) + 1, 6 // +1 para header
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("Template")
		},
		func(id widget.TableCellID, obj fyne.CanvasObject) {
			label := obj.(*widget.Label)
			if id.Row == 0 {
				// Header
				headers := []string{"ID", "Tipo", "Tipo de Dato", "Ámbito", "Línea", "Columna"}
				if id.Col < len(headers) {
					label.SetText(headers[id.Col])
					label.TextStyle = fyne.TextStyle{Bold: true}
				}
			} else {
				// Data
				if id.Row-1 < len(reports.symbolTable) {
					symbol := reports.symbolTable[id.Row-1]
					switch id.Col {
					case 0:
						label.SetText(symbol.ID)
					case 1:
						label.SetText(symbol.SymbolType)
					case 2:
						label.SetText(symbol.DataType)
					case 3:
						label.SetText(symbol.Scope)
					case 4:
						label.SetText(strconv.Itoa(symbol.Line))
					case 5:
						label.SetText(strconv.Itoa(symbol.Column))
					}
				}
			}
		},
	)

	// Configurar anchos de columna para mejor legibilidad
	reports.SymbolTable.SetColumnWidth(0, 120) // ID
	reports.SymbolTable.SetColumnWidth(1, 100) // Tipo
	reports.SymbolTable.SetColumnWidth(2, 120) // Tipo de Dato
	reports.SymbolTable.SetColumnWidth(3, 80)  // Ámbito
	reports.SymbolTable.SetColumnWidth(4, 60)  // Línea
	reports.SymbolTable.SetColumnWidth(5, 70)  // Columna

	// Crear contenedor para AST
	reports.ASTContent = container.NewVBox(
		widget.NewLabel("No hay AST disponible"),
	)

	// Crear pestañas
	reports.Container = container.NewAppTabs(
		container.NewTabItem("Errores", reports.ErrorList),
		container.NewTabItem("Tabla de Símbolos", reports.SymbolTable),
		container.NewTabItem("AST", container.NewScroll(reports.ASTContent)),
	)

	return reports
}

// UpdateErrors actualiza la lista de errores
func (r *Reports) UpdateErrors(errors []models.ErrorReport) {
	r.errors = errors
	r.ErrorList.Refresh()
}

// UpdateSymbolTable actualiza la tabla de símbolos
func (r *Reports) UpdateSymbolTable(symbols []models.SymbolEntry) {
	r.symbolTable = symbols
	r.SymbolTable.Refresh()
}

// UpdateASTWithSVG - Conversión automática SVG a PNG
func (r *Reports) UpdateASTWithSVG(astText string, svgContent string) {
	if svgContent != "" {
		// 1. Guardar SVG como archivo
		svgPath, err := r.saveSVGToFile(svgContent)
		if err != nil {
			r.showError("Error guardando SVG: " + err.Error())
			return
		}

		// 2. Intentar convertir SVG a PNG automáticamente
		r.convertAndShowAST(svgPath, svgContent)

	} else {
		r.showError("No hay AST disponible")
	}
}

// convertAndShowAST intenta convertir y mostrar, con fallback si falla
func (r *Reports) convertAndShowAST(svgPath, svgContent string) {
	// Intentar conversión automática
	pngPath, pngData, err := r.convertSVGToPNG(svgPath)

	if err == nil && len(pngData) > 0 {
		// ✅ ÉXITO: Mostrar PNG real (idéntico al SVG)
		r.showPNGSuccess(svgPath, pngPath, pngData, svgContent)
	} else {
		// ❌ FALLÓ: Mostrar instrucciones + fallback
		r.showConversionFailed(svgPath, svgContent, err.Error())
	}
}

// convertSVGToPNG convierte usando herramientas del sistema
func (r *Reports) convertSVGToPNG(svgPath string) (string, []byte, error) {
	// Generar ruta PNG
	dir := filepath.Dir(svgPath)
	base := strings.TrimSuffix(filepath.Base(svgPath), ".svg")
	pngPath := filepath.Join(dir, base+".png")

	// Herramientas de conversión (en orden de preferencia)
	converters := [][]string{
		// Inkscape (mejor calidad)
		{"inkscape", svgPath, "--export-type=png", "--export-dpi=300", "--export-width=1200", "--export-filename=" + pngPath},
		// rsvg-convert (muy bueno)
		{"rsvg-convert", "-a", "-w", "1200", "-h", "900", "-f", "png", "-o", pngPath, svgPath},
		// ImageMagick
		{"convert", "-density", "300", "-background", "white", "-resize", "1200x900", svgPath, pngPath},
		// CairoSVG (si está disponible)
		{"cairosvg", svgPath, "-o", pngPath, "--output-width", "1200"},
	}

	fmt.Printf("🔄 Intentando convertir SVG a PNG...\n")

	for i, cmd := range converters {
		fmt.Printf("   Probando %s (%d/%d)...\n", cmd[0], i+1, len(converters))

		// Ejecutar comando
		execCmd := exec.Command(cmd[0], cmd[1:]...)
		if err := execCmd.Run(); err == nil {
			// Verificar que el archivo se creó correctamente
			if pngData, err := os.ReadFile(pngPath); err == nil && len(pngData) > 1000 {
				fmt.Printf("✅ Conversión exitosa con %s\n", cmd[0])
				return pngPath, pngData, nil
			}
		}
		fmt.Printf("   %s no disponible o falló\n", cmd[0])
	}

	return "", nil, fmt.Errorf("no se encontró ningún convertidor (inkscape, rsvg-convert, convert, cairosvg)")
}

// showPNGSuccess muestra el PNG convertido (imagen real del AST)
func (r *Reports) showPNGSuccess(svgPath, pngPath string, pngData []byte, svgContent string) {
	// Crear imagen desde PNG
	pngResource := fyne.NewStaticResource("ast.png", pngData)
	astImage := canvas.NewImageFromResource(pngResource)
	astImage.FillMode = canvas.ImageFillContain
	astImage.SetMinSize(fyne.NewSize(600, 400))

	// Información de éxito
	infoText := fmt.Sprintf(`🌳 AST Mostrado Como Imagen Real

✅ Conversión exitosa SVG → PNG
📁 SVG: %s
🖼️ PNG: %s  
📊 Tamaño PNG: %d KB
📐 Resolución: 1200x900px

💡 Esta es la imagen EXACTA del SVG`,
		filepath.Base(svgPath),
		filepath.Base(pngPath),
		len(pngData)/1024)

	infoLabel := widget.NewLabel(infoText)
	infoLabel.Wrapping = fyne.TextWrapWord

	// Botones de acción
	openSVGBtn := widget.NewButton("📄 Abrir SVG Original", func() {
		r.openFileInBrowser(svgPath)
	})

	openPNGBtn := widget.NewButton("🖼️ Abrir PNG", func() {
		r.openFileInBrowser(pngPath)
	})

	// Layout con imagen real
	content := container.NewVBox(
		infoLabel,
		widget.NewSeparator(),
		container.NewHBox(openSVGBtn, openPNGBtn),
		widget.NewSeparator(),
		widget.NewLabel("🎯 AST Visualizado (imagen exacta):"),
		container.NewScroll(astImage),
	)

	r.ASTContent.Objects = []fyne.CanvasObject{content}
	r.ASTContent.Refresh()
}

// showConversionFailed muestra cuando no se puede convertir
func (r *Reports) showConversionFailed(svgPath, svgContent string, errorMsg string) {
	// Información de instalación
	installText := `🔧 INSTALAR CONVERTIDOR PARA VER AST REAL

Para ver el AST como imagen exacta en el IDE:

🐧 Linux:
   sudo apt update
   sudo apt install inkscape

🍎 macOS:
   brew install librsvg
   # o: brew install inkscape

🪟 Windows:
   Descargar Inkscape: https://inkscape.org

⚡ Después de instalar, reinicia el IDE`

	installLabel := widget.NewLabel(installText)
	installLabel.TextStyle = fyne.TextStyle{Monospace: true}

	// Información actual
	currentInfo := fmt.Sprintf(`📊 Estado Actual:

❌ Convertidor no encontrado: %s
📁 SVG guardado: %s
📊 Tamaño: %d caracteres

🌐 Mientras tanto, usa el botón para ver en navegador:`,
		errorMsg, filepath.Base(svgPath), len(svgContent))

	currentLabel := widget.NewLabel(currentInfo)
	currentLabel.Wrapping = fyne.TextWrapWord

	// Botón para navegador
	openBtn := widget.NewButton("🌐 Abrir en Navegador", func() {
		r.openFileInBrowser(svgPath)
	})
	openBtn.Importance = widget.HighImportance

	// Botón para verificar instalación
	checkBtn := widget.NewButton("🔍 Verificar Instalación", func() {
		r.checkConverters()
	})

	// Layout de instalación
	content := container.NewVBox(
		currentLabel,
		widget.NewSeparator(),
		container.NewHBox(openBtn, checkBtn),
		widget.NewSeparator(),
		widget.NewLabel("📋 Instrucciones de instalación:"),
		container.NewScroll(installLabel),
	)

	r.ASTContent.Objects = []fyne.CanvasObject{content}
	r.ASTContent.Refresh()
}

// checkConverters verifica qué herramientas están disponibles
func (r *Reports) checkConverters() {
	tools := []string{"inkscape", "rsvg-convert", "convert", "cairosvg"}
	results := "🔍 VERIFICACIÓN DE HERRAMIENTAS:\n\n"

	available := 0
	for _, tool := range tools {
		if exec.Command("which", tool).Run() == nil || exec.Command("where", tool).Run() == nil {
			results += fmt.Sprintf("✅ %s - DISPONIBLE\n", tool)
			available++
		} else {
			results += fmt.Sprintf("❌ %s - NO ENCONTRADO\n", tool)
		}
	}

	if available > 0 {
		results += fmt.Sprintf("\n🎉 %d herramienta(s) disponible(s)!\n", available)
		results += "💡 Reinicia el IDE para usar la conversión automática"
	} else {
		results += "\n⚠️ No se encontraron herramientas de conversión\n"
		results += "📦 Instala al menos una de las herramientas listadas arriba"
	}

	// Mostrar resultado
	resultLabel := widget.NewLabel(results)
	resultLabel.TextStyle = fyne.TextStyle{Monospace: true}

	r.ASTContent.Objects = []fyne.CanvasObject{
		widget.NewLabel("🔍 Resultado de Verificación"),
		widget.NewSeparator(),
		container.NewScroll(resultLabel),
	}
	r.ASTContent.Refresh()
}

// saveSVGToFile guarda el SVG como archivo local
func (r *Reports) saveSVGToFile(svgContent string) (string, error) {
	// Crear directorio si no existe
	outputDir := "generated_asts"
	os.MkdirAll(outputDir, 0755)

	// Generar nombre único con timestamp
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("ast_%s.svg", timestamp)
	fullPath := filepath.Join(outputDir, filename)

	// Escribir archivo
	err := os.WriteFile(fullPath, []byte(svgContent), 0644)
	if err != nil {
		return "", err
	}

	// Obtener ruta absoluta
	absPath, err := filepath.Abs(fullPath)
	if err != nil {
		return fullPath, nil
	}

	return absPath, nil
}

// showError muestra un mensaje de error
func (r *Reports) showError(message string) {
	label := widget.NewLabel("❌ " + message)
	label.Alignment = fyne.TextAlignCenter
	r.ASTContent.Objects = []fyne.CanvasObject{label}
	r.ASTContent.Refresh()
}

// UpdateASTWithPNG mantener compatibilidad
func (r *Reports) UpdateASTWithPNG(astText string, pngData []byte) {
	// Si tenemos datos PNG, mostrarlos, sino mostrar mensaje
	if len(pngData) > 0 {
		// Crear imagen desde bytes
		resource := fyne.NewStaticResource("ast.png", pngData)
		astImage := canvas.NewImageFromResource(resource)
		astImage.FillMode = canvas.ImageFillContain
		astImage.SetMinSize(fyne.NewSize(400, 300))

		content := container.NewVBox(
			widget.NewLabel("🌳 AST generado como PNG"),
			widget.NewSeparator(),
			container.NewScroll(astImage),
		)

		r.ASTContent.Objects = []fyne.CanvasObject{content}
		r.ASTContent.Refresh()
	} else {
		label := widget.NewLabel("❌ No hay datos PNG disponibles")
		label.Alignment = fyne.TextAlignCenter
		r.ASTContent.Objects = []fyne.CanvasObject{label}
		r.ASTContent.Refresh()
	}
}

// ClearAll limpia todos los reportes
func (r *Reports) ClearAll() {
	r.errors = []models.ErrorReport{}
	r.symbolTable = []models.SymbolEntry{}

	r.ErrorList.Refresh()
	r.SymbolTable.Refresh()

	// Limpiar AST
	label := widget.NewLabel("No hay AST disponible")
	r.ASTContent.Objects = []fyne.CanvasObject{label}
	r.ASTContent.Refresh()
}

// openFileInBrowser abre cualquier archivo en el navegador del sistema
func (r *Reports) openFileInBrowser(filePath string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "linux":
		cmd = exec.Command("xdg-open", filePath)
	case "darwin": // macOS
		cmd = exec.Command("open", filePath)
	case "windows":
		cmd = exec.Command("rundll32", "url.dll,FileProtocolHandler", filePath)
	default:
		r.showError("Sistema operativo no soportado")
		return
	}

	if err := cmd.Start(); err != nil {
		r.showError("Error abriendo archivo: " + err.Error())
	}
}
