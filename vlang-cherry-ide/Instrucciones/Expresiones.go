package instrucciones

import (
	"main/tiposDeDato"
)

type funcionEvaluar func(tiposDeDato.ValorInterno, tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) // takes 2 values and returns a value
type funcionConv func(tiposDeDato.ValorInterno) tiposDeDato.ValorInterno                                              // takes a value and returns a value (different type)

type VerificacionBinaria struct {
	TipoIzq  string // allowed left type
	TipoDcha string // allowed right type
	ConvIzq  funcionConv
	ConvDcha funcionConv
	Evaluar  funcionEvaluar
}

type MetodoBinario struct {
	Nombre               string
	Validaciones         []VerificacionBinaria
	Viceversa            bool // if true, the validation is also performed in the opposite order
	EvaluacionPorDefecto funcionEvaluar
}

func (s *MetodoBinario) ValidarExp(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {

	// nil in any side is, by default return nil

	if left.Type() == tiposDeDato.TIPO_NULO || right.Type() == tiposDeDato.TIPO_NULO {
		return false, "No es posible realizar operaciones con valores nulos", tiposDeDato.NuloPorDefecto
	}

	for _, valid := range s.Validaciones {

		if valid.TipoIzq == left.Type() && valid.TipoDcha == right.Type() {

			if valid.ConvIzq != nil {
				left = valid.ConvIzq(left)
			}

			if valid.ConvDcha != nil {
				right = valid.ConvDcha(right)
			}

			if valid.Evaluar != nil {
				return valid.Evaluar(left, right)
			}

			return s.EvaluacionPorDefecto(left, right)
		}

		if s.Viceversa && valid.TipoIzq == right.Type() && valid.TipoDcha == left.Type() {

			if valid.ConvIzq != nil {
				right = valid.ConvIzq(right)
			}

			if valid.ConvDcha != nil {
				left = valid.ConvDcha(left)
			}

			if valid.Evaluar != nil {
				return valid.Evaluar(left, right)
			}

			return s.EvaluacionPorDefecto(left, right)
		}

	}

	msg := "No es posible realizar la operación '" + s.Nombre + "' con los tipos '" + left.Type() + "' y '" + right.Type() + "'"

	return false, msg, tiposDeDato.NuloPorDefecto
}

// * arithmetic operators

// int + int; float + float; float + int (viceversa); string + string
var SumaExpresion = MetodoBinario{
	Nombre:               "+",
	Viceversa:            true,
	EvaluacionPorDefecto: nil,
	Validaciones: []VerificacionBinaria{
		{
			TipoIzq:  tiposDeDato.TIPO_ENTERO,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorEntero{
					InternalValor: left.(*tiposDeDato.ValorEntero).InternalValor + right.(*tiposDeDato.ValorEntero).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_DECIMAL,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorDecimal{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor + right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: func(v tiposDeDato.ValorInterno) tiposDeDato.ValorInterno {
				return &tiposDeDato.ValorDecimal{
					InternalValor: float64(v.(*tiposDeDato.ValorEntero).InternalValor),
				}
			},
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorDecimal{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor + right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_CADENA,
			TipoDcha: tiposDeDato.TIPO_CADENA,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorCadena{
					InternalValor: left.(*tiposDeDato.ValorCadena).InternalValor + right.(*tiposDeDato.ValorCadena).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_CARACTER,
			TipoDcha: tiposDeDato.TIPO_CADENA,
			ConvIzq: func(v tiposDeDato.ValorInterno) tiposDeDato.ValorInterno {
				return &tiposDeDato.ValorCadena{
					InternalValor: string(v.(*tiposDeDato.ValorCaracter).InternalValor),
				}
			},
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorCadena{
					InternalValor: left.(*tiposDeDato.ValorCadena).InternalValor + right.(*tiposDeDato.ValorCadena).InternalValor,
				}
			},
		},
	},
}

// int - int; float - float; float - int (viceversa)
var RestaExpresion = MetodoBinario{
	Nombre:               "-",
	Viceversa:            true,
	EvaluacionPorDefecto: nil,
	Validaciones: []VerificacionBinaria{
		{
			TipoIzq:  tiposDeDato.TIPO_ENTERO,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorEntero{
					InternalValor: left.(*tiposDeDato.ValorEntero).InternalValor - right.(*tiposDeDato.ValorEntero).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_DECIMAL,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorDecimal{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor - right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: func(v tiposDeDato.ValorInterno) tiposDeDato.ValorInterno {
				return &tiposDeDato.ValorDecimal{
					InternalValor: float64(v.(*tiposDeDato.ValorEntero).InternalValor),
				}
			},
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorDecimal{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor - right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
	},
}

// int * int; float * float; float * int (viceversa)
var MultipliacionExpresion = MetodoBinario{
	Nombre:               "*",
	Viceversa:            true,
	EvaluacionPorDefecto: nil,
	Validaciones: []VerificacionBinaria{
		{
			TipoIzq:  tiposDeDato.TIPO_ENTERO,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorEntero{
					InternalValor: left.(*tiposDeDato.ValorEntero).InternalValor * right.(*tiposDeDato.ValorEntero).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_DECIMAL,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorDecimal{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor * right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: func(v tiposDeDato.ValorInterno) tiposDeDato.ValorInterno {
				return &tiposDeDato.ValorDecimal{
					InternalValor: float64(v.(*tiposDeDato.ValorEntero).InternalValor),
				}
			},
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorDecimal{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor * right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
	},
}

// int / int; float / float; float / int (viceversa) !division by zero
var DivisionExpresion = MetodoBinario{
	Nombre:               "/",
	Viceversa:            true,
	EvaluacionPorDefecto: nil,
	Validaciones: []VerificacionBinaria{
		{
			TipoIzq:  tiposDeDato.TIPO_ENTERO,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {

				if right.(*tiposDeDato.ValorEntero).InternalValor == 0 {
					return false, "No se puede dividir entre cero", tiposDeDato.NuloPorDefecto
				}

				return true, "", &tiposDeDato.ValorEntero{
					InternalValor: left.(*tiposDeDato.ValorEntero).InternalValor / right.(*tiposDeDato.ValorEntero).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_DECIMAL,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {

				if right.(*tiposDeDato.ValorDecimal).InternalValor == 0 {
					return false, "No se puede dividir entre cero", tiposDeDato.NuloPorDefecto
				}

				return true, "", &tiposDeDato.ValorDecimal{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor / right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: func(v tiposDeDato.ValorInterno) tiposDeDato.ValorInterno {
				return &tiposDeDato.ValorDecimal{
					InternalValor: float64(v.(*tiposDeDato.ValorEntero).InternalValor),
				}
			},
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {

				if right.(*tiposDeDato.ValorDecimal).InternalValor == 0 {
					return false, "No se puede dividir entre cero", tiposDeDato.NuloPorDefecto
				}

				return true, "", &tiposDeDato.ValorDecimal{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor / right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
	},
}

// int % int; !division by zero
var ModuloExpresion = MetodoBinario{
	Nombre:               "%",
	Viceversa:            true,
	EvaluacionPorDefecto: nil,
	Validaciones: []VerificacionBinaria{
		{
			TipoIzq:  tiposDeDato.TIPO_ENTERO,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {

				if right.(*tiposDeDato.ValorEntero).InternalValor == 0 {
					return false, "No se puede dividir entre cero", tiposDeDato.NuloPorDefecto
				}

				return true, "", &tiposDeDato.ValorEntero{
					InternalValor: left.(*tiposDeDato.ValorEntero).InternalValor % right.(*tiposDeDato.ValorEntero).InternalValor,
				}
			},
		},
	},
}

// * comparison operators

// int == int; float == float; bool == bool; string == string; char == char
func ExpresionTipoIgual(Nombre string, eval funcionEvaluar) MetodoBinario {
	return MetodoBinario{
		Nombre:               Nombre,
		Viceversa:            true, // Habilita la validación inversa: a == b y b == a son tratadas igual
		EvaluacionPorDefecto: eval, // Función por defecto para evaluar si no se especifica una Evaluar explícita
		Validaciones: []VerificacionBinaria{
			// Comparaciones válidas entre tipos idénticos
			{
				TipoIzq:  tiposDeDato.TIPO_ENTERO,
				TipoDcha: tiposDeDato.TIPO_ENTERO,
				ConvIzq:  nil,
				ConvDcha: nil,
				Evaluar:  nil, // Se usará la EvaluacionPorDefecto
			},
			{
				TipoIzq:  tiposDeDato.TIPO_DECIMAL,
				TipoDcha: tiposDeDato.TIPO_DECIMAL,
				ConvIzq:  nil,
				ConvDcha: nil,
				Evaluar:  nil,
			},
			{
				TipoIzq:  tiposDeDato.TIPO_BOOLEAN,
				TipoDcha: tiposDeDato.TIPO_BOOLEAN,
				ConvIzq:  nil,
				ConvDcha: nil,
				Evaluar:  nil,
			},
			{
				TipoIzq:  tiposDeDato.TIPO_CADENA,
				TipoDcha: tiposDeDato.TIPO_CADENA,
				ConvIzq:  nil,
				ConvDcha: nil,
				Evaluar:  nil,
			},
			{
				TipoIzq:  tiposDeDato.TIPO_CARACTER,
				TipoDcha: tiposDeDato.TIPO_CARACTER,
				ConvIzq:  nil,
				ConvDcha: nil,
				Evaluar:  nil,
			},

			// Comparaciones cruzadas int == float64 (conversión implícita del entero a decimal)
			{
				TipoIzq:  tiposDeDato.TIPO_ENTERO,
				TipoDcha: tiposDeDato.TIPO_DECIMAL,
				ConvIzq: func(v tiposDeDato.ValorInterno) tiposDeDato.ValorInterno {
					return &tiposDeDato.ValorDecimal{
						InternalValor: float64(v.(*tiposDeDato.ValorEntero).InternalValor),
					}
				},
				ConvDcha: nil,
				Evaluar:  nil,
			},
			{
				TipoIzq:  tiposDeDato.TIPO_DECIMAL,
				TipoDcha: tiposDeDato.TIPO_ENTERO,
				ConvIzq:  nil,
				ConvDcha: func(v tiposDeDato.ValorInterno) tiposDeDato.ValorInterno {
					return &tiposDeDato.ValorDecimal{
						InternalValor: float64(v.(*tiposDeDato.ValorEntero).InternalValor),
					}
				},
				Evaluar: nil,
			},
		},
	}
}

var IgualExpresion = ExpresionTipoIgual("==", func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
	return true, "", &tiposDeDato.ValorBool{
		InternalValor: left.Value() == right.Value(),
	}
})

var NotExpresion = ExpresionTipoIgual("!=", func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
	return true, "", &tiposDeDato.ValorBool{
		InternalValor: left.Value() != right.Value(),
	}
})

var MenorExpresion = MetodoBinario{
	Nombre:               "<",
	Viceversa:            true,
	EvaluacionPorDefecto: nil,
	Validaciones: []VerificacionBinaria{
		{
			TipoIzq:  tiposDeDato.TIPO_ENTERO,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorEntero).InternalValor < right.(*tiposDeDato.ValorEntero).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_DECIMAL,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor < right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_CADENA,
			TipoDcha: tiposDeDato.TIPO_CADENA,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorCadena).InternalValor < right.(*tiposDeDato.ValorCadena).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_CARACTER,
			TipoDcha: tiposDeDato.TIPO_CARACTER,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorCaracter).InternalValor < right.(*tiposDeDato.ValorCaracter).InternalValor,
				}
			},
		},
	},
}

var MenorQUeExpresion = MetodoBinario{
	Nombre:               "<=",
	Viceversa:            true,
	EvaluacionPorDefecto: nil,
	Validaciones: []VerificacionBinaria{
		{
			TipoIzq:  tiposDeDato.TIPO_ENTERO,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorEntero).InternalValor <= right.(*tiposDeDato.ValorEntero).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_DECIMAL,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor <= right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_CADENA,
			TipoDcha: tiposDeDato.TIPO_CADENA,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorCadena).InternalValor <= right.(*tiposDeDato.ValorCadena).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_CARACTER,
			TipoDcha: tiposDeDato.TIPO_CARACTER,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorCaracter).InternalValor <= right.(*tiposDeDato.ValorCaracter).InternalValor,
				}
			},
		},
	},
}

var MayorExpresion = MetodoBinario{
	Nombre:               ">",
	Viceversa:            true,
	EvaluacionPorDefecto: nil,
	Validaciones: []VerificacionBinaria{
		{
			TipoIzq:  tiposDeDato.TIPO_ENTERO,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorEntero).InternalValor > right.(*tiposDeDato.ValorEntero).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_DECIMAL,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor > right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_CADENA,
			TipoDcha: tiposDeDato.TIPO_CADENA,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorCadena).InternalValor > right.(*tiposDeDato.ValorCadena).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_CARACTER,
			TipoDcha: tiposDeDato.TIPO_CARACTER,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorCaracter).InternalValor > right.(*tiposDeDato.ValorCaracter).InternalValor,
				}
			},
		},
	},
}

var MayorQueExpresion = MetodoBinario{
	Nombre:               ">=",
	Viceversa:            true,
	EvaluacionPorDefecto: nil,
	Validaciones: []VerificacionBinaria{
		{
			TipoIzq:  tiposDeDato.TIPO_ENTERO,
			TipoDcha: tiposDeDato.TIPO_ENTERO,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorEntero).InternalValor >= right.(*tiposDeDato.ValorEntero).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_DECIMAL,
			TipoDcha: tiposDeDato.TIPO_DECIMAL,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorDecimal).InternalValor >= right.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_CADENA,
			TipoDcha: tiposDeDato.TIPO_CADENA,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorCadena).InternalValor >= right.(*tiposDeDato.ValorCadena).InternalValor,
				}
			},
		},
		{
			TipoIzq:  tiposDeDato.TIPO_CARACTER,
			TipoDcha: tiposDeDato.TIPO_CARACTER,
			ConvIzq:  nil,
			ConvDcha: nil,
			Evaluar: func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: left.(*tiposDeDato.ValorCaracter).InternalValor >= right.(*tiposDeDato.ValorCaracter).InternalValor,
				}
			},
		},
	},
}

