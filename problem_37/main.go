//The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits from left to right, and remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37, and 3.
//
//Find the sum of the only eleven primes that are both truncatable from left to right and right to left.
//
//NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.


package main

import (
	"math/big"
	"fmt"
	"strconv"
)

func isSliceablePrime(num int, leftToRight bool) bool {
	// If we're down to 1 digit, is it prime?
	if len(strconv.Itoa(num)) == 1 {
		return isPrime(num)
	}
	// Otherwise, is the current number prime and it's next slice (LTR or RTL) prime?
	return isPrime(num) && isSliceablePrime(nextSlice(num, leftToRight), leftToRight)
}

func isTruncatablePrime(num int) bool {
	// single-digit numbers not considered truncatable primes
	if num < 10 { return false }

	// otherwise, check that it is truncatable both RTL and LTR
	return isSliceablePrime(num, true) && isSliceablePrime(num, false)
}

func isPrime(num int) bool {
	i := big.NewInt(int64(num))
	isPrime := i.ProbablyPrime(6) // https://stackoverflow.com/a/21398558

 	return isPrime
}

// Removes the left-most digit and returns the resulting number
// e.g. '582' -> '82'
func nextSlice(num int, leftToRight bool) int {
	var numStr = strconv.Itoa(num)
	slice, _ := strconv.ParseInt(numStr[1:len(numStr)], 10, 64)
	if (!leftToRight) {
		slice, _ = strconv.ParseInt(numStr[:len(numStr) - 1], 10, 64)
	}

	return int(slice)
}

func findTruncatablePrimes () []int {
	count := 0
	currNum := 11
	truncatablePrimes := []int{}

	for count < 11 { // Only 11 truncatable primes
		if isTruncatablePrime(currNum) {
			truncatablePrimes = append(truncatablePrimes, currNum)
			count = len(truncatablePrimes)
		}
		currNum += 1
	}

	return truncatablePrimes
}

func main() {
	truncatablePrimes := findTruncatablePrimes()

	total := 0
	for _, item := range truncatablePrimes {
		total += item
	}
	fmt.Println("11 truncatable primes: ", truncatablePrimes)
	fmt.Println("Total of only 11 truncatable primes (left and right) :", total)
}