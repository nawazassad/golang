package main

import (
  "fmt"
  "time"
)

type Node struct{
  Next *Node
  Value interface{}
}

type circularBuffer struct{
  Length int
  Head *Node
  Tail *Node
}


func (l *circularBuffer) Set(new_node *Node){
  head := l.Head
  tail := l.Tail

  new_node.Next = head
  tail.Next     = new_node

}

func (l *circularBuffer) Push(val interface{}){
  n := &Node{Value: val}

  if l.Head == nil {
    l.Head = n
  }else{
    l.Set(n)
  }
  l.Tail = n
  l.Length = l.Length + 1
}

func (l *circularBuffer) Pop(){
  if l.Length ==0{
    return
  }

  l.Head = l.Head.Next
  l.Tail.Next = l.Head
  l.Length = l.Length - 1
}

func producers(cb *circularBuffer, number int){
  for i:=0; i<number+1; i++{
    cb.Push(i)
  }
}

func consumers(cb *circularBuffer, number int){
  for i:=0; i<number+1; i++{
    cb.Pop()
  }
}


func main(){
  var l circularBuffer
  var number int
  fmt.Println("Enter the number of producers you want: ")
  fmt.Scanf("%d", &number)

  start := time.Now()

  producers(&l, number)
  consumers(&l, number)

  elapsed := time.Since(start)
  fmt.Println("Time taken is: ", elapsed)



}