// * logical operators

func LogicaBinariaGenerica(Nombre string, eval funcionEvaluar) MetodoBinario {

	return MetodoBinario{
		Nombre:               Nombre,
		Viceversa:            true,
		EvaluacionPorDefecto: eval,
		Validaciones: []VerificacionBinaria{
			{
				TipoIzq:  tiposDeDato.TIPO_BOOLEAN,
				TipoDcha: tiposDeDato.TIPO_BOOLEAN,
				ConvIzq:  nil,
				ConvDcha: nil,
				Evaluar:  nil,
			},
		},
	}
}

var AndExpresion = LogicaBinariaGenerica("&&", func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
	return true, "", &tiposDeDato.ValorBool{
		InternalValor: left.(*tiposDeDato.ValorBool).InternalValor && right.(*tiposDeDato.ValorBool).InternalValor,
	}
})

var OrExpresion = LogicaBinariaGenerica("||", func(left, right tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
	return true, "", &tiposDeDato.ValorBool{
		InternalValor: left.(*tiposDeDato.ValorBool).InternalValor || right.(*tiposDeDato.ValorBool).InternalValor,
	}
})

var ExpresionBinaria = map[string]MetodoBinario{
	"+":  SumaExpresion,
	"-":  RestaExpresion,
	"*":  MultipliacionExpresion,
	"/":  DivisionExpresion,
	"%":  ModuloExpresion,
	"==": IgualExpresion,
	"!=": NotExpresion,
	"<":  MenorExpresion,
	"<=": MenorQUeExpresion,
	">":  MayorExpresion,
	">=": MayorQueExpresion,
	"&&": AndExpresion,
	"||": OrExpresion,
}

