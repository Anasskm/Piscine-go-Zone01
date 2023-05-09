package piscine

func ListMerge(l1 *List, l2 *List) {
	if l1.Head == nil {
		l1.Head = l2.Head
	} else {
		l1.Tail.Next = l2.Head
	}
}
