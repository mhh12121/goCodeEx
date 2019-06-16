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

type stack []*Node

func (s *stack) pop() (v *Node) {

	if len(*s) == 0 {
		return nil
	}
	v = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return
}
func (s *stack) push(node *Node) {
	*s = append(*s, node)
}
func isSymmetricIterative(root *Node) bool {
	if root == nil {
		return true
	}
	var s stack
	if root.Left != nil {
		if root.Right == nil {
			return false
		}
		s.push(root.Left)
		s.push(root.Right)
	} else if root.Right != nil {
		return false
	}
	for len(s) != 0 {
		if len(s)%2 != 0 {
			return false
		}
		right := s.pop()
		left := s.pop()

		if right.Val != left.Val {
			return false
		}
		if left.Left != nil {
			if right.Right == nil {
				return false
			}
			s.push(left.Left)
			s.push(right.Right)
		} else if right.Right != nil {
			return false
		}

		if left.Right != nil {
			if right.Left == nil {
				return false
			}
			s.push(left.Right)
			s.push(right.Left)
		} else if right.Left != nil {
			return false
		}

	}
	return true

}
