package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func createNode(value int) *treeNode { //一个工厂函数   可以模拟构造函数
	return &treeNode{value: value} //局部变量也可以返回给别人用
}

//类的方法  go语言所有的函数都是传递值，不能改变原始的
func (node treeNode) print() { //第一个括号为接收者
	fmt.Printf("%d ", node.value)
}

//若想改变原始的则需要
func (node *treeNode) setValue(value int) {
	node.value = value
}

/*
值接收者vs指针接收者
- 要改变内容必须使用指针接收者
- 结构过大也考虑使用指针接收者（性能问题）

go中一个目录下只能有一个包


*/

//使用中序遍历，遍历树
func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()

}

func main() {
	var root treeNode
	root = treeNode{value: 3}
	root.left = &treeNode{} //由于是指针，所以这块需要取地址
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode) //new是一个内建函数
	root.left.right = createNode(2)
	root.right.left.setValue(4)
	node := []treeNode{ //节点的切片
		{value: 3},
		{},
		{6, nil, nil},
	}

	fmt.Println(node)
	fmt.Println("================================")

	root.traverse()
}

//为结构定义的方法必须放在同一个包内，但可以是不同文件
//如何扩充系统类型或者别人的类型
//	定义别名
//	使用组合
type MyTreeNode struct {
	node *treeNode
}

//后序遍历
func (myNode *MyTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return

	}

	node := MyTreeNode{myNode.node.left}
	node.postOrder()
	node2 := MyTreeNode{myNode.node.right}
	node2.postOrder()
	myNode.node.print()
}
