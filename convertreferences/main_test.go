package convertreferences

import "testing"

func Benchmark_NewSlice0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		o := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
		n := make([]string, 0)
		for _, v := range o {
			n = append(n, v)
		}
	}
}

func Benchmark_NewSliceLen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		o := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
		n := make([]string, len(o))
		for i, v := range o {
			n[i] = v
		}
	}
}

func Benchmark_NewSliceCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		o := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
		n := make([]string, 0, len(o))
		for _, v := range o {
			n = append(n, v)
		}
	}
}

func Benchmark_NewMap0(b *testing.B) {
	for i := 0; i < b.N; i++ {
		o := map[string]string{"a": "a", "b": "b", "c": "c", "d": "d", "e": "e", "f": "f", "g": "g", "h": "h", "i": "i", "j": "j", "k": "k", "l": "l", "m": "m", "n": "n", "o": "o", "p": "p", "q": "q", "r": "r", "s": "s", "t": "t", "u": "u", "v": "v", "w": "w", "x": "x", "y": "y", "z": "z"}
		n := make(map[string]string)
		for k, v := range o {
			n[k] = v
		}
	}
}

func Benchmark_NewMapSize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		o := map[string]string{"a": "a", "b": "b", "c": "c", "d": "d", "e": "e", "f": "f", "g": "g", "h": "h", "i": "i", "j": "j", "k": "k", "l": "l", "m": "m", "n": "n", "o": "o", "p": "p", "q": "q", "r": "r", "s": "s", "t": "t", "u": "u", "v": "v", "w": "w", "x": "x", "y": "y", "z": "z"}
		n := make(map[string]string, len(o))
		for k, v := range o {
			n[k] = v
		}
	}
}
