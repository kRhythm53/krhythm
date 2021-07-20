package main

import "fmt"

type Node struct {
	data string
	left *Node
	right *Node
}
func (root *Node)PreOrderTraversal(){
	if root !=nil{
		fmt.Printf("%s ", root.data)
		root.left.PreOrderTraversal()
		root.right.PreOrderTraversal()
	}
	return
}

func (root *Node)PostOrderTraversal(){
	if root !=nil{
		root.left.PostOrderTraversal()
		root.right.PostOrderTraversal()
		fmt.Printf("%s ", root.data)
	}
	return
}


func main() {

	tree := Node{"+", &Node{"a",nil,nil}, &Node{"-",&Node{"b",nil,nil}, &Node{"c",nil,nil}}}
	tree.PreOrderTraversal()
	fmt.Println()
	tree.PostOrderTraversal()
	fmt.Println()
}