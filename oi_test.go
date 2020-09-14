package main

import "testing"

func TestOi(t *testing.T) {
	resultado := Oi("teste")
	esperado := "Deu bom, teste?"

	if resultado != esperado {
		t.Errorf("resultado: '%s', esperado: '%s'", resultado, esperado)
	}
}
