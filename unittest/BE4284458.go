package unittest

func HomeworkBE4284458(numbers []int) int {
	if len(numbers) > 25 {
		return -1
	}

	largest := numbers[0]
	for _, number := range numbers {
		if number == 0 {
			return -1
		}

		if number > largest {
			largest = number
		}
	}
	return largest
}
