package main

import (
	"fmt"

	"github.com/meomeocoj/pQueue/pQueue"
)

func main() {
	pq := pQueue.NewPQueueWithVals([]int{1, 5, 5, 2, 3, 4, 7, 10, 50, 70, 80})
	pq.Add(10)
	fmt.Println(pq)
	pq.Add(2)
	fmt.Println(pq)
	pq.Poll()
	fmt.Println(pq)
	pq.RemoveAt(4)
	fmt.Println(pq)
	fmt.Println(pq.IsMaxHeap(0))
}
