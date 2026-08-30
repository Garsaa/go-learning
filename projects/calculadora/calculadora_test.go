package calculadora

import "testing"

func TestSomar(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{name: "positivos", a: 2, b: 3, want: 5},
		{name: "negativos", a: -2, b: -3, want: -5},
		{name: "com zero", a: 4, b: 0, want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Somar(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("Somar(%d, %d) = %d; esperado %d", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDobrar(t *testing.T) {
	got := Dobrar(4)
	want := 8

	if got != want {
		t.Errorf("Dobrar(4) = %d; esperado %d", got, want)
	}
}
