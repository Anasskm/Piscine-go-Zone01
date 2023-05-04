package piscine

func CollatzCountdown(start int) int {
	count := 1
	if start <= 0 {
		return -1
	}
	for start != 0 {
		if start%2 != 0 {
			start = ((start * 3) + 1) / 2
			count++

		} else {
			start /= 2
			count++
		}
		if start == 1 {
			count++
			break
		}
	}

	return count
}
