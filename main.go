package main

import (
	"fmt"
	"math"
	"sync"
)

/*
type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
*/

func is_prime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func count_prime_parallel(numbers []int, how_many int, numThreads int) int {
	rowsPerThread := int(math.Ceil(float64(how_many) / float64(numThreads)))
	count := 0
	var wg sync.WaitGroup
	var mu sync.Mutex

	for t := 0; t < numThreads; t++ {
		wg.Add(1)
		go func(t int) {
			defer wg.Done()
			startRow := t * rowsPerThread
			endRow := min((t+1)*rowsPerThread, how_many)
			localCount := 0
			for i := startRow; i < endRow; i++ {
				if is_prime(numbers[i]) {
					localCount++
				}
			}
			mu.Lock()
			count += localCount
			mu.Unlock()
		}(t)
	}
	wg.Wait()
	return count
}

func main() {
	numbers := make([]int, 100)
	for i := range numbers {
		numbers[i] = i
	}
	fmt.Println("Prime count:", count_prime_parallel(numbers, len(numbers), 4))
}
