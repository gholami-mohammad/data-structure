package main

import "fmt"

// Linked lists have several advantages over arrays in Golang[1][2][3][4][6]. Here are some of the benefits of linked lists over arrays:

// - Dynamic size: Linked lists can grow or shrink dynamically, allowing for easy addition or removal of elements without worrying about resizing or reallocation[1][2][6].
// - Efficient insertion and deletion: Linked lists excel in insertion and deletion operations, making them ideal for scenarios with large data sets[1].
// - Scalability: Linked lists can add or remove elements at any position, making them more scalable than arrays[2].
// - Memory efficiency: Linked lists are more memory-efficient than arrays because they can store heterogeneous data and only require memory for the data and the pointers to the next node[6].

// However, linked lists also have some disadvantages compared to arrays[2]:

// - Lower efficiency at times: For certain operations, such as searching for an element or iterating through the list, arrays can be more efficient than linked lists[2].
// - Memory usage: Linked lists require more memory than arrays because each node requires a pointer to the next node in the list[2].

// In summary, linked lists are a powerful data structure that can be used to manage changing data structures efficiently in Golang. They offer unique advantages such as dynamic size, efficient insertion and deletion, scalability, and memory efficiency. However, they also have some disadvantages such as lower efficiency for certain operations and higher memory usage compared to arrays.

// Citations:
// [1] https://hackthedeveloper.com/golang-linked-list/
// [2] https://www.geeksforgeeks.org/advantages-and-disadvantages-of-linked-list/
// [3] https://stackoverflow.com/questions/166884/array-versus-linked-list
// [4] https://www.developer.com/languages/linked-list-go/
// [5] https://www.reddit.com/r/golang/comments/25aeof/building_a_stack_in_go_slices_vs_linked_list/
// [6] https://sweetcode.io/implementing-the-linkedlist-data-structure-in-go/

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	n.next = l.head
	l.head = n
	l.length++
}

func (l *linkedList) deleteHead() {
	if l.head == nil {
		return
	}

	l.head = l.head.next
	l.length--
}

func (l *linkedList) deleteAllByData(data int) {
	if l.head == nil {
		return
	}

	n := l.head
	if l.head.data == data {
		l.head = l.head.next
		l.length--
		l.deleteAllByData(data)
		return
	}
	prev := l.head
	n = n.next
	for {
		if n.data == data {
			prev.next = n.next
			n = prev.next
			l.length--
		} else {
			prev = n
			n = n.next
		}

		if prev.next == nil {
			break
		}
	}
}

func (l linkedList) print() {
	if l.head == nil {
		return
	}
	n := l.head
	for {
		fmt.Printf("%d ", n.data)
		n = n.next
		if n == nil {
			break
		}
	}
	fmt.Println("")
}

func (l linkedList) search(data int) *node {
	if l.head == nil {
		return nil
	}

	n := l.head
	for {
		if n.data == data {
			return n
		}

		if n.next == nil {
			break
		}
		n = n.next
	}

	return nil
}

func main() {
	data := []int{9, 32, 52, 62, 3, 235, 351, 52, 11}

	list := linkedList{}

	// generate nodes of linked list based on the given array and prepend them to the list
	for _, v := range data {
		n := &node{data: v}

		list.prepend(n)
	}

	list.print()

	fmt.Println("Deleting by data = 52")
	list.deleteAllByData(52)
	list.print()
	fmt.Printf("Linked list length is: %d\n", list.length)
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")

	// no node has data equal to 500
	fmt.Println("Deleting by data = 500")
	list.deleteAllByData(500)
	list.print()
	fmt.Printf("Linked list length is: %d\n", list.length)
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")

	// delete by data of head
	fmt.Println("Deleting by data = 11; target is head node")
	list.deleteAllByData(11)
	list.print()
	fmt.Printf("Linked list length is: %d\n", list.length)
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")

	// delete head
	fmt.Println("Deleting head")
	list.deleteHead()
	list.print()
	fmt.Printf("Linked list length is: %d\n", list.length)
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")

	// search
	fmt.Println("Searching for 235")
	r0 := list.search(235)
	fmt.Printf("search result: %v\n", r0)
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")

	// search
	fmt.Println("Searching for 666")
	r1 := list.search(666)
	fmt.Printf("search result: %v\n", r1)
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")

	fmt.Println("Empty list")
	emptyList := linkedList{}
	emptyList.deleteAllByData(10)
	emptyList.print()
	fmt.Printf("Empty Linked list length is: %d\n", emptyList.length)
	fmt.Println("=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=")

}
