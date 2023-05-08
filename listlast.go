package piscine

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	var n *NodeL
	for l.Head != nil {
		n = l.Head
		l.Head = l.Head.Next

	}
	return n.Data
}
