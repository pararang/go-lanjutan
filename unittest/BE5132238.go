package unittest

func HomeworkBE5132238(number int) int {
	if number > 8 {
		return -1
	}
	if number < 0 {
		return 0
	}
	if number == 0 {
		return 1
	}
	return number * HomeworkBE5132238(number-1)
}
