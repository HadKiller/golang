package tree

import "fmt"

//二叉树
type NodeTree struct {
	 left,right *NodeTree
	 value int
}

func (treeNode NodeTree) printNode()  {
	fmt.Println(treeNode.value)
}



func CreateTree() *NodeTree{

	root:=NodeTree{nil,nil,1}
	root.left=&NodeTree{nil,nil,2}
	root.left.left=&NodeTree{nil,nil,4}
	root.left.right=&NodeTree{nil,nil,5}
	root.left.right.right=&NodeTree{nil,nil,6}
	root.right=&NodeTree{value:3}
	return &root
}
//先序遍历
func NLR(tree *NodeTree){
	if tree!=nil{
		fmt.Print(tree.value," ")
		NLR(tree.left)
		NLR(tree.right)
	}
}
//中序遍历
func LNR(tree *NodeTree){
	if tree!=nil{

		tree.printNode()
		LNR(tree.left)
		fmt.Print(tree.value," ")
		LNR(tree.right)
	}
}
//后序遍历
func LRN(tree *NodeTree){
	if tree!=nil{
		LRN(tree.left)
		LRN(tree.right)
		fmt.Print(tree.value," ")


	}
}