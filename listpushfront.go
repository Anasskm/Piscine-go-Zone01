package piscine

func ListPushFront(l *List, data interface{}) {
	n := &Nodel{Data: data}
	if l.Tail == nil {
		l.Tail = n
		l.Head = n
		return
	}
	l.Head.Next = n
	l.Head = l.Head.Next
}
