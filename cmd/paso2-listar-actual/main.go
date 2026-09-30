// Command paso2-listar-actual lista las entradas del directorio de trabajo
// actual mostrando su tipo, tamaño y nombre.
// Corresponde al Paso 2 de la Parte 1 del proyecto.
package main

import (
	"fmt"
	"io/fs"
	"os"

	"github.com/AndresFZV/golang-cloud-project/files"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		exitWithError(err)
	}

	entries, err := files.List(dir)
	if err != nil {
		exitWithError(err)
	}

	fmt.Printf("Directorio: %s\n\n", dir)
	for _, info := range entries {
		fmt.Println(formatEntry(info))
	}
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
