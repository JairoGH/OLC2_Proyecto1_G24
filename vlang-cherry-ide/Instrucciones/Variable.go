package instrucciones

import (
	"main/tiposDeDato"

	"github.com/antlr4-go/antlr/v4"
)

type Variable struct {
	Name     string
	Value    tiposDeDato.ValorInterno
	Type     string
	IsConst  bool
	AllowNil bool
	Token    antlr.Token
	isProp   bool
}

func (v *Variable) TipoValidacion() (bool, string) {

    if v.Value == tiposDeDato.ValorNoIniPorDefecto {
        return true, ""
    }

    if v.Value == tiposDeDato.NuloPorDefecto && v.AllowNil {
        return true, ""
    }

    if v.Type != v.Value.Type() {

        // vector type validation
        if EsTipoVector(v.Type) && EsTipoVector(v.Value.Type()) {

            // CORRECCIÓN: Manejar tanto vectores con tipo "[]" como vectores vacíos
            vectorValue := v.Value.(*TipoVector)
            
            // Vector vacío genérico - asignar el tipo correcto
            if v.Value.Type() == "[]" || len(vectorValue.InternalValor) == 0 {
                // Modificar el tipo del vector para que coincida con el esperado
                vectorValue.ItemType = EliminarCorchetes(v.Type)
                vectorValue.FullType = v.Type
                return true, ""
            }

            // implicit vector conversion para vectores no vacíos
            targetType := EliminarCorchetes(v.Type) // inner type
            newConvertedItems := make([]tiposDeDato.ValorInterno, 0)

            for _, item := range vectorValue.InternalValor {
                convertedValue, ok := tiposDeDato.Casteo(targetType, item)

                if !ok {
                    break
                }
                newConvertedItems = append(newConvertedItems, convertedValue)
            }

            if len(newConvertedItems) == len(vectorValue.InternalValor) {
                // Actualizar el vector con los elementos convertidos
                vectorValue.InternalValor = newConvertedItems
                vectorValue.ItemType = targetType
                vectorValue.FullType = v.Type
                return true, ""
            }

            msg := "No se puede asignar un vector de tipo " + v.Value.Type() + " a una vector de tipo " + v.Type
            v.Value = tiposDeDato.NuloPorDefecto
            return false, msg
        }

        // Caso especial: vector vacío genérico "[]" asignado a tipo específico
        if v.Value.Type() == "[]" && EsTipoVector(v.Type) {
            if vectorValue, ok := v.Value.(*TipoVector); ok {
                vectorValue.ItemType = EliminarCorchetes(v.Type)
                vectorValue.FullType = v.Type
                return true, ""
            }
        }

        // try implicit primitive conversion
        convertedValue, ok := tiposDeDato.Casteo(v.Type, v.Value)

        if !ok {
            // Si la expresión tiene un tipo de dato diferente al definido previamente se tomará como error y la variable obtendrá el valor de nil para fines prácticos.
            msg := "No se puede asignar un valor de tipo " + v.Value.Type() + " a una variable de tipo " + v.Type
            v.Value = tiposDeDato.NuloPorDefecto
            return false, msg
        }

        v.Value = convertedValue
    }

    return true, ""
}

func (v *Variable) AsignarVariable(val tiposDeDato.ValorInterno, isMutatingContext bool) (bool, string) {

	if v.IsConst {
		return false, "No se puede asignar un valor a una constante"
	}

	if v.isProp {
		if !isMutatingContext {
			return false, "No se puede asignar un valor a una propiedad desde un contexto de función no mutable"
		}
	}

	v.Value = val

	// if obj, ok := val.(*ObjectValue); ok {
	// 	v.Value = obj.Copy()
	// }

	return v.TipoValidacion()
}
