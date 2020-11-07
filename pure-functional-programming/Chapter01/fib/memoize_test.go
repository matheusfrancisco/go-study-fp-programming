package fib

import "testing"

var fibTests = []struct {
	a        int
	expected int
}{
	{1, 1},
	{2, 2},
	{3, 3},
	{4, 5},
	{20, 10946},
	{42, 433494437},
}

func TestSimple(t *testing.T) {
	for _, ft := range fibTests {
		if v := FibSimple(ft.a); v != ft.expected {
			t.Errorf("FibSimple(%d) returned %d, expected %d", ft.a, v, ft.expected)
		}
	}
}

func BenchmarkFibSimple(b *testing.B) {
	fn := FibSimple
	for i := 0; i < b.N; i++ {
		_ = fn(8)
	}
}

func benchmarkFibSimple(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibSimple(i)
	}
}

func BenchmarkFibSimple1(b *testing.B)  { benchmarkFibSimple(1, b) }
func BenchmarkFibSimple2(b *testing.B)  { benchmarkFibSimple(2, b) }
func BenchmarkFibSimple3(b *testing.B)  { benchmarkFibSimple(3, b) }
func BenchmarkFibSimple10(b *testing.B) { benchmarkFibSimple(4, b) }
func BenchmarkFibSimple20(b *testing.B) { benchmarkFibSimple(20, b) }
func BenchmarkFibSimple40(b *testing.B) { benchmarkFibSimple(42, b) }

func TestMemoized(t *testing.T) {
	for _, ft := range fibTests {
		if v := FibMemoized(ft.a); v != ft.expected {
			t.Errorf("FibMemoized(%d) returned %d, expected %d", ft.a, v, ft.expected)
		}
	}
}

func BenchmarkFibMemoized(b *testing.B) {
	fn := FibMemoized
	for i := 0; i < b.N; i++ {
		_ = fn(8)
	}
}

func benchmarkFibMemoized(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibMemoized(i)
	}
}

func BenchmarkFibMemoized1(b *testing.B)  { benchmarkFibMemoized(1, b) }
func BenchmarkFibMemoized2(b *testing.B)  { benchmarkFibMemoized(2, b) }
func BenchmarkFibMemoized3(b *testing.B)  { benchmarkFibMemoized(3, b) }
func BenchmarkFibMemoized10(b *testing.B) { benchmarkFibMemoized(4, b) }
func BenchmarkFibMemoized20(b *testing.B) { benchmarkFibMemoized(20, b) }
func BenchmarkFibMemoized40(b *testing.B) { benchmarkFibMemoized(42, b) }
