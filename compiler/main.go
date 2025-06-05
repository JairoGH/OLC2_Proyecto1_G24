package main

import (
	// Ajusta si tu módulo se llama diferente
	"compiler/Errores"
	compiler "compiler/parser"
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type EvalVisitor struct {
	*compiler.BaseVGrammarVisitor
}

func (v *EvalVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(v)
}

func (v *EvalVisitor) VisitPrograma(ctx *compiler.ProgramaContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *EvalVisitor) VisitAddSub(ctx *compiler.AddSubContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(int)
	right := v.Visit(ctx.Expr(1)).(int)

	switch ctx.GetOp().GetTokenType() {
	case compiler.VLexerOP_SUMA:
		return left + right
	case compiler.VLexerOP_RESTA:
		return left - right
	}
	return 0
}

func (v *EvalVisitor) VisitMulDiv(ctx *compiler.MulDivContext) interface{} {
	left := v.Visit(ctx.Expr(0)).(int)
	right := v.Visit(ctx.Expr(1)).(int)

	switch ctx.GetOp().GetTokenType() {
	case compiler.VLexerOP_MULT:
		return left * right
	case compiler.VLexerOP_DIV:
		return left / right
	}
	return 0
}

func (v *EvalVisitor) VisitInt(ctx *compiler.IntContext) interface{} {
	var value int
	fmt.Sscanf(ctx.GetText(), "%d", &value)
	return value
}

func (v *EvalVisitor) VisitParens(ctx *compiler.ParensContext) interface{} {
	return v.Visit(ctx.Expr())
}

func main() {
	input := antlr.NewInputStream("100 + 4 * (2 - 1) @ ")

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
	parser.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	tree := parser.Programa()

	visitor := &EvalVisitor{}
	result := visitor.Visit(tree)

	fmt.Println("Resultado:", result)

	// Imprimir los errores léxicos detectados
	fmt.Println("\n--- Errores Léxicos Detectados ---")
	errores := errorLexico.TablaErrores.GetErrors()
	if len(errores) > 0 {
		for i, err := range errores {
			fmt.Printf("No. %d:\n", i+1)
			fmt.Printf("  Línea: %d\n", err.Linea)
			fmt.Printf("  Columna: %d\n", err.Columna)
			fmt.Printf("  Mensaje: %s\n", err.Descripcion)
			fmt.Printf("  Tipo: %s\n", err.Tipo)
			fmt.Println()
		}
	} else {
		fmt.Println("No se detectaron errores léxicos")
	}
}
