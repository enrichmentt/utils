package utils

func InSlice(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}

func ContainsInt(a []int, x int) bool {
	for _, v := range a {
		if v == x {
			return true
		}
	}

	return false
}
