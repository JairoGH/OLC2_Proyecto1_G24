package instrucciones

import (
	"main/tiposDeDato"
	"regexp"
	"strings"
)

func CadenaAVector(s *tiposDeDato.ValorCadena) *TipoVector {

	items := make([]tiposDeDato.ValorInterno, 0)

	for _, c := range s.InternalValor {
		items = append(items, &tiposDeDato.ValorCaracter{InternalValor: string(c)})
	}

	return NewTipoVector(items, "["+tiposDeDato.TIPO_CARACTER+"]", tiposDeDato.TIPO_CARACTER)

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
