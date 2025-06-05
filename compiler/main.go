package main

import (
	"compiler/Errores"
	compiler "compiler/parser"
	"fmt"
	"os"

	"github.com/antlr4-go/antlr/v4"
)

type EvalVisitor struct {
	*compiler.BaseVGrammarVisitor
}

func (v *EvalVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

// Visit the root program node
func (v *EvalVisitor) VisitPrograma(ctx *compiler.ProgramaContext) interface{} {
	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}
	return nil
}

// Implementaciones mínimas para evitar nil pointer

func (v *EvalVisitor) VisitStmtDecl(ctx *compiler.StmtDeclContext) interface{} {
	fmt.Println("→ Declaración detectada.")
	return nil
}

func (v *EvalVisitor) VisitStmtAssign(ctx *compiler.StmtAssignContext) interface{} {
	fmt.Println("→ Asignación detectada.")
	return nil
}

func (v *EvalVisitor) VisitStmtFuncCall(ctx *compiler.StmtFuncCallContext) interface{} {
	fmt.Println("→ Llamada a función detectada.")
	return nil
}

func (v *EvalVisitor) VisitStmtIf(ctx *compiler.StmtIfContext) interface{} {
	fmt.Println("→ Sentencia if detectada.")
	return nil
}

func (v *EvalVisitor) VisitStmtFor(ctx *compiler.StmtForContext) interface{} {
	fmt.Println("→ Bucle for detectado.")
	return nil
}

func (v *EvalVisitor) VisitStmtFuncDecl(ctx *compiler.StmtFuncDeclContext) interface{} {
	fmt.Println("→ Función declarada.")
	return nil
}

func (v *EvalVisitor) VisitStmtStructDecl(ctx *compiler.StmtStructDeclContext) interface{} {
	fmt.Println("→ Struct declarado.")
	return nil
}

func (v *EvalVisitor) VisitStmtFlow(ctx *compiler.StmtFlowContext) interface{} {
	fmt.Println("→ Sentencia de flujo (return, break, continue) detectada.")
	return nil
}

func (v *EvalVisitor) VisitPrintCall(ctx *compiler.PrintCallContext) interface{} {
	fmt.Println("→ print() detectado.")
	return nil
}

func (v *EvalVisitor) VisitPrintLnCall(ctx *compiler.PrintLnCallContext) interface{} {
	fmt.Println("→ println() detectado.")
	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: go run main.go <archivo.vch>")
		return
	}
	archivo := os.Args[1]
	contenido, err := os.ReadFile(archivo)
	if err != nil {
		fmt.Println("Error al leer archivo:", err)
		return
	}

	input := antlr.NewInputStream(string(contenido))

	errorLexico := Errores.NewErrorLexico()
	lexer := compiler.NewVLexer(input)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(errorLexico)

	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	parser := compiler.NewVGrammar(tokens)
	syntaxError := Errores.NewSyntaxError(errorLexico.TablaErrores)
	parser.RemoveErrorListeners()
	parser.SetErrorHandler(Errores.NewErrorPersonalizado())
	parser.AddErrorListener(syntaxError)

	tree := parser.Programa()

	fmt.Println("\n--- Árbol Sintáctico ---")
	fmt.Println(tree.ToStringTree(nil, parser))

	visitor := &EvalVisitor{}
	visitor.Visit(tree)

	fmt.Println("\n--- Errores Léxicos Detectados ---")
	errores := errorLexico.TablaErrores.GetErrors()
	if len(errores) > 0 {
		for i, err := range errores {
			fmt.Printf("No. %d:\n  Línea: %d\n  Columna: %d\n  Mensaje: %s\n  Tipo: %s\n\n",
				i+1, err.Linea, err.Columna, err.Descripcion, err.Tipo)
		}
	} else {
		fmt.Println("No se detectaron errores léxicos")
	}
}
