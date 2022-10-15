package dlinklist

import (
	"errors"
	"fmt"
)

type Node[T comparable] struct {
	prev, next *Node[T]
	val        T
}

type Dlist[T comparable] struct {
	head, tail *Node[T]
	len        int
}

func (dl *Dlist[T]) Push(val T) {
	if dl.len == 0 {
		dl.head = &Node[T]{nil, nil, val}
		dl.tail = dl.head
	} else {
		dl.head.prev = &Node[T]{nil, dl.head, val}
		dl.head = dl.head.prev
	}
	dl.len++

}

func (dl *Dlist[T]) Append(val T) {
	if dl.len == 0 {
		dl.head = &Node[T]{nil, nil, val}
		dl.tail = dl.head
	} else {
		dl.tail.next = &Node[T]{dl.tail, nil, val}
		dl.tail = dl.tail.next
	}
	dl.len++

}

func (dl *Dlist[T]) Pop() (T, error) {
	if dl.len == 0 {
		return *new(T), errors.New("empty list")
	}
	dl.len--
	val := dl.head.val
	dl.head = dl.head.next
	if dl.len == 0 {
		dl.tail = nil
	} else {
		dl.head.prev = nil
	}
	return val, nil

}

func (dl *Dlist[T]) Poll() (T, error) {
	if dl.len == 0 {
		return *new(T), errors.New("empty list")
	}
	dl.len--
	val := dl.tail.val
	dl.tail = dl.tail.prev

	if dl.len == 0 {
		dl.head = nil
	} else {
		dl.tail.next = nil
	}

	return val, nil
}

func (dl *Dlist[T]) IndexOf(val T) int {
	for i, tmp := 0, dl.head; i < dl.len; i++ {
		if tmp.val == val {
			return i
		}
		tmp = tmp.next
	}
	return -1
}

func (dl *Dlist[T]) RemoveValue(val T) {
	for i := dl.IndexOf(val); i != -1; {
		dl.RemoveFromPosition(i)
		i = dl.IndexOf(val)
	}
}

func (dl Dlist[T]) String() string {
	return fmt.Sprintf("len: %v, head: %p, tail %p\n", dl.len, dl.head, dl.tail)
}

func (dl *Dlist[T]) PrintTraverseForward() {
	if dl.len == 0 {
		fmt.Println(nil)
		return
	}
	fmt.Printf("Dl len: %v\n", dl.len)
	for i, tmp := 0, dl.head; i < dl.len; i++ {
		fmt.Printf("Node %v; val %v\n", i, tmp.val)
		tmp = tmp.next
	}
}

func (dl *Dlist[T]) removeNode(node *Node[T]) error {
	if dl.len == 0 {
		return fmt.Errorf("index out of range")
	}

	if node.prev == nil {
		dl.Pop()
	}
	if node.next == nil {
		dl.Poll()
	}

	node.next.prev = node.prev
	node.prev.next = node.next
	node.next = nil
	node.prev = nil
	node.val = *new(T)

	dl.len--

	return nil
}

func (dl *Dlist[T]) RemoveFromPosition(index int) error {
	if index <= dl.len/2 {
		for i, tmp := 0, dl.head; i <= dl.len/2; i, tmp = i+1, tmp.next {
			if i == index {
				dl.removeNode(tmp)
				break
			}
		}
	} else {
		for i, tmp := dl.len-1, dl.tail; i > dl.len/2; i, tmp = i-1, tmp.prev {
			if i == index {
				dl.removeNode(tmp)
				break
			}

		}
	}

	return nil
}
