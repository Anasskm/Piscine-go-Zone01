package piscine

func StrRev(s string) string {
	var v string
	for _, t := range s {
		v = string(t) + v
	}
	return v
}
