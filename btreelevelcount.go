package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftLevels := BTreeLevelCount(root.Left)
	rightLevels := BTreeLevelCount(root.Right)
	if leftLevels > rightLevels {
		return leftLevels + 1
	}
	return rightLevels + 1
}
