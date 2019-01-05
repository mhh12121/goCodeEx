package main

/*
Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3
*/

func isSymmetric(root *Node) bool {
	if root == nil {
		return true
	}

	return traverse(root.Left, root.Right)
}

func traverse(rootL *Node, rootR *Node) bool {
	if rootL == nil && rootR == nil {
		return true
	}
	if rootL == nil || rootR == nil || rootL.Val != rootR.Val {
		return false
	}

	return traverse(rootL.Left, rootR.Right) && traverse(rootL.Right, rootR.Left)

}
