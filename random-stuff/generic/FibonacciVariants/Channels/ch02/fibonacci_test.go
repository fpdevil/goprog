package ch02

import "testing"

var fibTests = []struct {
	a        int
	expected int
}{
	{0, 0},
	{1, 1},
	{1, 2},
	{3, 3},
	{4, 5},
	{10, 55},
	{20, 10946},
	{50, 12586269025},
}

func TestChan(t *testing.T) {
	for _, ft := range fibTests {
		if v := FibChan(ft.a); v != ft.expected {
			t.Errorf("FibChan(%d) returned %d, expected %d", ft.a, v, ft.expected)
		}
	}
}

func BenchmarkFibChan(b *testing.B) {
	fc := FibChan
	for i := 0; i < b.N; i++ {
		_ = fc(8)
	}
}

func benchmarkFibChan(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibChan(i)
	}
}

func BenchmarkFibChan0(b *testing.B) { benchmarkFibChan(0, b) }
func BenchmarkFibChan1(b *testing.B) { benchmarkFibChan(1, b) }
func BenchmarkFibChan2(b *testing.B) { benchmarkFibChan(2, b) }
func BenchmarkFibChan3(b *testing.B) { benchmarkFibChan(3, b) }
func BenchmarkFibChan4(b *testing.B) { benchmarkFibChan(4, b) }