// UnaryStrats

type ValidarUnaria struct {
	Type       string // allowed type
	Conversion funcionConv
	Eval       funcionEvaluar
}

type ExpresionUnaria struct {
	Nombre               string
	Validaciones         []ValidarUnaria
	EvaluacionPorDefecto funcionEvaluar
}

func (s *ExpresionUnaria) ValidarExp(val tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {

	if val.Type() == tiposDeDato.TIPO_NULO {
		return false, "No es posible realizar operaciones con valores nulos", tiposDeDato.NuloPorDefecto
	}

	for _, valid := range s.Validaciones {

		if valid.Type == val.Type() {

			if valid.Conversion != nil {
				val = valid.Conversion(val)
			}

			if valid.Eval != nil {
				return valid.Eval(val, nil)
			}

			return s.EvaluacionPorDefecto(val, nil)
		}

	}

	msg := "No es posible realizar la operación '" + s.Nombre + "' con el tipo '" + val.Type() + "'"

	return false, msg, tiposDeDato.NuloPorDefecto
}

// * Not

var ExpresionNot = ExpresionUnaria{
	Nombre:               "!",
	EvaluacionPorDefecto: nil,
	Validaciones: []ValidarUnaria{
		{
			Type:       tiposDeDato.TIPO_BOOLEAN,
			Conversion: nil,
			Eval: func(i1, i2 tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorBool{
					InternalValor: !i1.(*tiposDeDato.ValorBool).InternalValor,
				}
			},
		},
	},
}

// * Minus

var ExpresionMenos = ExpresionUnaria{
	Nombre:               "-",
	EvaluacionPorDefecto: nil,
	Validaciones: []ValidarUnaria{
		{
			Type:       tiposDeDato.TIPO_ENTERO,
			Conversion: nil,
			Eval: func(i1, i2 tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorEntero{
					InternalValor: -i1.(*tiposDeDato.ValorEntero).InternalValor,
				}
			},
		},
		{
			Type:       tiposDeDato.TIPO_DECIMAL,
			Conversion: nil,
			Eval: func(i1, i2 tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {
				return true, "", &tiposDeDato.ValorDecimal{
					InternalValor: -i1.(*tiposDeDato.ValorDecimal).InternalValor,
				}
			},
		},
	},
}

var ExpresionesUnarias = map[string]ExpresionUnaria{
	"!": ExpresionNot,
	"-": ExpresionMenos,
}

// Early return strats

// * And

var RetornoAnd = ExpresionUnaria{
	Nombre: "&&",
	Validaciones: []ValidarUnaria{
		{
			Type:       tiposDeDato.TIPO_BOOLEAN,
			Conversion: nil,
			Eval: func(i1, i2 tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {

				if !i1.(*tiposDeDato.ValorBool).InternalValor {
					return true, "", &tiposDeDato.ValorBool{
						InternalValor: false,
					}
				}

				return false, "", nil
			},
		},
	},
}

// * Or

var RetornoOr = ExpresionUnaria{
	Nombre: "||",
	Validaciones: []ValidarUnaria{
		{
			Type:       tiposDeDato.TIPO_BOOLEAN,
			Conversion: nil,
			Eval: func(i1, i2 tiposDeDato.ValorInterno) (bool, string, tiposDeDato.ValorInterno) {

				if i1.(*tiposDeDato.ValorBool).InternalValor {
					return true, "", &tiposDeDato.ValorBool{
						InternalValor: true,
					}
				}

				return false, "", nil
			},
		},
	},
}

var RetornoAnticipado = map[string]ExpresionUnaria{
	"&&": RetornoAnd,
	"||": RetornoOr,
}
