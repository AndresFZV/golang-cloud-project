// Package files provee funciones para leer directorios y trabajar con
// archivos. Las funciones devuelven datos y errores; no imprimen nada.
package files

import (
	"fmt"
	"io/fs"
	"os"
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
