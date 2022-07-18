package fib

import "testing"

func TestFib(t *testing.T) {
	for _, tt := range fibTests {
		actual := Fib(tt.n)
		if actual != tt.expected {
			t.Errorf("Fibo(%d): expected %d, actual %d", tt.n, tt.expected, actual)
		}
	}
}
