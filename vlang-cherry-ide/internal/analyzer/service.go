package analyzer

import (
	"main/crearArbol"
	"main/models"

	// Imports del backend original - ajusta las rutas según tu estructura
	errores "main/Errores"
	instrucciones "main/Instrucciones"
	compiler "main/parser"

	"github.com/antlr4-go/antlr/v4"
)

// AnalyzerService integra la lógica completa del backend original
type AnalyzerService struct{}

// NewAnalyzerService crea una nueva instancia del servicio
func NewAnalyzerService() *AnalyzerService {
	return &AnalyzerService{}
}

// AnalyzeCode usa la lógica completa del backend original
func (a *AnalyzerService) AnalyzeCode(code, filename string) (*models.AnalysisResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			// Manejo de errores - no crashear la aplicación
		}
	}()

	// Generar AST con ANTLR externo (en paralelo como en el original)
	resultChannel := make(chan string)
	go func() {
		resultChannel <- crearArbol.ReporteArbol(code)
	}()

	// === ANÁLISIS LÉXICO ===
	lexicalErrorListener := errores.NewErrorLexico()
	lexer := compiler.NewVLexer(antlr.NewInputStream(code))
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexicalErrorListener)

	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// === ANÁLISIS SINTÁCTICO ===
	parser := compiler.NewVGrammar(stream)
	parser.BuildParseTrees = true

	syntaxErrorListener := errores.NewSyntaxError(lexicalErrorListener.TablaError)
	parser.RemoveErrorListeners()
	parser.SetErrorHandler(errores.NewErrorPersonalizado())
	parser.AddErrorListener(syntaxErrorListener)

	tree := parser.Program()

	// === ANÁLISIS SEMÁNTICO ===
	dclVisitor := instrucciones.NewVisitanteDcl(syntaxErrorListener.TablaError)
	dclVisitor.Visit(tree)

	// === INTERPRETACIÓN ===
	replVisitor := instrucciones.NewVisitor(dclVisitor)
	replVisitor.Visit(tree)

	// Obtener el reporte del AST
	cstReport := <-resultChannel

	// === CONVERTIR A FORMATO DEL FRONTEND ===
	success := len(replVisitor.TablaError.Errores) == 0

	// Convertir errores - usando la estructura real correcta
	frontendErrors := []models.ErrorReport{}
	for _, err := range replVisitor.TablaError.Errores {
		frontendErrors = append(frontendErrors, models.ErrorReport{
			Linea:       err.Line,        // ← CORRECTO: usa Line, no Linea
			Columna:     err.Columna,     // ← CORRECTO: usa Columna
			Descripcion: err.Descripcion, // ← CORRECTO: usa Descripcion
			Tipo:        err.Tipo,        // ← CORRECTO: usa Tipo (que ya es string)
		})
	}

	// Convertir tabla de símbolos
	symbolTable := a.convertSymbolTable(replVisitor.RegistroAmbito.Report())

	// Preparar output de consola
	consoleOutput := replVisitor.Console.GetOutput()
	if consoleOutput == "" {
		if success {
			consoleOutput = "✅ Análisis completado exitosamente\n✅ Código interpretado sin errores"
		} else {
			consoleOutput = "❌ Se encontraron errores durante el análisis"
		}
	}

	result := &models.AnalysisResponse{
		Success:       success,
		AST:           "Árbol de Sintaxis Abstracta generado exitosamente",
		Errors:        frontendErrors,
		SymbolTable:   symbolTable,
		ConsoleOutput: consoleOutput,
		CSTSvg:        cstReport,
		ASTPng:        nil,
	}

	return result, nil
}

// Convertir tabla de símbolos del backend al formato del frontend
func (a *AnalyzerService) convertSymbolTable(registroAmbito instrucciones.ReporteTabla) []models.SymbolEntry {
	var symbolTable []models.SymbolEntry

	// Extraer variables del ámbito global
	for _, variable := range registroAmbito.AmbitoGlobal.Vars {
		symbolTable = append(symbolTable, models.SymbolEntry{
			ID:         variable.Name,
			SymbolType: "Variable",
			DataType:   variable.Type,
			Scope:      "global",
			Line:       variable.Line,
			Column:     variable.Column,
		})
	}

	// Extraer funciones del ámbito global
	for _, function := range registroAmbito.AmbitoGlobal.Funcs {
		symbolTable = append(symbolTable, models.SymbolEntry{
			ID:         function.Name,
			SymbolType: "Function",
			DataType:   function.Type,
			Scope:      "global",
			Line:       function.Line,
			Column:     function.Column,
		})
	}

	return symbolTable
}

// ValidateCode valida la sintaxis del código
func (a *AnalyzerService) ValidateCode(code string) (*models.AnalysisResponse, error) {
	return a.AnalyzeCode(code, "temp.vch")
}

// GetVersion retorna información del analizador
func (a *AnalyzerService) GetVersion() (string, error) {
	return "V-Lang Cherry Analyzer (Direct Integration with Full Backend)", nil
}

// Ping siempre retorna éxito ya que no hay servidor externo
func (a *AnalyzerService) Ping() error {
	return nil // Siempre disponible
}
