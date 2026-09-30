package web

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"

	"github.com/AndresFZV/golang-cloud-project/gallery"
)

var testDir = filepath.Join("..", "testdata", "directorio-prueba")

func newTestHandler(t *testing.T) http.Handler {
	t.Helper()
	h, err := NewHandler(Config{ImagesDir: testDir, Hostname: "host-prueba", ImageCount: 4})
	if err != nil {
		t.Fatalf("NewHandler() error inesperado: %v", err)
	}
	return h
}

func TestIndexRendersHostnameAndImages(t *testing.T) {
	rec := httptest.NewRecorder()
	newTestHandler(t).ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("código = %d, se esperaba %d", rec.Code, http.StatusOK)
	}
	if ct := rec.Header().Get("Content-Type"); !strings.HasPrefix(ct, "text/html") {
		t.Errorf("Content-Type = %q, se esperaba text/html", ct)
	}

	body := rec.Body.String()
	if !strings.Contains(body, "host-prueba") {
		t.Error("la página no contiene el nombre del host")
	}
	if got := strings.Count(body, `src="data:image/`); got != 4 {
		t.Errorf("la página tiene %d imágenes en Base64, se esperaban 4", got)
	}
	if got := strings.Count(body, "<figcaption>"); got != 4 {
		t.Errorf("la página tiene %d nombres de imagen, se esperaban 4", got)
	}
}

func TestUnknownRouteReturnsNotFound(t *testing.T) {
	rec := httptest.NewRecorder()
	newTestHandler(t).ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/otra", nil))

	if rec.Code != http.StatusNotFound {
		t.Errorf("código = %d, se esperaba %d", rec.Code, http.StatusNotFound)
	}
}

func TestPostReturnsMethodNotAllowed(t *testing.T) {
	rec := httptest.NewRecorder()
	newTestHandler(t).ServeHTTP(rec, httptest.NewRequest(http.MethodPost, "/", nil))

	if rec.Code != http.StatusMethodNotAllowed {
		t.Errorf("código = %d, se esperaba %d", rec.Code, http.StatusMethodNotAllowed)
	}
}

func TestNewHandlerNotEnoughImages(t *testing.T) {
	_, err := NewHandler(Config{ImagesDir: testDir, Hostname: "h", ImageCount: 5})
	if !errors.Is(err, gallery.ErrNotEnoughItems) {
		t.Errorf("NewHandler() error = %v, se esperaba ErrNotEnoughItems", err)
	}
}

func TestNewHandlerNonexistentDirectory(t *testing.T) {
	_, err := NewHandler(Config{ImagesDir: "no-existe", Hostname: "h", ImageCount: 4})
	if err == nil {
		t.Error("NewHandler() con un directorio inexistente no devolvió error")
	}
}
