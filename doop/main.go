package main

import "os"

func atoi(s string) (int, bool) {
	var res int
	var neg bool

	if s == "" {
		return 0, false
	}

	if s[0] == '-' {
		neg = true
		s = s[1:]
	}

	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		res = res*10 + int(c-'0')
		if res < 0 {
			return 0, false
		}
	}

	if neg {
		return -res, true
	}

	return res, true
}

func itoa(i int) string {
	if i == 0 {
		return "0"
	}

	neg := false
	if i < 0 {
		neg = true
		i = -i
	}

	buf := make([]byte, 0, 20)

	for i > 0 {
		digit := i % 10
		buf = append([]byte{'0' + byte(digit)}, buf...)
		i /= 10
	}

	if neg {
		buf = append([]byte{'-'}, buf...)
	}

	return string(buf)
}

func main() {
	if len(os.Args) != 4 {
		return
	}

	val1, ok1 := atoi(os.Args[1])
	val2, ok2 := atoi(os.Args[3])

	if !ok1 || !ok2 {
		return
	}

	operator := os.Args[2]

	if operator != "+" && operator != "-" && operator != "*" && operator != "/" && operator != "%" {
		return
	}

	if operator == "/" && val2 == 0 {
		os.Stdout.WriteString("No division by 0\n")
		return
	}

	if operator == "%" && val2 == 0 {
		os.Stdout.WriteString("No modulo by 0\n")
		return
	}

	switch operator {
	case "+":
		result := val1 + val2
		if (val1 > 0 && val2 > 0 && result < 0) || (val1 < 0 && val2 < 0 && result > 0) {
			return
		}
		os.Stdout.WriteString(itoa(result) + "\n")
	case "-":
		result := val1 - val2
		if (val1 > 0 && val2 < 0 && result < 0) || (val1 < 0 && val2 > 0 && result > 0) {
			return
		}
		os.Stdout.WriteString(itoa(result) + "\n")
	case "*":
		result := val1 * val2
		if (val2 != 0 && result/val2 != val1) || (val1 == -1 && val2 == -9223372036854775808) {
			return
		}
		os.Stdout.WriteString(itoa(result) + "\n")
	case "/":
		result := val1 / val2
		os.Stdout.WriteString(itoa(result) + "\n")
	case "%":
		result := val1 % val2
		os.Stdout.WriteString(itoa(result) + "\n")
	}
}
