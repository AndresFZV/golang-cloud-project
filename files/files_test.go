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
