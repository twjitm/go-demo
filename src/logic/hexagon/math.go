package hexagon

func Abs(a int32) int32 {
	if a > 0 {
		return a
	}
	return -a
}

func Max(a, b int32) int32 {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int32) int32 {
	if a < b {
		return a
	}
	return b
}
