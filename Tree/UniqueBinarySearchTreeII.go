package main

//Given an integer n, generate all structurally unique BST's (binary search trees) that store values 1 ... n.

/*Example:

Input: 3
Output:
[
  [1,null,3,2],
  [3,2,null,1],
  [3,1,null,null,2],
  [2,1,3],
  [1,null,2,null,3]
]
Explanation:
The above output corresponds to the 5 unique BST's shown below:

   1         3     3      2      1
    \       /     /      / \      \
     3     2     1      1   3      2
    /     /       \                 \
   2     1         2                 3

*/
func generateTrees(n int) []*Node {
	return genSubTree(1, n)
}
func genSubTree(start, end int) []*Node {
	res := make([]*Node, 0)
	if start > end {
		res = append(res, nil)
		return res
	}
	for i := start; i <= end; i++ {
		leftSubTree := genSubTree(start, i-1)
		rightSubTree := genSubTree(i+1, end)
		for left := range leftSubTree {
			for right := range rightSubTree {
				root := &Node{Val: i, Left: leftSubTree[left], Right: rightSubTree[right]}
				res = append(res, root)
			}
		}
	}
	return res
}
