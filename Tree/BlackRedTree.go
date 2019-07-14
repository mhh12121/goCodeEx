package main

import (
	"errors"
	"fmt"
	"log"
)

const (
	RED   bool = true
	BLACK bool = false
)

type BRNode struct {
	color  bool
	left   *BRNode
	right  *BRNode
	parent *BRNode
	val    int
}

func (brnode *BRNode) getParent() *BRNode { //之前都需要自己确认是否为空
	return brnode.parent
}
func (brnode *BRNode) getGrandParent() *BRNode { //之前都需要自己确认是否为空
	return brnode.parent.parent
}
func (brnode *BRNode) getsibling() *BRNode {
	if brnode.parent.left == brnode {
		return brnode.parent.right
	}
	return brnode.parent.left
}
func (brnode *BRNode) getUncle() *BRNode {
	if brnode.parent.parent != nil {
		if brnode.parent == brnode.parent.parent.left {
			return brnode.parent.parent.right
		}
		return brnode.parent.parent.left
	}
	return nil
}

/*
	2
	  \
	   8
	  /  \
	 7   10

==>
	  8
    /   \
   2     10
    \
     7
*/
//如上面例子，2即是传入的brnode
func (brnode *BRNode) rotateLeft() (*BRNode, error) { //AVL
	var temp *BRNode
	if brnode == nil {
		return temp, nil //这个时候是叶子节点，没有任何内容而已
	}
	if brnode.right == nil { //左旋的时候右节点不能为空
		return temp, errors.New("rightchild cannot be empty if do left rotate")
	}
	parent := brnode.parent //暂存当前节点的父节点，以防旋转后当前节点被替换了找不到父节点
	var isLeft bool         //这个用来判断当前节点是左子节点与否
	if parent != nil {
		isLeft = brnode == parent.left
	}
	grandson := brnode.right.left
	brnode.right.left = brnode   //右子节点的左子节点指向自己（断开7,8,连接2,8）
	brnode.parent = brnode.right //自己的父节点指向了自己的右节点，这样自己的右节点就取代了自己（连接2,8）
	brnode.right = grandson      //把暂存的放到自己的右节点处，因为自己的右节点已经取代了自己（连接2,7）

	//是否是根节点
	if parent == nil {
		//更新新的根节点的父节点（即为nil）
		brnode.parent.parent = nil //brnode现在是在左子节点
		temp = brnode.parent       //暂存新的根节点
	} else { //更新原父节点的左或右子树，取决于旋转的是左子树还是右子树
		if isLeft {
			parent.left = brnode.parent
		} else {
			parent.right = brnode.parent
		}
		brnode.parent.parent = parent //更新新的节点的父节点
	}
	return temp, nil //不是根节点的话temp是空的
}

/*
	10
   /
  7
 / \
2	8

==>
	  7
    /   \
   2     10
         /
        8
*/

func (brnode *BRNode) rotateRight() (*BRNode, error) { //AVL
	var temp *BRNode
	if brnode == nil {
		return temp, nil //这个时候是叶子节点，没有任何内容而已
	}
	if brnode.left == nil { //右旋的时候左节点不能为空
		return temp, errors.New("lefttchild cannot be empty if do right rotate")
	}
	parent := brnode.parent //暂存当前节点的父节点，以防旋转后当前节点被替换了找不到父节点
	var isLeft bool         //这个用来判断当前节点是左子节点与否
	if parent != nil {
		isLeft = brnode == parent.left
	}
	//----------------主要区别在这里----------------

	grandson := brnode.left.right
	brnode.left.right = brnode  //右子节点的左子节点指向自己
	brnode.parent = brnode.left //自己的父节点指向了自己的右节点，这样自己的右节点就取代了自己
	brnode.left = grandson      //把暂存的放到自己的右节点处，因为自己的右节点已经取代了自己

	//--------------------------------------------
	//根节点是否被换了???
	if parent == nil {
		brnode.parent.parent = nil //brnode现在是在左子节点
		temp = brnode.parent       //
	} else { //更新原父节点的左或右子树，取决于旋转的是左子树还是右子树
		if isLeft {
			parent.left = brnode.parent
		} else {
			parent.right = brnode.parent
		}
		brnode.parent.parent = parent //更新新的节点的父节点
	}
	return temp, nil
}

type BRTree struct { //树的开始
	root *BRNode
}

//该节点下整个树的旋转
func (brtree *BRTree) rotateLeft(node *BRNode) {
	if tmproot, err := node.rotateLeft(); err != nil {
		log.Println(err)
	} else {
		if tmproot != nil { //根节点
			brtree.root = tmproot
		}
	}
}
func (brtree *BRTree) rotateRight(node *BRNode) {
	if tmproot, err := node.rotateRight(); err != nil {
		log.Println(err)
	} else {
		if tmproot != nil {
			brtree.root = tmproot
		}
	}
}

