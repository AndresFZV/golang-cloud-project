package cli

import "testing"

func TestDirFromArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    string
		wantErr bool
	}{
		{name: "un argumento", args: []string{"imagenes"}, want: "imagenes"},
		{name: "sin argumentos", args: []string{}, wantErr: true},
		{name: "dos argumentos", args: []string{"a", "b"}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DirFromArgs(tt.args)
			if (err != nil) != tt.wantErr {
				t.Fatalf("DirFromArgs(%v) error = %v, wantErr %v", tt.args, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("DirFromArgs(%v) = %q, se esperaba %q", tt.args, got, tt.want)
			}
		})
	}
}
