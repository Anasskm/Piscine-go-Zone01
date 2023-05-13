package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil || root.Data == elem {
		return root
	}
	if result := BTreeSearchItem(root.Left, elem); result != nil {
		return result
	}
	return BTreeSearchItem(root.Right, elem)
}