//主要思路，先像AVL一样插入，之后再调整
func (brtree *BRTree) insertNode(pnode *BRNode) {
	tempRoot := brtree.root
	var temp *BRNode
	for tempRoot != nil { //已有根节点，就往下loop找到插入点
		temp = tempRoot //每次更新
		if pnode.val > tempRoot.val {
			tempRoot = tempRoot.right

		} else {
			tempRoot = tempRoot.left
		}
	}
	pnode.parent = temp //找到最后，把pnode的parent指向找到的最后一个点
	if temp != nil {    //不是根节点
		if temp.val < pnode.val { //判断放在左子树还是右子树
			temp.right = pnode
		} else {
			temp.left = pnode
		}
		pnode.color = RED //直接把这个颜色先设置为红色！！！主要为了满足：
		// 从一个节点到该节点的所有叶子节点的路径上包含相同数目的黑节点（可以试想一下，最下方插入一个红色节点，因为目前红色节点的叶子节点肯定是黑色的nil，所以必会满足这个特性）
		brtree.insertCheck(pnode) //再检查
	} else { //根节点
		pnode.color = BLACK
		brtree.root = pnode
	}

}

//插入时进行检查
func (brtree *BRTree) insertCheck(pnode *BRNode) {
	//1. 该节点没有父节点，即为root（与上面insertNode插入点的检查并不重复，因为接下来会进行递归，这个属于边界条件，所以这个检查root必不可少）
	if pnode.parent == nil {
		brtree.root = pnode
		brtree.root.color = BLACK
		return
	}
	//2.父节点是黑色直接添加(不用管)，红色接着处理
	if pnode.parent.color == RED {
		//2.1 父，叔叔节点不为空而且其颜色是红色 ,则将该父叔都改为黑色,将祖父改成红色
		/*
							  10,B
							 /     \
						    6,B	    15,B
						   /   \     /  \
						4,R   8,R  11,R  19,R
						/
					pnode,R

					==>父叔节点都变成黑色，祖父变成红色
							  10,B
							 /     \
						    6,R	    15,B
						   /   \     /  \
						4,B   8,B  11,R  19,R
			             /
					  pnode,R

					==>递归，从祖父（6）开始，因为此时父节点属于根节点，是黑色，所以完成

					           10,B
							 /     \
						    6,R	    15,B
						   /   \     /  \
						4,B   8,B  11,R  19,R
			             /
					  pnode,R
		*/
		if pnode.getUncle() != nil && pnode.getUncle().color == RED {
			pnode.parent.color = BLACK
			pnode.getUncle().color = BLACK
			pnode.getGrandParent().color = RED
			brtree.insertCheck(pnode.getGrandParent()) //对该树上的节点递归上去处理
		} else {
			//2.2 父节点为红色，叔叔节点不存在或者是黑色

			isLeft := pnode == pnode.parent.left
			isParentLeft := pnode.parent == pnode.getGrandParent().left

			if isLeft && isParentLeft {
				//2.2.1 是左子树，其父节点也是左子树
				/*       ...
						  /
				        8,B
					    /  \
					   7,R   /nil
					  /
					pnode,R
				*/
				//=》
				/*         /
					     7,B
				         /   \
				     pnode,R  8,R
				               \
				             /nil
				*/
				pnode.parent.color = BLACK
				pnode.getGrandParent().color = RED
				brtree.rotateRight(pnode.getGrandParent())

			} else if !isLeft && !isParentLeft {
				//2.2.2 不是左子树，父亲也不是左子树
				/*
						  7,B
					      /  \
					2,B/nil   8,R
					           \
					           10,pnode
				*/
				pnode.parent.color = BLACK
				pnode.getGrandParent().color = RED
				brtree.rotateLeft(pnode.getGrandParent())

			} else if isLeft && !isParentLeft {
				//2.2.3 是左子树，但父亲不是左子树
				/*
						8,B
					    /  \
					6,B/nil   10,R
					        /
					      pnode，R
				*/
				/*
						pnode,R
					    /  \
					6,B/nil  8,B
					            \
					            10,R
				*/
				brtree.rotateRight(pnode.parent)
				//实际上变成了2.2.2
				brtree.rotateLeft(pnode.parent) //pnode现在在原先pnode的父节点位置
				pnode.color = BLACK
				pnode.left.color = RED
				pnode.right.color = RED

			} else if !isLeft && isParentLeft {
				//2.2.4不是左子树，但父亲是左子树
				/*
							 \
					        8,B
						    /  \
						   7,R   10,B/nil
						      \
						      pnode，R
				*/
				//==>
				/*
						 8，B
					    /    \
					  pnode,R  10,B/nil
					  /
					7,R
				*/
				/*
						 pnode，R
					    /    \
					   8,B   10,B/nil
					  /
					7,R
				*/

				brtree.rotateLeft(pnode.parent)
				//实际上变成了2.2.1
				brtree.rotateRight(pnode.parent)
				pnode.color = BLACK
				pnode.left.color = RED
				pnode.right.color = RED

			}
		}

	}
}

