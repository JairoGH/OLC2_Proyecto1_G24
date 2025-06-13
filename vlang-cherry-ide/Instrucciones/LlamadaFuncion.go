package instrucciones

import "main/tiposDeDato"

const (
	Detener   = "break"
	Continuar = "continue"
	Retornar  = "return"
)

type LlamadaFunciones struct {
	RetornarValor tiposDeDato.ValorInterno
	Type          []string
	Accion        string
}

func (csi *LlamadaFunciones) EsDeTipo(t string) bool {
	for _, i := range csi.Type {
		if i == t {
			return true
		}
	}
	return false
}

func (csi *LlamadaFunciones) EsAccion(a string) bool {
	return csi.Accion == a
}

func (csi *LlamadaFunciones) ResetAccion() {
	csi.Accion = ""
}

type PilaLlamada struct {
	Items []*LlamadaFunciones
}

func (cs *PilaLlamada) Push(item *LlamadaFunciones) {
	cs.Items = append(cs.Items, item)
}

func (cs *PilaLlamada) Pop() *LlamadaFunciones {
	item := cs.Items[len(cs.Items)-1]
	cs.Items = cs.Items[:len(cs.Items)-1]
	return item
}

func (cs *PilaLlamada) Peek() *LlamadaFunciones {
	return cs.Items[len(cs.Items)-1]
}

func (cs *PilaLlamada) In(item *LlamadaFunciones) bool {
	for _, i := range cs.Items {
		if i == item {
			return true
		}
	}
	return false
}

// Elimina elementos de la pila hasta encontrar el elemento buscado
func (cs *PilaLlamada) Limpiar(item *LlamadaFunciones) {
	if !cs.In(item) {
		return
	}

	for {
		peek := cs.Pop()
		if peek == item {
			break
		}
	}
}

func (cs *PilaLlamada) PuedeContinuar() (bool, *LlamadaFunciones) {
	// La instrucción 'continue' solo puede estar en un bucle
	// No puede atravesar un entorno de retorno de función
	start := len(cs.Items) - 1

	for i := start; i >= 0; i-- {
		if cs.Items[i].EsDeTipo(Continuar) {
			return true, cs.Items[i]
		}

		if cs.Items[i].EsDeTipo(Retornar) {
			return false, nil
		}
	}

	return false, nil
}

func (cs *PilaLlamada) PuedeDetener() (bool, *LlamadaFunciones) {
	// El 'break' debe ser el elemento superior de la pila
	if len(cs.Items) == 0 {
		return false, nil
	}

	if cs.Items[len(cs.Items)-1].EsDeTipo(Detener) {
		return true, cs.Items[len(cs.Items)-1]
	}

	return false, nil
}

func (cs *PilaLlamada) PuedeRetornar() (bool, *LlamadaFunciones) {
	// El 'return' puede interferir con cualquier otro elemento
	for i := len(cs.Items) - 1; i >= 0; i-- {
		if cs.Items[i].EsDeTipo(Retornar) {
			return true, cs.Items[i]
		}
	}

	return false, nil
}

func (cs *PilaLlamada) Len() int {
	return len(cs.Items)
}

func NuevaLlamadaFuncion() *PilaLlamada {
	return &PilaLlamada{Items: []*LlamadaFunciones{}}
}
