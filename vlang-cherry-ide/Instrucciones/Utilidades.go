package instrucciones

import (
	"main/tiposDeDato"
	"regexp"
	"strings"
)

func CadenaAVector(s *tiposDeDato.ValorCadena) *TipoVector {
	items := make([]tiposDeDato.ValorInterno, 0, len(s.InternalValor))
	for _, rune := range s.InternalValor {
		// cada rune lo convertimos en string de 1 carácter
		items = append(items, &tiposDeDato.ValorCadena{
			InternalValor: string(rune),
		})
	}
	// Aquí indicamos que es un vector de cadenas
	tag := "[" + tiposDeDato.TIPO_CADENA + "]"
	return NewTipoVector(items, tag, tiposDeDato.TIPO_CADENA)
}

func EsTipoVector(_type string) bool {

	// Vector starts with only one [ and ends with only one ]
	formatoVector := "^\\[.*\\]$"

	// Matrix starts with AT LEAST two [[ and ends with at least two ]]
	formatoMatriz := "^\\[\\[.*\\]\\](\\[.*\\]\\])*$"

	// match vector pattern but not matrix pattern

	match, _ := regexp.MatchString(formatoVector, _type)
	match2, _ := regexp.MatchString(formatoMatriz, _type)

	return match && !match2
}

func EliminarCorchetes(s string) string {
	return strings.Trim(s, "[]")
}

func EsTipoMatriz(_type string) bool {

	// Matrix starts with AT LEAST two [[ and ends with at least two ]]
	formatoMatriz := "^\\[\\[.*\\]\\](\\[.*\\]\\])*$"

	match, _ := regexp.MatchString(formatoMatriz, _type)

	return match
}
