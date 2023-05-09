package piscine

func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}
	for l.Head != nil && data_ref == l.Head.Data {
		l.Head = l.Head.Next
	}
	i := l.Head
	for i != nil && i.Next != nil {
		if i.Next.Data == data_ref {
			i.Next = i.Next.Next
		} else {
			i = i.Next
		}
	}
}
