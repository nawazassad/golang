package main

import (
	"fmt"
  "time"
  "sync"
)

type Node struct {
	Next  *Node
	Value interface{}
}

//Creating Struct For link list
type LinkedList struct {
  sync.Mutex
	Length int // will be used for counting the Length
	Head   *Node
	Tail   *Node
}

func (l *LinkedList)Operate(tail *Node, node *Node, state bool){
  l.Lock()

  if stat{
    tail.Next = n
    l.Tail = n
    l.Length = l.Length + 1
  }else{
		for i := 1; i < l.Length-1; i++ {
      fmt.Println(node.Value)
			node = node.Next
		}
    node.Next = node.Next.Next
	  l.Length = l.Length - 1
  }

  l.Unlock()
}

// to push a new node  at the end of the link list
func (l *LinkedList) Push(val interface{}) {
	n := &Node{Value: val}
  l.Lock()
	if l.Head == nil {
		l.Head = n
	} else {
    l.Operate(l.Tail, n, true)
		//l.Tail.Next = n
	}
}

// this will pop the last element in the link list
func (l *LinkedList) Pop() {
  l.Lock()
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
  l.Unlock()
}


func producers(l *LinkedList, c chan int){
  for i:=1; i<l.Length+1;i++{
    l.Push(i)
    fmt.Println("pushing-->", i)
    c <- i
  }
  close(c)
}

func consumers(l *LinkedList){
  l.Pop()
}

func main() {
	var l LinkedList
  var c = make(chan int)

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
