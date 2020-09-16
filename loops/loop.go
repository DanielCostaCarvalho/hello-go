package loops

func Loop(received string, times int) string {
  result := ""
  for i := 0; i < times; i++ {
    result += received
  }
  return result
}
