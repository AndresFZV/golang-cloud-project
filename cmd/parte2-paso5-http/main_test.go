package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloPage(t *testing.T) {
	rec := httptest.NewRecorder()
	newMux().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/", nil))

	if rec.Code != http.StatusOK {
		t.Fatalf("código = %d, se esperaba %d", rec.Code, http.StatusOK)
	}
	if !strings.Contains(rec.Body.String(), "Hola Mundo!") {
		t.Errorf("el cuerpo no contiene \"Hola Mundo!\": %q", rec.Body.String())
	}
}

func TestUnknownRouteReturnsNotFound(t *testing.T) {
	rec := httptest.NewRecorder()
	newMux().ServeHTTP(rec, httptest.NewRequest(http.MethodGet, "/otra", nil))

	if rec.Code != http.StatusNotFound {
		t.Errorf("código = %d, se esperaba %d", rec.Code, http.StatusNotFound)
	}
}
