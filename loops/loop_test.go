package loops

import "testing"

func TestLoop(t *testing.T){
  result := Loop("a", 5)
  expect := "aaaaa"

  if expect != result {
    t.Errorf("esperado: '%s', recebido: '%s'", expect, result)
  }
}

func BenchmarkLopp(b *testing.B) {
  for i := 0; i < b.N; i++ {
    Loop("a", 1)
  }
}
