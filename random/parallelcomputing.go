package main

import (
	"runtime"
	"sync"
)

type Vector []float64

// // Convolve computes w = u *v , where w[k] Σ u[i]*v[j], i + j = k.
// // Precondition: len(u) > 0, len(v) > 0.
// func Convolve(u, v Vector) Vector {
// 	n := len(u) + len(v) - 1
// 	w := make(Vector, n)
// 	for k := 0; k < n; k++ {
// 		w[k] = mul(u, v, k)
// 	}
// 	return w
// }

// // mul return s Σ u[i]*v[j], i + j = k.
// func mul(u, v Vector, k int) float64 {
// 	var res float64
// 	n := min(k+1, len(u))
// 	j := min(k, len(v)-1)
// 	for i := k - j; i < n; j = i + 1, j - 1 {
// 		res += u[i] * v[j]
// 	}
// 	return res
// }

func init() {
	numcpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numcpu) // try to use all availabe CPUS.
}

func Convolve(u, v Vector) Vector {
	n := len(u) + len(v) - 1
	w := make(Vector, n)
}

func main() {
	// Divide w into work units that take ~100μs-1ms to compute.
	size := max(1, 1000000/n)

	var wg sync.WaitGroup
	for i, j := 0, size; i < n; i, j = j, j+size {
		if j > n {
			j = n
		}

		// These goroutines share memory, but oly for reading,
		wg.Add(1)
		go func(i, j int) {
			for k := i; k < j; k++ {
				w[k] = mul(u, v, k)
			}
			wg.Done()
		}(i, j)
	}
	wg.Wait()
	return w
}

// No idea how this work.
