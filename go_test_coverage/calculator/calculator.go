package calculator

type Calculator struct {
	name string
}

func sum(a, b uint8) uint16 {
	return uint16(a) + uint16(b)
}

func subtract(a, b int8) int16 {
	return int16(a) - int16(b)
}

func multiply(a, b int8) int16 {
	return int16(a) * int16(b)
}

func abs(a int8) int8 {
	if a < 0 {
		return -a
	}

	return a
}
