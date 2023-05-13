package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil {
		return nil
	}
	s := []*TreeNode{root}
	for len(s) > 0 {
		n := s[len(s)-1]
		s = s[:len(s)-1]
		if n.Data == elem {
			return n
		}
		if elem < n.Data {
			s = append(s, n.Left)
		} else {
			s = append(s, n.Right)
		}

	}
	return nil
}
