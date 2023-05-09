package piscine

func IsPositiveNode(node *NodeL) bool {
	switch node.Data.(type) {
	case int, float32, float64, byte:
		return node.Data.(int) > 0
	default:
		return false
	}
}

func IsAlNode(node *NodeL) bool {
	switch node.Data.(type) {
	case string:
		return true
	default:
		return false
	}
}

func ListForEachIf(l *List, f func(*NodeL), cond func(*NodeL) bool) {
	for i := l.Head; i != nil; i = i.Next {
		if cond(i) {
			f(i)
		}
	}
}
