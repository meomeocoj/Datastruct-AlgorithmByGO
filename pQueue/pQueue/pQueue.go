package pQueue

import (
	"fmt"
	"math"

	"github.com/emirpasic/gods/sets/treeset"
	"golang.org/x/exp/constraints"
)

type PQueue[T constraints.Ordered] struct {
	heap          []T
	mapValToIndex map[T]*treeset.Set
}

func NewPQueue[T constraints.Ordered](size int) *PQueue[T] {
	return &PQueue[T]{heap: make([]T, size), mapValToIndex: make(map[T]*treeset.Set)}
}

func NewPQueueWithVals[T constraints.Ordered](elems []T) *PQueue[T] {
	var mapVal map[T]*treeset.Set = make(map[T]*treeset.Set)

	pq := &PQueue[T]{heap: elems, mapValToIndex: mapVal}

	for index, val := range pq.heap {
		s, ok := mapVal[val]
		if !ok {
			mapVal[val] = treeset.NewWithIntComparator()
			mapVal[val].Add(index)
		} else {
			s.Add(index)
		}
	}
	for i := int(math.Max(0, float64(len(elems)-1))); i >= 0; i-- {
		pq.sink(i)
	}
	return pq
}

func (pq *PQueue[T]) IsEmpty() bool {
	return len(pq.heap) == 0
}

func (pq *PQueue[T]) Free() {
	pq = nil
}

func (pq *PQueue[T]) Size() int {
	return len(pq.heap)
}

func (pq *PQueue[T]) Peek() (T, error) {
	if pq.IsEmpty() {
		return *new(T), fmt.Errorf("empty queue")
	}
	return pq.heap[0], nil
}

func (pq *PQueue[T]) Poll() (T, error) {
	return pq.RemoveAt(0)
}

func (pq *PQueue[T]) Contains(val T) bool {
	_, ok := pq.mapValToIndex[val]
	return ok
}

func (pq *PQueue[T]) Add(val T) error {
	position := len(pq.heap)
	pq.heap = append(pq.heap, val)
	_, ok := pq.mapValToIndex[val]
	if !ok {
		pq.mapValToIndex[val] = treeset.NewWithIntComparator()
	}
	pq.mapValToIndex[val].Add(position)
	pq.swim(position)
	return nil
}

func (pq *PQueue[T]) RemoveAt(k int) (T, error) {
	if len(pq.heap) == 0 {
		return *new(T), fmt.Errorf("empty queue")
	}

	lastPosition := len(pq.heap) - 1
	removed_data := pq.heap[k]

	pq.swap(k, lastPosition)
	pq.heap = pq.heap[:lastPosition]
	if len(pq.heap) == 0 {
		return removed_data, nil
	}
	if index := pq.sink(k); index == 0 {
		pq.swim(k)
	}

	return removed_data, nil
}

func (pq *PQueue[T]) sink(index int) int {
	for {
		left := 2*index + 1
		right := 2*index + 2
		largest := left

		// Compare node right and node left
		if right < len(pq.heap) && (pq.heap[right] > pq.heap[left]) {
			largest = right
		}

		// compare left and length of the heap, value of index and smallest
		if left >= len(pq.heap) || pq.heap[index] > pq.heap[largest] {
			break
		}
		// Swap the positions of the elems
		pq.swap(largest, index)
		index = largest
	}
	return index
}

func (pq *PQueue[T]) swim(index int) {
	parent := (index - 1) / 2
	for {
		if parent >= 0 && pq.heap[parent] < pq.heap[index] {
			pq.swap(parent, index)
			index = parent
			parent = (index - 1) / 2
		} else {
			break
		}
	}
}

func (pq *PQueue[T]) swap(i, j int) {
	// Swap index of the map before swaping value
	treeI := pq.mapValToIndex[pq.heap[i]]
	treeII := pq.mapValToIndex[pq.heap[j]]
	treeI.Remove(i)
	treeII.Remove(j)
	treeI.Add(j)
	treeII.Add(i)
	// Swap value in the heap
	pq.heap[i], pq.heap[j] = pq.heap[j], pq.heap[i]
}

func (pq *PQueue[T]) IsMaxHeap(i int) bool {

	if i >= len(pq.heap) {
		return true
	}

	left, right := 2*i+1, 2*i+2

	if left < len(pq.heap) && pq.heap[left] > pq.heap[i] {
		return false
	}
	if right < len(pq.heap) && pq.heap[right] > pq.heap[i] {
		return false
	}

	return pq.IsMaxHeap(left) && pq.IsMaxHeap(right)
}

func (pq *PQueue[T]) String() string {
	return fmt.Sprintf("heap: %v, mapValToIndex: %v", pq.heap, pq.mapValToIndex)
}
