package unittest

func ToDays(amount int, unit string) int {
	// 1-600
	if amount < 1 || amount > 600 {
		return -1
	}

	if unit == "" {
		return -1
	}

	// d -> day
	if unit == "d" {
		return amount
	}

	// m -> bulan
	if unit == "m" {
		return amount * 30
	}

	// w -> minggu
	if unit == "w" {
		return amount * 7
	}

	return -1
}