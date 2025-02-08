package main

import (
	"fmt"
	"os"
	"sort"
	"time"
)

// Function to generate prime numbers up to a limit
func getPrimes(n int) []int {
	primes := []int{}
	for num := 2; num < n; num++ {
		isPrime := true
		for i := 2; i < num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, num)
		}
	}
	return primes
}

// Bubble sort implementation (inefficient)
func bubbleSort(arr []int) []int {
	n := len(arr)
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

// Function to compute sum of squares
func sumOfSquares(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num * num
	}
	return sum
}

func main() {
	start := time.Now()

	// Generate prime numbers up to 10,000
	primes := getPrimes(123456)
	fmt.Printf("Found %d prime numbers.\n", len(primes))

	// Sort the primes
	var sortedPrimes []int
	if os.Getenv("USE_BUBBLE") != "" {
		sortedPrimes = bubbleSort(primes)
	} else {
		sort.Ints(primes)
		sortedPrimes = primes
	}
	fmt.Printf("Sorted %d prime numbers.\n", len(sortedPrimes))

	// Compute sum of squares
	total := sumOfSquares(sortedPrimes)
	fmt.Printf("Sum of squares: %d\n", total)

	fmt.Printf("Execution time: %v\n", time.Since(start))
	fmt.Printf("USE_BUBBLE was : %v\n", os.Getenv("USE_BUBBLE"))
}
