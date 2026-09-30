// Command paso3-listar-directorio lista las entradas del directorio indicado
// como argumento, mostrando su tipo, tamaño y nombre.
// Corresponde al Paso 3 de la Parte 1 del proyecto.
//
// Uso:
//
//	paso3-listar-directorio <directorio>
package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/AndresFZV/golang-cloud-project/files"
)

const usage = "uso: paso3-listar-directorio <directorio>"

func main() {
	dir, err := dirFromArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		fmt.Fprintln(os.Stderr, usage)
		os.Exit(2)
	}

	absDir, err := filepath.Abs(dir)
	if err != nil {
		exitWithError(err)
	}

	entries, err := files.List(absDir)
	if err != nil {
		exitWithError(err)
	}

	fmt.Printf("Directorio: %s\n\n", absDir)
	for _, info := range entries {
		fmt.Println(formatEntry(info))
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

func formatEntry(info fs.FileInfo) string {
	kind := "archivo"
	if info.IsDir() {
		kind = "directorio"
	}
	return fmt.Sprintf("%-10s %10d bytes  %s", kind, info.Size(), info.Name())
}

func exitWithError(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}
