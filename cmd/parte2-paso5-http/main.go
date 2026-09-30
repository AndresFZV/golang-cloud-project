// Command parte2-paso5-http inicia un servidor HTTP en el puerto 8080 con
// una página de prueba que muestra "Hola Mundo!".
// Corresponde al Paso 5 de la Parte 2 del proyecto.
package main

import (
	"io"
	"log"
	"net/http"
	"time"
)

const (
	addr      = ":8080"
	helloPage = `<!DOCTYPE html>
<html lang="es">
<head><meta charset="utf-8"><title>Hola Mundo</title></head>
<body><h1>Hola Mundo!</h1></body>
</html>`
)

func main() {
	srv := &http.Server{
		Addr:              addr,
		Handler:           newMux(),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Printf("servidor escuchando en http://localhost%s", addr)
	log.Fatal(srv.ListenAndServe())
}

func newMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", helloHandler)
	return mux
}

func helloHandler(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := io.WriteString(w, helloPage); err != nil {
		log.Printf("escribir respuesta: %v", err)
	}
}
