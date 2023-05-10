package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeInsertData(root *TreeNode, data string) *TreeNode {
	newNode := &TreeNode{Data: data}

	if root == nil {
		root = newNode
		return root
	}

	i := root
	for {
		if data < i.Data {
			if i.Left == nil {
				i.Left = newNode
				newNode.Parent = i
				break
			}
			i = i.Left
		} else {
			if i.Right == nil {
				i.Right = newNode
				newNode.Parent = i
				break
			}
			i = i.Right
		}
	}

	return root
}
