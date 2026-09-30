// Command parte2-paso3-base64 selecciona al azar una imagen del directorio
// indicado y la codifica en Base64.
// Corresponde al Paso 3 de la Parte 2 del proyecto.
//
// Uso:
//
//	parte2-paso3-base64 <directorio>
package main

import (
	"fmt"
	"os"

	"github.com/AndresFZV/golang-cloud-project/gallery"
	"github.com/AndresFZV/golang-cloud-project/internal/cli"
)

const (
	usage         = "uso: parte2-paso3-base64 <directorio>"
	previewLength = 76
)

func main() {
	dir, err := cli.DirFromArgs(os.Args[1:])
	if err != nil {
		cli.FailUsage(err, usage)
	}

	images, err := gallery.RandomImages(dir, 1)
	if err != nil {
		cli.Fail(fmt.Errorf("seleccionar y codificar imagen: %w", err))
	}
	img := images[0]

	fmt.Printf("Imagen seleccionada: %s\n", img.Name)
	fmt.Printf("Tipo MIME: %s\n", img.MIMEType)
	fmt.Printf("Longitud en Base64: %d caracteres\n", len(img.Base64))
	fmt.Printf("Inicio del Base64: %s\n", preview(img.Base64, previewLength))
}

// preview devuelve los primeros n caracteres de s seguidos de "..." si s
// es más largo que n.
func preview(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}
