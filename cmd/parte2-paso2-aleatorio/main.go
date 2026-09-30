// Command parte2-paso2-aleatorio selecciona al azar el nombre de una de las
// imágenes del directorio indicado.
// Corresponde al Paso 2 de la Parte 2 del proyecto.
//
// Uso:
//
//	parte2-paso2-aleatorio <directorio>
package main

import (
	"fmt"
	"os"

	"github.com/AndresFZV/golang-cloud-project/files"
	"github.com/AndresFZV/golang-cloud-project/gallery"
	"github.com/AndresFZV/golang-cloud-project/internal/cli"
)

const usage = "uso: parte2-paso2-aleatorio <directorio>"

func main() {
	dir, err := cli.DirFromArgs(os.Args[1:])
	if err != nil {
		cli.FailUsage(err, usage)
	}

	names, err := files.ImageNames(dir)
	if err != nil {
		cli.Fail(err)
	}

	picked, err := gallery.PickRandom(names, 1)
	if err != nil {
		cli.Fail(fmt.Errorf("seleccionar imagen: %w", err))
	}

	fmt.Printf("Imágenes disponibles: %d\n", len(names))
	fmt.Printf("Imagen seleccionada: %s\n", picked[0])
}
