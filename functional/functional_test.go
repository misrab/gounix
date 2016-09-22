package functional


import (
  "testing"
)


func TestMapString(t *testing.T) {
  input := []string{"aaa", "bbb", "ccc"}
  f := func(string) string {
    return "ddd"
  }
  MapString(input, f)
}
