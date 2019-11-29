package main

import (
	"fmt"
	"sync"
)

// Creating Struct for Node
type node struct {
	next  *node
	value interface{}
}

//Creating Struct For link list
type linkList struct {
	sync.Mutex
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
func (l *linkList) Push(val interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	l.Lock()
	defer l.Unlock()
	n := &node{value: val}

	if l.head == nil {
		l.head = n
	} else {
		l.tail.Set(n)
	}
	l.tail = n
	l.length = l.length + 1
	//fmt.Printf("%+v ", l.Value())
	l.Print()
	//l.Unlock()
}

// this will pop the last element in the link list
func (l *linkList) Pop(wg *sync.WaitGroup) {
	defer wg.Done()
	l.Lock()
	defer l.Unlock()
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
	l.Print()
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
	var wg sync.WaitGroup
	var number int
	fmt.Println("Enter the number of producers you want: ")
	fmt.Scanf("%d", &number)
	wg.Add(number)
	fmt.Println("Now we Produce :")
	for i := 1; i < number+1; i++ {
		// for checking if there is race condition we use below go l.Push()
		//go l.Push(i)
		go l.Push(i, &wg)
		//l.Print()
	}
	wg.Wait()

	wg.Add(number)
	fmt.Println("Now we Consume:")
	l.Print()
	for k := 1; k < number+1; k++ {
		// for checking if there is race condition we use below go l.Pop()
		//go l.Pop()
		go l.Pop(&wg)
		//l.Print()
	}
	wg.Wait()
	fmt.Println("Linked List is empty ")
	//l.Print()

	//fmt.Printf("linked list took %s", elapsed)
	fmt.Println("")
	fmt.Println("")

}
