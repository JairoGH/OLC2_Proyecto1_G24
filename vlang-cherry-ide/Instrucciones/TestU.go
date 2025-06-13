package instrucciones

import "testing"

func ProbarTipoVector(t *testing.T) {

	matrix := map[string]bool{
		"[int]":   true,
		"int":     false,
		"[[int]]": false,
		"[]":      true,
	}

	for k, v := range matrix {
		if EsTipoVector(k) != v {
			t.Errorf("isVector(%s) != %t", k, v)
		}
	}

}

func ProbarTipoMatriz(t *testing.T) {
	matrix := map[string]bool{
		"[int]":     false,
		"int":       false,
		"[[int]]":   true,
		"[]":        false,
		"[[[int]]]": true,
		"[":         false,
		"[[]":       false,
	}

	for k, v := range matrix {
		if EsTipoMatriz(k) != v {
			t.Errorf("isMatrix(%s) != %t", k, v)
		}
	}
}
