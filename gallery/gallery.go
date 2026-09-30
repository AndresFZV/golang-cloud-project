// Package gallery selecciona imágenes al azar de un directorio y las
// codifica en Base64. No imprime nada ni depende de HTTP.
package gallery

import (
	"errors"
	"fmt"
	"math/rand/v2"
	"path/filepath"

	"github.com/AndresFZV/golang-cloud-project/files"
)

// ErrNotEnoughItems indica que se pidieron más elementos de los disponibles.
var ErrNotEnoughItems = errors.New("no hay suficientes elementos")

// Image es una imagen codificada en Base64.
type Image struct {
	Name     string
	MIMEType string
	Base64   string
}

// DataURI devuelve la imagen como URI de datos (data:<mime>;base64,<datos>),
// lista para usarse en el atributo src de una etiqueta img.
func (img Image) DataURI() string {
	return "data:" + img.MIMEType + ";base64," + img.Base64
}

// PickRandom devuelve n elementos distintos de items elegidos al azar.
// No modifica items.
func PickRandom(items []string, n int) ([]string, error) {
	if n < 0 {
		return nil, fmt.Errorf("cantidad inválida: %d", n)
	}
	if n > len(items) {
		return nil, fmt.Errorf("se pidieron %d de %d: %w", n, len(items), ErrNotEnoughItems)
	}

	picked := make([]string, 0, n)
	for _, i := range rand.Perm(len(items))[:n] {
		picked = append(picked, items[i])
	}
	return picked, nil
}

// Encode lee la imagen name del directorio dir y la codifica en Base64.
func Encode(dir, name string) (Image, error) {
	mimeType := files.ImageMIMEType(name)
	if mimeType == "" {
		return Image{}, fmt.Errorf("%q no es una imagen soportada", name)
	}

	data, err := files.ReadBase64(filepath.Join(dir, name))
	if err != nil {
		return Image{}, err
	}

	return Image{Name: name, MIMEType: mimeType, Base64: data}, nil
}

// RandomImages elige n imágenes distintas al azar del directorio dir y las
// devuelve codificadas en Base64.
func RandomImages(dir string, n int) ([]Image, error) {
	names, err := files.ImageNames(dir)
	if err != nil {
		return nil, err
	}

	picked, err := PickRandom(names, n)
	if err != nil {
		return nil, err
	}

	images := make([]Image, 0, n)
	for _, name := range picked {
		img, err := Encode(dir, name)
		if err != nil {
			return nil, err
		}
		images = append(images, img)
	}
	return images, nil
}
