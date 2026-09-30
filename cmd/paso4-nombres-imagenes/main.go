// Command paso4-nombres-imagenes muestra los nombres de los archivos de
// imagen (.jpg, .jpeg, .png) del directorio indicado, uno por línea.
// Corresponde al Paso 4 de la Parte 1 del proyecto.
//
// Uso:
//
//	paso4-nombres-imagenes <directorio>
package main

import (
	"fmt"
	"os"

	"github.com/AndresFZV/golang-cloud-project/files"
)

const usage = "uso: paso4-nombres-imagenes <directorio>"

func main() {
	dir, err := dirFromArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(2)
	}

	names, err := files.ImageNames(dir)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	for _, name := range names {
		fmt.Println(name)
	}
}

// dirFromArgs devuelve el directorio a listar a partir de los argumentos
// recibidos, sin incluir el nombre del programa.
func dirFromArgs(args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("se esperaba 1 argumento, se recibieron %d", len(args))
	}
	return args[0], nil
}
