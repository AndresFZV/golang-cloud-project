package main

import (
	"errors"
	"flag"
	"io"
	"testing"
)

func TestParseConfig(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    config
		wantErr bool
	}{
		{name: "puerto y directorio", args: []string{"-port", "9090", "-dir", "imagenes"}, want: config{port: 9090, dir: "imagenes"}},
		{name: "puerto por defecto", args: []string{"-dir", "imagenes"}, want: config{port: 8080, dir: "imagenes"}},
		{name: "sin directorio", args: []string{"-port", "9090"}, wantErr: true},
		{name: "puerto cero", args: []string{"-port", "0", "-dir", "imagenes"}, wantErr: true},
		{name: "puerto fuera de rango", args: []string{"-port", "70000", "-dir", "imagenes"}, wantErr: true},
		{name: "puerto no numérico", args: []string{"-port", "abc", "-dir", "imagenes"}, wantErr: true},
		{name: "argumento extra", args: []string{"-dir", "imagenes", "sobra"}, wantErr: true},
		{name: "flag desconocido", args: []string{"-dir", "imagenes", "-x"}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseConfig(tt.args, io.Discard)
			if (err != nil) != tt.wantErr {
				t.Fatalf("parseConfig(%v) error = %v, wantErr %v", tt.args, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("parseConfig(%v) = %+v, se esperaba %+v", tt.args, got, tt.want)
			}
		})
	}
}

func TestParseConfigHelp(t *testing.T) {
	_, err := parseConfig([]string{"-h"}, io.Discard)
	if !errors.Is(err, flag.ErrHelp) {
		t.Errorf("parseConfig(-h) error = %v, se esperaba flag.ErrHelp", err)
	}
}
