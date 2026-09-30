package files

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func TestListReturnsEntriesSortedByName(t *testing.T) {
	dir := t.TempDir()
	createFile(t, filepath.Join(dir, "b.txt"))
	createFile(t, filepath.Join(dir, "a.png"))
	if err := os.Mkdir(filepath.Join(dir, "sub"), 0o755); err != nil {
		t.Fatalf("crear subdirectorio: %v", err)
	}

	got, err := List(dir)
	if err != nil {
		t.Fatalf("List() error inesperado: %v", err)
	}

	want := []string{"a.png", "b.txt", "sub"}
	if names := namesOf(got); !slices.Equal(names, want) {
		t.Errorf("List() nombres = %v, se esperaba %v", names, want)
	}
}

func TestListEmptyDirectory(t *testing.T) {
	got, err := List(t.TempDir())
	if err != nil {
		t.Fatalf("List() error inesperado: %v", err)
	}
	if len(got) != 0 {
		t.Errorf("List() devolvió %d entradas, se esperaba 0", len(got))
	}
}

func TestListNonexistentDirectory(t *testing.T) {
	_, err := List(filepath.Join(t.TempDir(), "no-existe"))
	if !errors.Is(err, fs.ErrNotExist) {
		t.Errorf("List() error = %v, se esperaba fs.ErrNotExist", err)
	}
}

func createFile(t *testing.T, path string) {
	t.Helper()
	if err := os.WriteFile(path, nil, 0o644); err != nil {
		t.Fatalf("crear archivo %q: %v", path, err)
	}
}

func namesOf(infos []fs.FileInfo) []string {
	names := make([]string, 0, len(infos))
	for _, info := range infos {
		names = append(names, info.Name())
	}
	return names
}

func TestIsImage(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"foto.jpg", true},
		{"paisaje.jpeg", true},
		{"logo.png", true},
		{"captura.PNG", true},
		{"FOTO.JPG", true},
		{"documento.pdf", false},
		{"pagina.html", false},
		{"informe.docx", false},
		{"respaldo.jpg.bak", false},
		{"falso.png.txt", false},
		{"sin-extension", false},
		{"jpg", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsImage(tt.name); got != tt.want {
				t.Errorf("IsImage(%q) = %v, se esperaba %v", tt.name, got, tt.want)
			}
		})
	}
}

func TestImageNamesWithTestDirectory(t *testing.T) {
	got, err := ImageNames(filepath.Join("..", "testdata", "directorio-prueba"))
	if err != nil {
		t.Fatalf("ImageNames() error inesperado: %v", err)
	}

	want := []string{"captura.PNG", "foto.jpg", "logo.png", "paisaje.jpeg"}
	if !slices.Equal(got, want) {
		t.Errorf("ImageNames() = %v, se esperaba %v", got, want)
	}
}

func TestImageNamesExcludesDirectories(t *testing.T) {
	dir := t.TempDir()
	createFile(t, filepath.Join(dir, "foto.jpg"))
	if err := os.Mkdir(filepath.Join(dir, "album.png"), 0o755); err != nil {
		t.Fatalf("crear subdirectorio: %v", err)
	}

	got, err := ImageNames(dir)
	if err != nil {
		t.Fatalf("ImageNames() error inesperado: %v", err)
	}

	want := []string{"foto.jpg"}
	if !slices.Equal(got, want) {
		t.Errorf("ImageNames() = %v, se esperaba %v", got, want)
	}
}

func TestImageNamesWithoutImages(t *testing.T) {
	dir := t.TempDir()
	createFile(t, filepath.Join(dir, "documento.pdf"))

	got, err := ImageNames(dir)
	if err != nil {
		t.Fatalf("ImageNames() error inesperado: %v", err)
	}
	if got == nil || len(got) != 0 {
		t.Errorf("ImageNames() = %#v, se esperaba un slice vacío no nil", got)
	}
}

func TestImageNamesNonexistentDirectory(t *testing.T) {
	_, err := ImageNames(filepath.Join(t.TempDir(), "no-existe"))
	if !errors.Is(err, fs.ErrNotExist) {
		t.Errorf("ImageNames() error = %v, se esperaba fs.ErrNotExist", err)
	}
}
