// Command parte2-web inicia la aplicación web de la galería. Muestra el
// nombre del host y 4 imágenes elegidas al azar del directorio indicado,
// codificadas en Base64. Cubre los pasos 6 a 11 de la Parte 2.
//
// Uso:
//
//	parte2-web -port 8080 -dir imagenes/coleccion-1
package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/AndresFZV/golang-cloud-project/web"
)

const imagesPerPage = 4

type config struct {
	port int
	dir  string
}

func main() {
	cfg, err := parseConfig(os.Args[1:], os.Stderr)
	if errors.Is(err, flag.ErrHelp) {
		return
	}
	if err != nil {
		os.Exit(2) // parseConfig ya imprimió el error y el uso.
	}

	if err := run(cfg); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// parseConfig lee los argumentos de la línea de comandos. Los errores y el
// mensaje de uso se escriben en output.
func parseConfig(args []string, output io.Writer) (config, error) {
	flags := flag.NewFlagSet("parte2-web", flag.ContinueOnError)
	flags.SetOutput(output)
	port := flags.Int("port", 8080, "puerto HTTP (1-65535)")
	dir := flags.String("dir", "", "directorio con las imágenes (obligatorio)")

	if err := flags.Parse(args); err != nil {
		return config{}, err
	}

	cfg := config{port: *port, dir: *dir}
	if err := cfg.validate(flags.NArg()); err != nil {
		fmt.Fprintln(output, "error:", err)
		flags.Usage()
		return config{}, err
	}
	return cfg, nil
}

func (c config) validate(extraArgs int) error {
	if extraArgs > 0 {
		return fmt.Errorf("argumentos no reconocidos: %d", extraArgs)
	}
	if c.dir == "" {
		return errors.New("falta el directorio de imágenes (-dir)")
	}
	if c.port < 1 || c.port > 65535 {
		return fmt.Errorf("puerto inválido %d: debe estar entre 1 y 65535", c.port)
	}
	return nil
}

func run(cfg config) error {
	hostname, err := os.Hostname()
	if err != nil {
		return fmt.Errorf("obtener nombre del host: %w", err)
	}

	handler, err := web.NewHandler(web.Config{
		ImagesDir:  cfg.dir,
		Hostname:   hostname,
		ImageCount: imagesPerPage,
	})
	if err != nil {
		return err
	}

	srv := &http.Server{
		Addr:              net.JoinHostPort("", strconv.Itoa(cfg.port)),
		Handler:           handler,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("servidor escuchando en http://localhost:%d (host: %s, directorio: %s)",
		cfg.port, hostname, cfg.dir)
	return srv.ListenAndServe()
}
