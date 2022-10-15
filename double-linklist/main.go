package main

import (
	"github.com/meomeocoj/double-linklist/dlinklist"
)

func main() {
	dl := dlinklist.Dlist[int]{}
	dl.Push(5)
	dl.Push(4)
	dl.Push(4)
	dl.Push(3)
	dl.Push(2)
	dl.Push(1)
	dl.PrintTraverseForward()
	// fmt.Println("After pop and poll")
	// dl.printTraverseForward()
	dl.RemoveValue(4)
	// fmt.Println(err.Error())
	dl.PrintTraverseForward()

}
