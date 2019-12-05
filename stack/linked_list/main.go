package main

import (
	"fmt"
  "time"
)

type Node struct {
	Next  *Node
	Value interface{}
}

//Creating Struct For link list
type LinkedList struct {
	Length int // will be used for counting the Length
	Head   *Node
	Tail   *Node
}

// to push a new node  at the end of the link list
func (l *LinkedList) Push(val interface{}) {
	node := &Node{Value: val}

	if l.Head == nil {
		l.Head = node
	} else {
		l.Tail.Next = node
	}
	l.Tail = node
	l.Length = l.Length + 1
}

// this will pop the last element in the link list
func (l *LinkedList) Pop() {
  fmt.Println("popping")
	node := l.Head
	if l.Length == 1 {
		l.Head = nil
	} else {
		for i := 1; i < l.Length-1; i++ {
			node = node.Next
		}
    node.Next = nil
	}
	l.Length = l.Length - 1
  //print_ll(l)
}

func producers(l *LinkedList, number int){
  for i:=1; i< number+1 ;i++{
    l.Push(i)
  }
}

func consumers(l *LinkedList, number int){
  for i:=1;i<number+1; i++{
    l.Pop()
  }
}

func print_ll(l *LinkedList){
  for node:=l.Head;node!=nil;node=node.Next{
    fmt.Print(node.Value)
  }
}

func main() {
	var l LinkedList
  var number int

	fmt.Println("Enter the number of producers you want: ")
	fmt.Scanf("%d", &number)

  start := time.Now()

  producers(&l, number)
  consumers(&l, number)

  elapsed := time.Since(start)
  fmt.Println("Time taken is: ", elapsed)
}
