package main

// `Array` has a fixed size
// var array = [2]int
// array := [2]int{1,2}

// `Slice` hasn't a fixed size
// var slice []int
// slice := []int{1,2,3}

// Sum returns the sum of all the numbers in `Slice`
func Sum(numbers []int) int {
	sum := 0
	// range returns index and and value
	// ignoring the index using `_` as a placeholder (blank identifier)
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAll returns the sum of every `Slice` and puts the results in a `Slice`
// This function can be called with many arguments (of the same type) as we want
func SumAll(numbersToSum ...[]int) []int {

	/*lenghtOfNumbers := len(numbersToSum)
	// `make` initialize an Array with a given length
	sums := make([]int, lenghtOfNumbers)
	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}*/

	var sums []int
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

// SumAllTails returns the sum of all the tails of provided `Slice`s in a `Slice`
func SumAllTails(numnersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numnersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			// selecting all the elements but the first
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
