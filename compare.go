package compare

// TCompatible represents a type compatible with Helper and Fatalf from
// testing.T. This allows T itself to be mocked
type TCompatible interface {
	Helper()
	Fatalf(string, ...any)
}

// Compare compares two values of the same type for equality and fails the
// current test if they are not equal.
func Compare[T comparable](t TCompatible, v1, v2 T, matches bool) {
	t.Helper()
	if matches {
		if v1 != v2 {
			t.Fatalf("%v != %v", v1, v2)
		}
	} else {
		if v1 == v2 {
			t.Fatalf("%v == %v", v1, v2)
		}
	}
}
