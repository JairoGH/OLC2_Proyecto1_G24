package instrucciones

import (
    "main/parser"
    "main/tiposDeDato"

    "github.com/antlr4-go/antlr/v4"
)

type Struct struct {
    Name   string
    Fields []parser.IStruct_propContext
    Token  antlr.Token
}

type StructField struct {
    Name  string
    Type  string
    Token antlr.Token
}

type StructInitValue struct {
    Name  string
    Value tiposDeDato.ValorInterno
    Token antlr.Token
}
