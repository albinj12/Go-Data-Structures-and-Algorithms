package main

import (
	"fmt"
	"math"
)

type node struct{
	data int
	height int
	left *node
	right *node
}

type avlTree struct{
	root *node
}

func main(){
	avlt := avlTree{}
	// avlt.insert(26)
	// avlt.insert(5)
	// avlt.insert(38)
	// avlt.insert(46)
	// avlt.insert(48)
	// avlt.insert(63)
	// avlt.insert(49)
	// avlt.insert(59)
	// avlt.insert(76)
	// avlt.insert(89)
	// avlt.delete(46)
	// avlt.insert(60)
	// avlt.insert(50)
	// avlt.preorderTraversal(avlt.root)
}

func (avlt *avlTree) insert(value int){
	if avlt.root == nil{
		avlt.root = &node{
			data: value,
			height: 1,
		}
		return
	}

	avlt.root = insertNode(avlt.root, value)
}

func insertNode(avltNode *node, value int) *node{	
	if avltNode == nil{
		return &node{
			data: value,
			height: 1,
		}
	}
	if value < avltNode.data {
		avltNode.left = insertNode(avltNode.left, value)
	}else{
		avltNode.right = insertNode(avltNode.right, value)
	}

	avltNode.height = int(math.Max(getHeight(avltNode.left), getHeight(avltNode.right))) + 1

	balanceFactor := getBalance(avltNode)

	if balanceFactor > 1 && value < avltNode.left.data{
		avltNode = rightRotate(avltNode)
	}

	if balanceFactor < -1 && value > avltNode.right.data{
		avltNode = leftRotate(avltNode)
	}

	if balanceFactor > 1 && value > avltNode.left.data{
		avltNode.left = leftRotate(avltNode.left)
		avltNode = rightRotate(avltNode)
	}

	if balanceFactor < -1 && value < avltNode.right.data{
		avltNode.right = rightRotate(avltNode.right)
		avltNode = leftRotate(avltNode)
	}
	return avltNode
}

func (avlt *avlTree) delete(value int){
	if avlt.root == nil{
		return
	}

	deleteNode(avlt.root, value)
	
}

func deleteNode(avltNode *node, value int) *node{
	if value < avltNode.data{
		avltNode.left = deleteNode(avltNode.left, value)
	}else if value > avltNode.data{
		avltNode.right = deleteNode(avltNode.right, value)
	}else{
		if avltNode.left == nil{
			return avltNode.right
		}else if avltNode.right == nil{
			return avltNode.left
		}else{
			avltNode.data = findInorderSuccessor(avltNode.right)
			avltNode.right = deleteNode(avltNode.right, avltNode.data)
		}
	}
	avltNode.height = int(math.Max(getHeight(avltNode.left), getHeight(avltNode.right)))+1
	balanceFactor := getBalance(avltNode)
	if balanceFactor > 1 && getBalance(avltNode.left) >= 0{
		avltNode = rightRotate(avltNode)
	}

	if balanceFactor < -1 && getBalance(avltNode.right) <=0{
		avltNode = leftRotate(avltNode.left)
	}

	if balanceFactor > 1 && getBalance(avltNode.left) < 0{
		avltNode.left = leftRotate(avltNode.left)
		avltNode = rightRotate(avltNode)
	}

	if balanceFactor < -1 && getBalance(avltNode.right) >0{
		avltNode.right = rightRotate(avltNode.right)
		avltNode = leftRotate(avltNode)
	}
	return avltNode
}

func getHeight(avltNode *node) float64{
	if avltNode == nil{
		return 0
	}

	return float64(avltNode.height)
}

func getBalance(avltNode *node) int{
	if avltNode.left == nil && avltNode.right == nil{
		return 0
	}else if avltNode.left == nil{
		return 0 - avltNode.right.height
	}else if avltNode.right == nil{
		return avltNode.left.height
	}else{
		return avltNode.left.height - avltNode.right.height
	}
}

func rightRotate(avltNode *node) *node{
	temp1 := avltNode
	temp2 := avltNode.left
	temp3 := avltNode.left.right

	avltNode = temp2
	avltNode.right = temp1
	avltNode.right.left = temp3

	avltNode.right.height = int(math.Max(getHeight(avltNode.right.left), getHeight(avltNode.right.right)))+1
	avltNode.height = int(math.Max(getHeight(avltNode.left), getHeight(avltNode.right)))+1
	return avltNode
}

func leftRotate(avltNode *node) *node{
	temp1 := avltNode
	temp2 := avltNode.right
	temp3 := avltNode.right.left

	avltNode = temp2
	avltNode.left = temp1
	avltNode.left.right = temp3

	avltNode.left.height = int(math.Max(getHeight(avltNode.left.left), getHeight(avltNode.left.right)))+1
	avltNode.height = int(math.Max(getHeight(avltNode.left), getHeight(avltNode.right)))+1
	return avltNode
}

func findInorderSuccessor(avltNode *node) int {
	for avltNode.left != nil{
		avltNode = avltNode.left
	}

	return avltNode.data
}

func (avlt *avlTree) inorderTraversal(avltNode *node){
	if avltNode != nil{
		avlt.inorderTraversal(avltNode.left)
		fmt.Println(avltNode.data)
		avlt.inorderTraversal(avltNode.right)
		
	}
}

func (avlt *avlTree) preorderTraversal(avltNode *node){
	if avltNode != nil{
		fmt.Println(avltNode)
		avlt.preorderTraversal(avltNode.left)
		avlt.preorderTraversal(avltNode.right)
		
	}
}