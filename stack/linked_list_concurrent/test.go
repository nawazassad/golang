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
	node := l.Head
	if l.Length == 1 {
		l.Head = nil
	} else {
		for i := 1; i < l.Length-1; i++ {
			node = node.Next
		}
    node.Next = node.Next.Next
	}
	l.Length = l.Length - 1
}

func producers(l *LinkedList, c1, c2 chan string, number int){
  for i:=1; i< number+1 ;i++{
    
    l.Push(i)
    c1 <- "produced"
    <-c2
  }
  close(c1)
}

func consumer(l *LinkedList, c1, c2 chan string){
  for _ = range c1{
    l.Pop()
    c2 <- "consumed"
  }
  close(c2)
}

func main() {
	var l LinkedList
  var number int
  var c1 = make(chan string, 10)
  var c2 = make(chan string, 10)

	fmt.Println("Enter the number of producers you want: ")
	fmt.Scanf("%d", &number)

  start := time.Now()

  go producers(&l, c1, c2, number)
  go consumer(&l, c1, c2)

  elapsed := time.Since(start)
  fmt.Println("Time taken is: ", elapsed)
}
