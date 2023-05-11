package piscine

func BTreeApplyPostorder(root *TreeNode, f func(...interface{}) (int, error)) {
	if root != nil {
		BTreeApplyInorder(root.Right, f)
		BTreeApplyInorder(root.Left, f)
		f(root.Data)
	}
}
