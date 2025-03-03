package patterns

import "testing"

func TestMap(t *testing.T) {
	for name, mm := range map[string]SafeMap[string, int]{
		"MapMu":   NewMapMu[string, int](),
		"MapSync": NewMapSync[string, int](),
	} {
		t.Run(name, func(t *testing.T) {
			mm.LoadOrStore("a", 1)
			mm.LoadOrStore("b", 2)
			mm.LoadOrStore("c", 3)
			mm.LoadOrStore("d", 4)
			if mm.Len() != 4 {
				t.Errorf("expected 4, got %d", mm.Len())
			}

			mm.Clear()

			if mm.Len() != 0 {
				t.Errorf("expected 0, got %d", mm.Len())
			}

			mm.LoadOrStore("a", 1)
			mm.LoadOrStore("b", 2)
			mm.LoadOrStore("c", 3)
			v, ok := mm.Load("b")
			if !ok {
				t.Errorf("expected true, got false")
			}

			if v != 2 {
				t.Errorf("expected 2, got %d", v)
			}

			if mm.CompareAndDelete("b", 1) {
				t.Errorf("expected false, got true")
			}

			if !mm.CompareAndDelete("b", v) {
				t.Errorf("expected false, got true")
			}

			if !mm.CompareAndSwap("c", 3, 5) {
				t.Errorf("expected true, got false")
			}

			v, ok = mm.Load("c")
			if !ok {
				t.Errorf("expected true, got false")
			}

			if v != 5 {
				t.Errorf("expected 5, got %d", v)
			}

			m := mm.Map()

			if len(m) != mm.Len() {
				t.Errorf("expected %d, got %d", mm.Len(), len(m))
			}

			for k, v := range m {
				if c, ok := mm.Load(k); !ok || v != c {
					t.Errorf("expected %d, got %d", v, c)
				}
			}
		})
	}
}
