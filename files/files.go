// Package files provee funciones para leer directorios y trabajar con
// archivos. Las funciones devuelven datos y errores; no imprimen nada.
package files

import (
	"encoding/base64"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// List devuelve la información de las entradas del directorio dir,
// ordenadas por nombre. Incluye archivos y subdirectorios.
func List(dir string) ([]fs.FileInfo, error) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("leer directorio: %w", err)
	}

	infos := make([]fs.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return nil, fmt.Errorf("obtener información de archivo: %w", err)
		}
		infos = append(infos, info)
	}

	return infos, nil
}

// ImageMIMEType devuelve el tipo MIME de una imagen según su extensión
// (.jpg, .jpeg o .png), o una cadena vacía si no es una imagen soportada.
// La comparación no distingue mayúsculas de minúsculas.
func ImageMIMEType(name string) string {
	switch strings.ToLower(filepath.Ext(name)) {
	case ".jpg", ".jpeg":
		return "image/jpeg"
	case ".png":
		return "image/png"
	default:
		return ""
	}
}

// IsImage informa si name tiene extensión .jpg, .jpeg o .png.
// La comparación no distingue mayúsculas de minúsculas.
func IsImage(name string) bool {
	return ImageMIMEType(name) != ""
}

// ImageNames devuelve los nombres de los archivos de imagen del directorio
// dir, ordenados por nombre. Los subdirectorios se excluyen aunque su nombre
// tenga extensión de imagen.
func ImageNames(dir string) ([]string, error) {
	infos, err := List(dir)
	if err != nil {
		return nil, err
	}

	names := make([]string, 0, len(infos))
	for _, info := range infos {
		if !info.IsDir() && IsImage(info.Name()) {
			names = append(names, info.Name())
		}
	}

	return names, nil
}

// ReadBase64 lee el archivo path y devuelve su contenido codificado en
// Base64 estándar (RFC 4648).
func ReadBase64(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("leer archivo: %w", err)
	}
	return base64.StdEncoding.EncodeToString(data), nil
}
