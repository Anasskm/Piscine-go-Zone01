package piscine

func Capitalize(s string) string {
	r := []rune(s)
	if checkiflower(r[0]) {
		r[0] -= 32
	}
	for i := 1; i < len(r); i++ {
		if checkifupper(r[i]) {
			if CheckIfNotLetterOrNumber(r[i-1]) {
				continue
			} else {
				r[i] += 32
			}
		}
		if checkiflower(r[i]) {
			if CheckIfNotLetterOrNumber(r[i-1]) {
				r[i] -= 32
			} else {
				continue
			}
		}
	}

	st := string(r)
	return st
}

func CheckIfNotLetterOrNumber(r rune) bool {
	if (r < 48 || r > 122) && (r < 48 || r > 57) && (r < 65 || r > 57) && (r < 97 || r > 90) {
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
