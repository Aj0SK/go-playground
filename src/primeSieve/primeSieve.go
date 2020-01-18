package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func isPrime(x int) bool {
	if x == 2 {
		return true
	}

	if (x < 2) || ((x % 2) == 0) {
		return false
	}

	for i := 3; i*i <= x; i += 2 {
		if (x % i) == 0 {
			return false
		}
	}
	return true
}

func sieveOfErat(n int) []int {
	var sqrtN int = int(math.Sqrt(float64(n))) + 1
	sieve := make([]bool, n)
	primes := make([]int, 1)

	primes[0] = 2

	for i := 3; i < n; i += 2 {
		if sieve[i] == false {
			primes = append(primes, i)

			if i > sqrtN {
				continue
			}
			for j := i * i; j < n; j += 2 * i {
				sieve[j] = true
			}
		}
	}
	return primes
}

func segmentErat(n int) []int {
	var sqrtN int = int(math.Sqrt(float64(n))) + 1
	var segmentSize int = sqrtN

	var primes = sieveOfErat(segmentSize)
	segmentSieve := make([]bool, segmentSize)

	for b := segmentSize; b < n; b += segmentSize {
		for _, p := range primes {
			if p > sqrtN {
				break
			}

			var helper int = b - (b % p)
			if helper != b {
				helper = helper + p
			}

			for j := helper - b; j < segmentSize; j += p {
				segmentSieve[j] = true
			}
		}

		for i := 0; i < segmentSize; i++ {
			if segmentSieve[i] == false && b+i < n {
				primes = append(primes, b+i)
			}
			segmentSieve[i] = false
		}
	}

	return primes
}

func measureTime(maxPrime int) {
	fmt.Printf("n = %d\n", maxPrime)

	start := time.Now()
	var erat1 = sieveOfErat(maxPrime)
	elapsed := time.Since(start)
	fmt.Printf("Classic sieve %s\n", elapsed)

	start = time.Now()
	var erat2 = segmentErat(maxPrime)
	elapsed = time.Since(start)
	fmt.Printf("Segmented sieve %s\n", elapsed)

	var sum int = 0
	for _, x := range erat1 {
		sum += x
	}
	for _, x := range erat2 {
		sum += x
	}
}

func main() {
	const testCases int = 10
	var maxPrime int = 100000

	primes := make([]int, 0)

	for i := 1; i < maxPrime; i++ {
		if isPrime(i) == true {
			primes = append(primes, i)
		}
	}

	var erat = sieveOfErat(maxPrime)
	var segmerat = segmentErat(maxPrime)

	if len(erat) != len(primes) {
		panic("Bad sieveOfErat!")
	}
	if len(segmerat) != len(primes) {
		panic("Bad segmentErat!")
	}

	for i := 0; i < len(primes); i++ {
		if erat[i] != primes[i] {
			panic("Bad sieveOfErat!")
		}
		if segmerat[i] != primes[i] {
			panic("Bad segmentErat!")
		}
	}

	maxPrime = 2000000000
	for i := 0; i < testCases; i++ {
		var upTo int = 10 + (rand.Int() % (maxPrime - 10))
		measureTime(upTo)
	}

}
