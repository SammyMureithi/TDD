package main

func Sum(numbers []int) int {
	total := 0
	for i := 0; i < len(numbers); i++ {
		total += numbers[i]
	}
	return total
}

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}

	return sums
}