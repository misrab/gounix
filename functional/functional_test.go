package functional


import (
  // "fmt"
  "testing"
)


func TestMapString(t *testing.T) {
  input := []string{"abc", "def", "ghi"}
  f := func(s string) string {
    return s[:len(s)-1] // pop the last character
  }
  desired := []string{"ab", "de", "gh"}
  result := MapString(input, f)
  if EqualString(desired, result) == false {
    t.Errorf("Result does not match desired.\n")
  }
  // fmt.Printf("%v\n", MapString(input, f))
}
