package main

import (
	"fmt"
)

// Creating Struct for Node
type node struct {
	next  *node
	value interface{}
}

//Creating Struct For link list
type linkList struct {
	length int // will be used for counting the length
	head   *node
	tail   *node
}

//for getting the current node value
func (n *node) Value() interface{} {
	return n.value
}

// for getting the next node
func (n *node) Next() *node {
	return n.next
}

// this will set the next node as current node , if its present
func (n *node) Set(next *node) {
	n.next = next
}

//for knowing the length of link list
func (l *linkList) Length() int {
	return l.length
}

// to push a new node  at the end of the link list
func (l *linkList) Push(val interface{}) {
	n := &node{value: val}

	if l.head == nil {
		l.head = n
	} else {
		l.tail.Set(n)
	}
	l.tail = n
	l.length = l.length + 1
}

// this will pop the last element in the link list
func (l *linkList) Pop() {
	node := l.head
	if l.Length() == 1 {
		l.head = nil
	} else {
		for i := 1; i < l.Length()-1; i++ {
			node = node.Next()
		}
		node.Set(node.Next().Next())
	}
	l.length = l.length - 1
}

// for printing the link list
func (l *linkList) Print() {
	if l.head != nil {
		//fmt.Printf("List elements: ")
		for node := l.head; node != nil; node = node.Next() {
			fmt.Printf("%+v ", node.Value())
		}
		fmt.Println()
	}
}

func main() {
	var l linkList
	var number int
	fmt.Println("Enter the number of producers you want: ")
	fmt.Scanf("%d", &number)

	fmt.Println("Now we Produce :")
	for i := 1; i < number+1; i++ {
		l.Push(i)
		l.Print()
	}

	fmt.Println("Now we Consume:")
	l.Print()
	for k := 1; k < number+1; k++ {
		l.Pop()
		l.Print()
	}
	fmt.Println("Linked List is empty ")
	//l.Print()

	//fmt.Printf("linked list took %s", elapsed)
	fmt.Println("")
	fmt.Println("")

}
