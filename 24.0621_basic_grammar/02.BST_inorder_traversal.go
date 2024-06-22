func inorderTraversal_inner(root *TreeNode, output *[]int) {
    if root == nil {
        return
    }
    inorderTraversal_inner(root.Left, output)
    *output = append(*output, root.Val)
    inorderTraversal_inner(root.Right, output)
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    var result []int
    inorderTraversal_inner(root, &result)
    return result   
}