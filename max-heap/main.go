package main

import "log"

// MaxHeap DEFINITION:
// A max-heap is a complete binary tree in which the value of each node is greater than or equal to the values of its children. In other words, the root node contains the maximum value in the heap.
// A complete binary tree is a binary tree in which all levels are completely filled except possibly the last level, which is filled from left to right.
// Max-heaps are commonly used to implement priority queues, where the highest priority element is always at the front of the queue. They are also used in sorting algorithms such as heapsort.
// In a max-heap, the maximum element can be found in constant time by simply accessing the root node. Inserting a new element into the heap involves adding it to the bottom level of the tree and then "bubbling up" the element by swapping it with its parent until the heap property is restored. Removing the maximum element involves swapping it with the last element in the heap, removing the last element, and then "bubbling down" the new root element by swapping it with its larger child until the heap property is restored.
// The time complexity of basic operations on a max-heap is O(log n), where n is the number of elements in the heap.
type MaxHeap struct {
	heap []int
}

// Insert adds new value to the heap and sort it again
func (h *MaxHeap) Insert(value int) {
	h.heap = append(h.heap, value)
	h.bubbleUp(len(h.heap) - 1)
}

func (h *MaxHeap) parentIndex(index int) int {
	i := (index - 1) / 2
	if i < 0 {
		return 0
	}
	return i
}

func (h *MaxHeap) rightIndex(index int) int {
	return (index * 2) + 2
}
func (h *MaxHeap) leftIndex(index int) int {
	return (index * 2) + 1
}

func (h *MaxHeap) parentOf(index int) int {
	return h.heap[h.parentIndex(index)]
}

// re-sort the heap when an element added to the end of the heap
func (h *MaxHeap) bubbleUp(index int) {
	if h.heap[index] > h.parentOf(index) { // in value of index is greater that its parent, swap them
		h.heap[index], h.heap[h.parentIndex(index)] = h.heap[h.parentIndex(index)], h.heap[index]
		h.bubbleUp(h.parentIndex(index))
		return
	}
	return
}

func (h *MaxHeap) bubbleDown(index int) {
	heapLen := len(h.heap)
	if heapLen == 1 { // last element remaining
		return
	}

	leftIndex := h.leftIndex(index)
	rightIndex := h.rightIndex(index)
	if heapLen-1 < leftIndex { // element has no left and no right child

	} else if heapLen-1 < rightIndex { // element has just left child
		if h.heap[leftIndex] > h.heap[index] { // left element is greater that its parent, swap them
			h.heap[leftIndex], h.heap[index] = h.heap[index], h.heap[leftIndex]
			h.bubbleDown(leftIndex)
			return
		}
	} else { // element has both right and left children
		left := h.heap[leftIndex]
		right := h.heap[rightIndex]

		// select max of right and left
		chosenValue := left
		chosenIndex := leftIndex
		if right > left {
			chosenValue = right
			chosenIndex = rightIndex
		}

		// compare current value with chosen value
		currentValue := h.heap[index]
		if currentValue < chosenValue { // swap them
			h.heap[index], h.heap[chosenIndex] = h.heap[chosenIndex], h.heap[index]
			h.bubbleDown(chosenIndex)
			return
		}
	}

}

// ExtractMax removes the max value from the heap and resort it again
func (h *MaxHeap) ExtractMax() int {
	if len(h.heap) == 0 {
		return 0
	}

	max := h.heap[0]

	length := len(h.heap)
	if length == 1 {
		return max
	}

	// move last index to the first element of the array
	h.heap[0] = h.heap[length-1]
	// remove the last index after moving it to the beginning
	h.heap = h.heap[:length-1]
	// bubble down the first element to re-sort the heap
	h.bubbleDown(0)
	return max
}

func main() {

	items := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	h := MaxHeap{}

	for _, i := range items {
		h.Insert(i)
		log.Println(h)
	}

	for i := 0; i < 5; i++ {
		h.ExtractMax()
		log.Println(h)
	}
}