/*
要删除的是n，有一个子节点（这里是左子节点）
	 					n
			    /    \
			   S   	nil
			  /  \
			 Sl		SR
1. 将S.parent=n.Parent
2. 如果n=n.Parent.left,让n.Parent.left=S,右边同理

要删除的是n，有两个子节点
	 					n
			    /    \
			   S   	  X
			  /  \   /  \
			 Sl	 SR  XL  XR

1. 选择后继节点，即X作为替换
2. 找到n的右子树X的最左非空子树，这里是XL，将XL的值复制到n，将删除的节点改为XL，
就可以转换成节点包含0个或1个节点的问题了

*/

//二叉查找树删除修正
func (brtree *BRTree) removeNode(n *BRNode) { //传入数据和之前一样可以构造一个brnode对象进来
	//
	var child, parent *BRNode
	var color bool
	if n.left != nil && n.right != nil {
		replace := n
		replace = n.right         //被删除的后继节点
		for replace.left != nil { //得到最左的子节点
			replace = replace.left
		}
		//此时实际上是要删除掉replace节点
		if n.parent != nil { //not root node
			if n.parent.left == n { //分别连接被删除节点的父节点和被删除节点的后继节点的最左的子节点
				n.parent.left = replace
			} else if n.parent.right == n {
				n.parent.right = replace
			}
		} else {
			brtree.root = replace
		}

		child = replace.right
		parent = replace.parent
		color = replace.color

		/*

		 */
		if parent == n {
			//删除节点是它后继节点的父节点
			parent = replace
		} else {

			if child != nil {
				//child 不为空
				child.parent = parent
			}
			parent.left = child     //把child节点放在parent的左节点处
			replace.right = n.right //被删除的节点右节点由replace节点指向
			n.right = replace
		}
		replace.parent = n.parent
		replace.color = n.color
		replace.left = n.left
		n.left.parent = replace
		if color == BLACK {
			removeCheck(child, parent)
		}
		n = nil
		return
	}

	if n.left != nil { //只有左节点非空
		child = n.left
	} else { //右节点非空
		child = n.right
	}

	parent = n.parent
	color = n.color
	if child != nil {
		child.parent = parent
	}

	if parent != nil {
		if parent.left == n {
			parent.left = child
		} else {
			parent.right = child
		}
	} else {
		brtree.root = child
	}

	if color == BLACK {
		removeCheck(child, parent)
	}
	n = nil

}

//颜色删除修正
func removeCheck(n *BRNode, parent *BRNode) {

}

func Search(root *BRNode, data int) *BRNode {
	return nil
}
func (brtree *BRTree) RemoveKey(data int) {
	node := Search(brtree.root, data)
	if node == nil {
		brtree.removeNode(node)
	}

}

func PreOrderTraverse(root *BRNode) {
	if root == nil {
		return
	}
	c := "BLACK"
	if root.color {
		c = "RED"
	}
	fmt.Printf("val:%v,color:%v,", root.val, c)
	PreOrderTraverse(root.left)
	PreOrderTraverse(root.right)
}
func InOrderTraverse(root *BRNode) {
	if root == nil {
		return
	}

	InOrderTraverse(root.left)
	c := "BLACK"
	if root.color {
		c = "RED"
	}
	fmt.Printf("val:%v,color:%v,", root.val, c)
	InOrderTraverse(root.right)
}
func PostOrderTraverse(root *BRNode) {
	if root == nil {
		return
	}

	PostOrderTraverse(root.left)

	PostOrderTraverse(root.right)
	fmt.Printf("val:%v,color:%v", root.val, root.color)
}

func buildBlackRedTree() {
	tnode := &BRNode{val: 10, color: BLACK}
	btree := &BRTree{}
	tnode2 := &BRNode{val: 6}
	tnode3 := &BRNode{val: 15}
	tnode4 := &BRNode{val: 11}
	// // tnode5 := &BRNode{val: 19}
	tnode6 := &BRNode{val: 8}
	tnode7 := &BRNode{val: 4}
	tnode8 := &BRNode{val: 2}
	tnode9 := &BRNode{val: 1}
	tnode10 := &BRNode{val: 0}
	btree.insertNode(tnode)
	btree.insertNode(tnode2)
	btree.insertNode(tnode3)
	btree.insertNode(tnode4)
	// // btree.insertNode(tnode5)
	btree.insertNode(tnode6)
	btree.insertNode(tnode7)
	btree.insertNode(tnode8)
	btree.insertNode(tnode9)
	btree.insertNode(tnode10)
	PreOrderTraverse(btree.root)
	fmt.Println("\n")
	InOrderTraverse(btree.root)
	fmt.Println("\n")

}
