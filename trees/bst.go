package trees

import "fmt"

type Node struct {
	val int
	left *Node
	right *Node
}


func Insert(val int, root *Node) *Node{
	if root == nil {
		node := Node {val,nil,nil}
		root = &node
		return root;
	}

	if(val < root.val) {
		root.left= Insert(val,root.left)
	}else{
		root.right = Insert(val,root.right)
	}

	return root
}


func Inorder(root *Node){

	if root==nil {
		return
	}
	Inorder(root.left)
	fmt.Println(root.val)
	Inorder(root.right)

}