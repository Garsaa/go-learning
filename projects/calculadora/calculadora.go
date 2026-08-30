// Package calculadora fornece operações matemáticas simples.
package calculadora

// Somar retorna a soma de a e b.
func Somar(a, b int) int {
	return a + b
}

// Dobrar retorna o dobro de n.
func Dobrar(n int) int {
	return Somar(n, n)
}
