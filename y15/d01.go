package y15

func CalcFloor(pattern []byte) int {
	f := 0

	for _, b := range pattern {
		if b == 40 {
			f++
		} else {
			f--
		}
	}

	return f
}

func CalcFloor2(pattern []byte) int {
	f := 0

	for i, b := range pattern {
		if b == 40 {
			f++
		} else {
			f--
		}

		if f < 0 {
			return i + 1
		}
	}

	return -1
}
