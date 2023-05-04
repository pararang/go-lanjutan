package unittest

type Wide struct {
	Unit  string
	Value float64
}

// bikin method untuk konversi ke meter. returnya float64
func (w Wide) ToMeter() float64 {
	var result float64

	if w.Unit == "" {
		return -1
	}

	if w.Unit == "m" {
		result = w.Value
	}

	if w.Unit == "cm" {
		result = w.Value / 100
	}

	return result
}
