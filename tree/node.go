package main

import "fmt"

type  treeNode struct {
	value int
	left,right *treeNode
}

/**
 * 使用自定义工厂函数
 */
func createNode(value int) *treeNode{
	return &treeNode{value:value}
}

func (node *treeNode) traverse(){
	if node == nil {
		return
	}

	node.left.traverse()
	node.print()
	node.right.traverse()

}

func (node treeNode) print(){
	fmt.Printf("%d ",node.value)
}

//如果接受指针，则直接改掉值
func (node *treeNode) setValue(value int){

	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}

	node.value = value
}



func main(){
	var root treeNode

	root = treeNode{value: 3}

	root.left = &treeNode{}
	root.right = &treeNode{5,nil,nil}

	root.right.left = new(treeNode)

	root.left.right = createNode(2)

	fmt.Println(root)

	root.print()

	root.traverse()
}


