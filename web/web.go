// Package web contiene el handler HTTP de la galería: muestra el nombre del
// host y un conjunto de imágenes elegidas al azar, codificadas en Base64.
package web

import (
	"bytes"
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/AndresFZV/golang-cloud-project/files"
	"github.com/AndresFZV/golang-cloud-project/gallery"
)

//go:embed templates/index.html
var templatesFS embed.FS

// Config define los parámetros del handler.
type Config struct {
	ImagesDir  string
	Hostname   string
	ImageCount int
}

type handler struct {
	cfg  Config
	tmpl *template.Template
}

type pageData struct {
	Hostname  string
	Directory string
	Images    []imageView
}

type imageView struct {
	Name string
	Src  template.URL
}

// NewHandler valida la configuración, carga la plantilla y devuelve el
// handler de la aplicación. Falla si el directorio no existe o no tiene al
// menos cfg.ImageCount imágenes.
func NewHandler(cfg Config) (http.Handler, error) {
	if err := checkImages(cfg.ImagesDir, cfg.ImageCount); err != nil {
		return nil, err
	}

	tmpl, err := template.ParseFS(templatesFS, "templates/index.html")
	if err != nil {
		return nil, fmt.Errorf("cargar plantilla: %w", err)
	}

	h := &handler{cfg: cfg, tmpl: tmpl}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", h.index)
	return mux, nil
}

func (h *handler) index(w http.ResponseWriter, _ *http.Request) {
	images, err := gallery.RandomImages(h.cfg.ImagesDir, h.cfg.ImageCount)
	if err != nil {
		log.Printf("seleccionar imágenes: %v", err)
		http.Error(w, "no fue posible cargar las imágenes", http.StatusInternalServerError)
		return
	}

	data := pageData{
		Hostname:  h.cfg.Hostname,
		Directory: h.cfg.ImagesDir,
		Images:    toViews(images),
	}

	var buf bytes.Buffer
	if err := h.tmpl.Execute(&buf, data); err != nil {
		log.Printf("generar página: %v", err)
		http.Error(w, "no fue posible generar la página", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if _, err := buf.WriteTo(w); err != nil {
		log.Printf("escribir respuesta: %v", err)
	}
}

func toViews(images []gallery.Image) []imageView {
	views := make([]imageView, 0, len(images))
	for _, img := range images {
		// El URI se arma con un tipo MIME fijo y datos Base64 generados por
		// la aplicación, por eso es seguro marcarlo como URL confiable.
		views = append(views, imageView{Name: img.Name, Src: template.URL(img.DataURI())})
	}
	return views
}

func checkImages(dir string, count int) error {
	names, err := files.ImageNames(dir)
	if err != nil {
		return err
	}
	if len(names) < count {
		return fmt.Errorf("el directorio %q tiene %d imágenes y se necesitan %d: %w",
			dir, len(names), count, gallery.ErrNotEnoughItems)
	}
	return nil
}
