package main

import "testing"

func TestOi(t *testing.T) {
	verificaMensagem := func(t *testing.T, resultado, esperado string) {
		t.Helper()
		if resultado != esperado {
			t.Errorf("resultado: '%s', esperado: '%s'", resultado, esperado)
		}
	}

	t.Run("Pergunta se deu bom para o nome recebido", func(t *testing.T) {
		resultado := Hello("teste")
		esperado := "Deu bom, teste?"
		
		verificaMensagem(t, resultado, esperado)
	})

	t.Run("Pergunta só se deu bom, se for passada uma string vazia", func(t *testing.T) {
		resultado := Hello("")
		esperado := "Deu bom?"

		verificaMensagem(t, resultado, esperado)
	})
}
