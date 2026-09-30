package main

import "testing"

func TestDirFromArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    string
		wantErr bool
	}{
		{name: "un argumento", args: []string{"testdata"}, want: "testdata"},
		{name: "sin argumentos", args: []string{}, wantErr: true},
		{name: "dos argumentos", args: []string{"a", "b"}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := dirFromArgs(tt.args)
			if (err != nil) != tt.wantErr {
				t.Fatalf("dirFromArgs(%v) error = %v, wantErr %v", tt.args, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("dirFromArgs(%v) = %q, se esperaba %q", tt.args, got, tt.want)
			}
		})
	}
}
