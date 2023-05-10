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

	itterator := root
	for {
		if data < itterator.Data {
			if itterator.Left == nil {
				itterator.Left = newNode
				newNode.Parent = itterator
				break
			}
			itterator = itterator.Left
		} else {
			if itterator.Right == nil {
				itterator.Right = newNode
				newNode.Parent = itterator
				break
			}
			itterator = itterator.Right
		}
	}

	return root
}
