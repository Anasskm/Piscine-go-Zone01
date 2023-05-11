package piscine

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}

	stack := []*TreeNode{root}
	var lastVisited *TreeNode

	for len(stack) > 0 {
		current := stack[len(stack)-1]

		if lastVisited == nil || lastVisited.Left == current || lastVisited.Right == current {
			if current.Left != nil {
				stack = append(stack, current.Left)
			} else if current.Right != nil {
				stack = append(stack, current.Right)
			}
		} else if current.Left == lastVisited {
			if current.Right != nil {
				stack = append(stack, current.Right)
			}
		} else {
			f(current.Data)
			stack = stack[:len(stack)-1]
		}

		lastVisited = current
	}
}
