// Command parte2-paso4-hostname imprime el nombre del host del computador
// donde se ejecuta la aplicación.
// Corresponde al Paso 4 de la Parte 2 del proyecto.
package main

import (
	"fmt"
	"os"

	"github.com/AndresFZV/golang-cloud-project/internal/cli"
)

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		cli.Fail(fmt.Errorf("obtener nombre del host: %w", err))
	}
	fmt.Printf("Nombre del host: %s\n", hostname)
}
