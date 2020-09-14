package main

import "testing"

func TestOi(t *testing.T) {
	resultado := Oi()
	esperado := "Deu bom?"

	if resultado != esperado {
		t.Errorf("resultado: '%s', esperado: '%s'", resultado, esperado)
	}
}
