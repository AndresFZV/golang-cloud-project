package main

import "testing"

func TestPreview(t *testing.T) {
	tests := []struct {
		name string
		s    string
		n    int
		want string
	}{
		{name: "más corto que n", s: "abc", n: 5, want: "abc"},
		{name: "igual a n", s: "abcde", n: 5, want: "abcde"},
		{name: "más largo que n", s: "abcdefgh", n: 5, want: "abcde..."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := preview(tt.s, tt.n); got != tt.want {
				t.Errorf("preview(%q, %d) = %q, se esperaba %q", tt.s, tt.n, got, tt.want)
			}
		})
	}
}
