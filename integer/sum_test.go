package integer

import "testing"
import "fmt"

func TestSum(t *testing.T) {
  sumResult := Sum(2,2)
  expect := 4

  if sumResult != expect {
    t.Errorf("esperado: '%d', recebido: '%d'", expect, sumResult)
  }
}

func ExampleSum() {
  sumResult := Sum(3,5)
  fmt.Println(sumResult)
  // Output: 8
}
