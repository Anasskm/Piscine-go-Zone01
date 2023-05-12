package piscine

func BTreeApplyPreorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	s := []*TreeNode{root}
	for len(s) > 0 {
		n := s[len(s)-1]
		s = s[:len(s)-1]
		if n != nil {
			f(n.Data)

			s = append(s, n.Right, n.Left)

		}

	}
}
