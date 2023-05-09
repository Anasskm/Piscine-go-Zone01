package piscine

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	originalhead := l.Head
	for i := l.Head; i != nil; i = i.Next {
		if comp(i.Data, ref) {
			return &i.Data
		}
	}
	l.Head = originalhead
	return nil
}
