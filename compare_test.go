package compare

import (
	"testing"
)

type mockT struct {
	fatal bool
}

func (m *mockT) Helper()               {}
func (m *mockT) Fatalf(string, ...any) { m.fatal = true }

func TestCompare(t *testing.T) {
	for _, v := range []struct {
		name    string
		v2      bool
		matches bool
		fatal   bool
	}{
		{
			name:    "ShouldNotMatchAndDoesnt",
			v2:      false,
			matches: false,
			fatal:   false,
		},
		{
			name:    "ShouldNotMatchAndDoes",
			v2:      true,
			matches: false,
			fatal:   true,
		},
		{
			name:    "ShouldMatchAndDoesnt",
			v2:      false,
			matches: true,
			fatal:   true,
		},
		{
			name:    "ShouldMatchAndDoes",
			v2:      true,
			matches: true,
			fatal:   false,
		},
	} {
		t.Run(v.name, func(t *testing.T) {
			m := &mockT{}
			Compare(m, true, v.v2, v.matches)
			if m.fatal != v.fatal {
				t.Fatalf("%v != %v", m.fatal, v.fatal)
			}
		})
	}
}
