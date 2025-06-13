package instrucciones

import (
	"main/tiposDeDato"

	"github.com/antlr4-go/antlr/v4"
)

type TipoObjeto struct {
	InternalScope *AmbitoBase
	AuxObject     interface{}
	ConcretType   string
	v             *PatronVIsitor
	t             antlr.Token
}

func (o TipoObjeto) Value() interface{} {
	return o
}

func (o TipoObjeto) Type() string {
	if o.ConcretType != "" {
		return o.ConcretType
	}

	return tiposDeDato.TIPO_OBJECTO
}

func (o *TipoObjeto) Copy() tiposDeDato.ValorInterno {
	args := make([]*Argumento, 0)

	for _, prop := range o.InternalScope.variables {
		args = append(args, &Argumento{
			Name:  prop.Name,
			Value: prop.Value,
		})
	}

	return NewTipoObjeto(o.v, o.ConcretType, o.t, args, true)
}

func NewTipoObjeto(v *PatronVIsitor, targetStruct string, targetToken antlr.Token, args []*Argumento, allowReinitialize bool) tiposDeDato.ValorInterno {

    // Check if struct exists
    structTemplate, msg := v.RegistroAmbito.AmbitoGlobal.GetEstructura(targetStruct)

    if structTemplate == nil {
        v.TablaError.NewErrorSemantico(targetToken, msg)
        return tiposDeDato.NuloPorDefecto
    }

    internalScope := NuevoAmbitoGlobal()

    prevScope := v.RegistroAmbito.AmbitoActual
    v.RegistroAmbito.AmbitoActual = internalScope

    defer func() {
        // restore scope
        v.RegistroAmbito.AmbitoActual = prevScope
    }()

    // MEJORADO: Add fields to internal scope con verificaciones
    for _, field := range structTemplate.Fields {
        if field != nil { // Verificar que field no sea nil
            result := v.Visit(field)
            if result == nil {
                // Si hay error procesando el campo, continuar con el siguiente
                continue
            }
        }
    }

    // Resto del código igual...
    // create args map
    argMap := make(map[string]*Argumento)

    for _, arg := range args {
        // repeat arg
        if _, ok := argMap[arg.Name]; ok {
            v.TablaError.NewErrorSemantico(arg.Token, "El argumento "+arg.Name+" ya fue definido")
            return tiposDeDato.NuloPorDefecto
        }

        argMap[arg.Name] = arg
    }

    // validate constructor args
    wasConst := false
    usedArgs := make(map[string]bool)

    for _, prop := range internalScope.variables {
        arg, found := argMap[prop.Name]

        if !found {
            if prop.Value == tiposDeDato.ValorNoIniPorDefecto {
                v.TablaError.NewErrorSemantico(targetToken, "El campo "+prop.Name+" no fue inicializado en el constructor")
                return tiposDeDato.NuloPorDefecto
            }

            continue
        }

        // then the arg exists
        if prop.IsConst {
            if (prop.Value != tiposDeDato.ValorNoIniPorDefecto) && !allowReinitialize {
                v.TablaError.NewErrorSemantico(targetToken, "El campo "+prop.Name+" es inmutable y ya fue inicializado")
                return tiposDeDato.NuloPorDefecto
            }

            wasConst = true
            prop.IsConst = false
        }

        var throwError bool = false
        var msg string = ""
        var assignValue tiposDeDato.ValorInterno = arg.Value

        // pointer support
        if arg.PassByReference {
            if arg.VariableRef == nil {
                msg = "No es posible pasar por referencia un valor que no este asociado a una variable"
                throwError = true
            }

            // create the pointer
            assignValue = &TipoPuntero{
                VariableAsociada: arg.VariableRef,
            }
        }

        if !throwError {
            throwError, msg = prop.AsignarVariable(assignValue, false)
        }

        if wasConst {
            prop.IsConst = true
            wasConst = false
        }

        if !throwError {
            v.TablaError.NewErrorSemantico(targetToken, msg)
            return tiposDeDato.NuloPorDefecto
        }

        usedArgs[prop.Name] = true
    }

    // validate unused args
    for _, arg := range args {
        if _, ok := usedArgs[arg.Name]; !ok {
            v.TablaError.NewErrorSemantico(arg.Token, "El argumento "+arg.Name+" no es utilizado en el constructor")
        }
    }

    // mutable related flags
    for _, prop := range internalScope.variables {
        prop.isProp = true
    }

    // self object
    selfObject := &TipoObjeto{
        InternalScope: internalScope,
        ConcretType:   tiposDeDato.TIPO_PROPIO,
    }

    instanceInternalScope := NuevoAmbitoGlobal()

    instanceInternalScope.AgregarVariable("self", tiposDeDato.TIPO_PROPIO, selfObject, true, false, nil)

    // make functions use the instance scope
    for _, function := range internalScope.functions {
        f, ok := function.(*Funcion)

        if !ok {
            continue
        }

        f.DefaultScope = instanceInternalScope
    }

    // create instance
    return &TipoObjeto{
        InternalScope: internalScope,
        ConcretType:   targetStruct,
        v:             v,
        t:             targetToken,
    }
}

func ArgumentoValidoEstructura(arg []*Argumento) bool {
	for _, a := range arg {

		if a.Name == "" {
			return false
		}
	}
	return true
}
