package main

import (
	"fmt"

	"../bst"
)

func main() {
	bst1 := bst.New()
	bst1.Add(8)
	bst1.Add(3)
	bst1.Add(5)
	bst1.Add(4)
	bst1.Add(1)
	bst1.Add(9)
	bst1.Add(6)
	bst1.Add(7)
	bst1.Add(10)
	bst1.Add(2)
	// bst1.TraverseMid()
	// bst1.TraverseFront()
	// bst1.TraverseBack()
	// fmt.Println(bst1.Search(5))
	// fmt.Println(bst1.Search(99))
	bst1.Graph()
	node := bst1.Search(8)
	// fmt.Println(node)
	bst1.Remove(node)
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println()
	bst1.Graph()
	bst1.TraverseMid()
	// fmt.Println(bst1.ExistNode(node))

}
