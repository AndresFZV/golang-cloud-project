// Command parte2-paso1-array guarda en un slice los nombres de las imágenes
// del directorio indicado y muestra cuántas hay.
// Corresponde al Paso 1 de la Parte 2 del proyecto.
//
// Uso:
//
//	parte2-paso1-array <directorio>
package main

import (
	"fmt"
	"os"

	"github.com/AndresFZV/golang-cloud-project/files"
	"github.com/AndresFZV/golang-cloud-project/internal/cli"
)

const usage = "uso: parte2-paso1-array <directorio>"

func main() {
	dir, err := cli.DirFromArgs(os.Args[1:])
	if err != nil {
		cli.FailUsage(err, usage)
	}

	names, err := files.ImageNames(dir)
	if err != nil {
		cli.Fail(err)
	}

	for i, name := range names {
		fmt.Printf("[%d] %s\n", i, name)
	}
	fmt.Printf("\nCantidad de imágenes: %d\n", len(names))
}
