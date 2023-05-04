package piscine

func CollatzCountdown(start int) int {
	count := 0
	if start <= 0 {
		return -1
	}
	for start != 0 {
		if start%2 != 0 {
			start = (start * 3) + 1
			count++

		} else {
			start /= 2
			count++
		}
		if start == 1 {
			break
		}
	}

	return count
}
