package creararbol

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
)

type RespuestaArbol struct {
	SVGTree string `json:"svgtree"`
}

func LeerArchivo(nombreArchivo string) string {
	archivo, err := os.Open(nombreArchivo)
	if err != nil {
		panic(err)
	}
	defer archivo.Close()

	contenido, _ := io.ReadAll(archivo)
	return string(contenido)
}

func ReporteArbol(entrada string) string {

	// obtener el contenido 

	contenidoParser := ""

	_, nombreArchivo, _, _ := runtime.Caller(0)

	ruta := filepath.Dir(nombreArchivo)

	// quitar \cst de la ruta
	ruta = ruta[:len(ruta)-4]

	parser, err := json.Marshal(LeerArchivo(filepath.Join(ruta, "/compiler/VGrammar.g4")))

	if err != nil {
		fmt.Println(err)
	}
	contenidoParser = string(parser)

	contenidoLexer := ""
	lexer, err := json.Marshal(LeerArchivo(filepath.Join(ruta, "/compiler/VLexer.g4")))

	if err != nil {
		fmt.Println(err)
	}

	contenidoLexer = string(lexer)

	entradaJ, err := json.Marshal(entrada)
	entradaF := string(entradaJ)

	carga := []byte(
		fmt.Sprintf(
			`{
				"grammar": %s,
				"input": %s,
				"lexgrammar": %s,
				"start": "%s"
			}`,
			contenidoParser,
			entradaF,
			contenidoLexer,
			"program",
		),
	)

	req, err := http.NewRequest("POST", "http://lab.antlr.org/parse/", bytes.NewBuffer(carga))

	if err != nil {
		fmt.Println("Error al crear la solicitud:", err)
		return ""
	}
	req.Header.Set("Content-Type", "application/json")

	cliente := &http.Client{}
	resp, err := cliente.Do(req)
	if err != nil {
		fmt.Println("Error al enviar la solicitud:", err)
		return ""
	}
	defer resp.Body.Close()

	cuerpo, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer el cuerpo:", err)
		return ""
	}

	// crear un mapa para almacenar el json
	var datos map[string]interface{}

	// deserializar el json
	err = json.Unmarshal(cuerpo, &datos)
	if err != nil {
		fmt.Println("Error al deserializar json:", err)
		return ""
	}

	resultado := datos["result"].(map[string]interface{})

	return resultado["svgtree"].(string)
}
