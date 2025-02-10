package main

import (
	"fmt"
	"os"
	"sync"
  "math"
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

// worker function to check if a number is prime and send it to the channel
func checkPrime(num int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Decrement WaitGroup counter when goroutine completes

	isPrime := true
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			isPrime = false
			break
		}
	}
	if isPrime {
		ch <- num // Send prime number to the channel
	}
}

// getPrimesConcurrent generates prime numbers up to 'n' using goroutines and channels
func getPrimesConcurrent(n int) []int {
	if n < 2 {
		return []int{}
	}

	ch := make(chan int, n) // Buffered channel to collect primes
	var wg sync.WaitGroup   // WaitGroup to track goroutines

	// Launch worker goroutines
	for num := 2; num < n; num++ {
		wg.Add(1) // Increment WaitGroup counter
		go checkPrime(num, ch, &wg)
	}

	// Close the channel once all goroutines finish
	go func() {
		wg.Wait()  // Wait for all goroutines to complete
		close(ch)  // Close the channel safely
	}()

	// Collect results from channel
	primes := []int{}
	for prime := range ch {
		primes = append(primes, prime)
	}

	return primes
}


func main() {
	start := time.Now()

	// Generate prime numbers up to 10,000
	var primes []int
  if os.Getenv("PRIME_ALG") == "none" {
    primes = getPrimes(123456)
  } else {
    primes = getPrimesConcurrent(123456)
  }

	fmt.Printf("Found %d prime numbers.\n", len(primes))

	// Sort the primes
	var sortedPrimes []int
	if os.Getenv("SORT_ALG") != "bubble" {
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
	fmt.Printf("SORT_ALG was : %v\n", os.Getenv("SORT_ALG"))
	fmt.Printf("PRIME_ALG was : %v\n", os.Getenv("PRIME_ALG"))
}
