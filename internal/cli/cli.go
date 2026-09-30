// Package cli contiene utilidades compartidas por los comandos de consola
// del proyecto.
package cli

import (
	"fmt"
	"os"
)

// DirFromArgs devuelve el único argumento recibido, que se interpreta como
// un directorio. args no debe incluir el nombre del programa.
func DirFromArgs(args []string) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf("se esperaba 1 argumento, se recibieron %d", len(args))
	}
	return args[0], nil
}

// Fail imprime err en la salida de errores y termina el programa con
// código 1 (error de ejecución).
func Fail(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}

// FailUsage imprime err y el mensaje de uso en la salida de errores y
// termina el programa con código 2 (uso incorrecto).
func FailUsage(err error, usage string) {
	fmt.Fprintln(os.Stderr, "error:", err)
	fmt.Fprintln(os.Stderr, usage)
	os.Exit(2)
}
