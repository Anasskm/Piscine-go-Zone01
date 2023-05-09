package piscine

func ListReverse(l *List) {
	ElmtCourant := l.Head
	var ElmtPrecedent *NodeL
	l.Tail = l.Head

	for ElmtCourant != nil {
		next := ElmtCourant.Next
		ElmtCourant.Next = ElmtPrecedent
		ElmtPrecedent = ElmtCourant
		ElmtCourant = next
	}

	l.Head = ElmtPrecedent
}
