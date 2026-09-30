package gallery

import (
	"encoding/base64"
	"errors"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

var testDir = filepath.Join("..", "testdata", "directorio-prueba")

func TestPickRandomReturnsDistinctItemsFromInput(t *testing.T) {
	items := []string{"a", "b", "c", "d", "e"}

	for range 50 {
		got, err := PickRandom(items, 3)
		if err != nil {
			t.Fatalf("PickRandom() error inesperado: %v", err)
		}
		if len(got) != 3 {
			t.Fatalf("PickRandom() devolvió %d elementos, se esperaban 3", len(got))
		}

		seen := make(map[string]bool)
		for _, item := range got {
			if !slices.Contains(items, item) {
				t.Fatalf("%q no pertenece a la entrada %v", item, items)
			}
			if seen[item] {
				t.Fatalf("%q está repetido en %v", item, got)
			}
			seen[item] = true
		}
	}
}

func TestPickRandomDoesNotModifyInput(t *testing.T) {
	items := []string{"a", "b", "c"}
	original := slices.Clone(items)

	if _, err := PickRandom(items, 3); err != nil {
		t.Fatalf("PickRandom() error inesperado: %v", err)
	}
	if !slices.Equal(items, original) {
		t.Errorf("PickRandom() modificó la entrada: %v, se esperaba %v", items, original)
	}
}

func TestPickRandomZero(t *testing.T) {
	got, err := PickRandom([]string{"a"}, 0)
	if err != nil {
		t.Fatalf("PickRandom() error inesperado: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("PickRandom() devolvió %v, se esperaba un slice vacío", got)
	}
}

func TestPickRandomInvalidCount(t *testing.T) {
	items := []string{"a", "b"}

	if _, err := PickRandom(items, 3); !errors.Is(err, ErrNotEnoughItems) {
		t.Errorf("PickRandom(n > len) error = %v, se esperaba ErrNotEnoughItems", err)
	}
	if _, err := PickRandom(items, -1); err == nil {
		t.Error("PickRandom(n < 0) no devolvió error")
	}
}

func TestEncodeRejectsNonImage(t *testing.T) {
	if _, err := Encode(testDir, "documento.pdf"); err == nil {
		t.Error("Encode() con un .pdf no devolvió error")
	}
}

func TestRandomImages(t *testing.T) {
	images, err := RandomImages(testDir, 4)
	if err != nil {
		t.Fatalf("RandomImages() error inesperado: %v", err)
	}
	if len(images) != 4 {
		t.Fatalf("RandomImages() devolvió %d imágenes, se esperaban 4", len(images))
	}

	for _, img := range images {
		if img.MIMEType != "image/jpeg" && img.MIMEType != "image/png" {
			t.Errorf("%s: tipo MIME inesperado %q", img.Name, img.MIMEType)
		}
		if _, err := base64.StdEncoding.DecodeString(img.Base64); err != nil {
			t.Errorf("%s: Base64 inválido: %v", img.Name, err)
		}
		if !strings.HasPrefix(img.DataURI(), "data:"+img.MIMEType+";base64,") {
			t.Errorf("%s: DataURI con formato inválido", img.Name)
		}
	}
}

func TestRandomImagesNotEnough(t *testing.T) {
	_, err := RandomImages(testDir, 5)
	if !errors.Is(err, ErrNotEnoughItems) {
		t.Errorf("RandomImages() error = %v, se esperaba ErrNotEnoughItems", err)
	}
}
