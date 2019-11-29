package main

import (
	"fmt"
  "time"
)

type Node struct {
	Next  *Node
	Value int
}

//Creating Struct For link list
type LinkedList struct {
	Length int // will be used for counting the Length
	Head   *Node
	Tail   *Node
}


// to push a new node  at the end of the link list
func (l *LinkedList) Push(val int, c chan int) {
	n := &Node{Value: val}
	if l.Head == nil {
		l.Head = n
    c <- val
	} else {
		l.Tail.Next = n
    c <- val
	}
	l.Tail = n
	l.Length = l.Length + 1
}

// this will pop the last element in the link list
func (l *LinkedList) Pop() {
	node := l.Head
	if l.Length == 1 {
		l.Head = nil
	} else {
		for i := 1; i < l.Length-1; i++ {
      fmt.Println(node.Value)
			node = node.Next
		}
    node.Next = node.Next.Next
	}
	l.Length = l.Length - 1
}


func producers(l *LinkedList, c chan int){
  for i:=1; i<l.Length+1;i++{
    l.Push(i, c)
    fmt.Println("pushing-->", i)
  }
  close(c)
}

func consumers(l *LinkedList){
  l.Pop()
}

func main() {
	var l LinkedList
  var c1 = make(chan int, 4)
  var c2 = make(chan int, 4)

	fmt.Println("Enter the number of producers you want: ")
	fmt.Scanf("%d", &l.Length)


  start := time.Now()

  go producers(&l, c)
  for msg := range c{
    fmt.Println("popping-->", msg)
    consumers(&l)
  }

  elapsed := time.Since(start)
  fmt.Println("Time taken is: ", elapsed)
}
