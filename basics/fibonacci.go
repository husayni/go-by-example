// basics/fibonacci.go
package main

import "fmt"

// Function to generate Fibonacci sequence up to n terms
func generateFibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}
	fibSeq := make([]int, n)
	fibSeq[0] = 0
	if n > 1 {
		fibSeq[1] = 1
		for i := 2; i < n; i++ {
			fibSeq[i] = fibSeq[i-1] + fibSeq[i-2]
		}
	}
	return fibSeq
}

func main() {
	n := 10 // Number of terms in the Fibonacci sequence
	fibSeq := generateFibonacci(n)
	fmt.Println("Fibonacci sequence:", fibSeq)
}