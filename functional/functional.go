package functional


import (
)


/*
  Map
*/

func MapString(arr []string, f func(string) string) []string {
  result := make([]string, len(arr))
  for i, v := range arr {
    result[i] = f(v)
  }

  return result
}

func MapFloat(arr []float64, f func(float64) float64) []float64 {
  result := make([]float64, len(arr))
  for i, v := range arr {
    result[i] = f(v)
  }

  return result
}

func MapInt(arr []int, f func(int) int) []int {
  result := make([]int, len(arr))
  for i, v := range arr {
    result[i] = f(v)
  }

  return result
}


/*
  Index
*/

// Returns the first index of the target string t, or -1 if no match is found.
func IndexString(vs []string, t string) int {
    for i, v := range vs {
        if v == t {
            return i
        }
    }
    return -1
}

func IndexFloat(vs []float64, t float64) int {
    for i, v := range vs {
        if v == t {
            return i
        }
    }
    return -1
}

func IndexInt(vs []int, t int) int {
    for i, v := range vs {
        if v == t {
            return i
        }
    }
    return -1
}

/*
  Include
*/

// Returns true if the target string t is in the slice.
func IncludeString(vs []string, t string) bool {
    return IndexString(vs, t) >= 0
}

func IncludeFloat(vs []float64, t float64) bool {
    return IndexFloat(vs, t) >= 0
}

func IncludeInt(vs []int, t int) bool {
    return IndexInt(vs, t) >= 0
}


/*
  Any
*/

// Returns true if one of the strings in the slice satisfies the predicate f.
func AnyString(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if f(v) {
            return true
        }
    }
    return false
}

func AnyFloat(vs []float64, f func(float64) bool) bool {
    for _, v := range vs {
        if f(v) {
            return true
        }
    }
    return false
}

func AnyInt(vs []int, f func(int) bool) bool {
    for _, v := range vs {
        if f(v) {
            return true
        }
    }
    return false
}

/*
  All
*/

// Returns true if all of the strings in the slice satisfy the predicate f.
func AllString(vs []string, f func(string) bool) bool {
    for _, v := range vs {
        if !f(v) {
            return false
        }
    }
    return true
}

func AllFloat(vs []float64, f func(float64) bool) bool {
    for _, v := range vs {
        if !f(v) {
            return false
        }
    }
    return true
}


func AllInt(vs []int, f func(int) bool) bool {
    for _, v := range vs {
        if !f(v) {
            return false
        }
    }
    return true
}


/*
  Filter
*/

// Returns a new slice containing all strings in the slice that satisfy the predicate f.
func FilterString(vs []string, f func(string) bool) []string {
    vsf := make([]string, 0)
    for _, v := range vs {
        if f(v) {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

func FilterFloat(vs []float64, f func(float64) bool) []float64 {
    vsf := make([]float64, 0)
    for _, v := range vs {
        if f(v) {
            vsf = append(vsf, v)
        }
    }
    return vsf
}

func FilterInt(vs []int, f func(int) bool) []int {
    vsf := make([]int, 0)
    for _, v := range vs {
        if f(v) {
            vsf = append(vsf, v)
        }
    }
    return vsf
}
