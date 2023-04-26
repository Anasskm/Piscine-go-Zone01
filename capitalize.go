package piscine

func Capitalize(s string) string {
	r := []rune(s)
	if checkiflower(r[0]) {
		r[0] -= 32
	}
	for i := 1; i < len(r); i++ {
		if checkifupper(r[i]) {
			if CheckIfLetterOrNumber(r[i-1]) {
				r[i] += 32
			}
			continue

		}
		if checkiflower(r[i]) {
			if CheckIfLetterOrNumber(r[i-1]) {
				continue
			}
			r[i] -= 32

		}
	}

	st := string(r)
	return st
}

func CheckIfLetterOrNumber(r rune) bool {
	if (r >= 48 && r <= 57) || (r >= 65 && r <= 90) || (r >= 97 && r <= 122) {
		return true
	}
	return false
}

func checkiflower(r rune) bool {
	if r < 97 || r > 122 {
		return false
	}

	return true
}

func checkifupper(r rune) bool {
	if r < 65 || r > 90 {
		return false
	}

	return true
}
