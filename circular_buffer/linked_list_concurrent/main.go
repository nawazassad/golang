package main

import (
  "fmt"
  "time"
)

type Node struct{
  Next *Node
  Value interface{}
}

type CircularBuffer struct{
  Length int
  Head *Node
  Tail *Node
}

func (l *CircularBuffer) Set(new_node *Node){
  head := l.Head
  tail := l.Tail

  new_node.Next = head
  tail.Next     = new_node

}

func (l *CircularBuffer) Push(val interface{}){
  n := &Node{Value: val}

  if l.Head == nil {
    l.Head = n
  }else{
    l.Set(n)
  }
  l.Tail = n
  l.Length = l.Length + 1
}

func (l *CircularBuffer) Pop(){
  if l.Length ==0{
    return
  }

  l.Head = l.Head.Next
  l.Tail.Next = l.Head
  l.Length = l.Length - 1
}

func producers(cb *CircularBuffer, c1, c2 chan string, number int){
  for i:=0; i<number+1; i++{
    cb.Push(i)
    c1 <- "produced"
    <- c2
  }
  close(c1)
}

func consumers(cb *CircularBuffer, c1, c2, c3 chan string){
  for _ = range c1{
    cb.Pop()
    c2 <- "consumed"
  }
  close(c2)
  c3 <- "completed"
}

func consumers1(cb *CircularBuffer, number int){
  for i:=0; i<number+1; i++{
    cb.Pop()
  }
}


func main(){
  var l CircularBuffer
  var number int
  var c1 = make(chan string)
  var c2 = make(chan string)
  var c3 = make(chan string)

  fmt.Println("Enter the number of producers you want: ")
  fmt.Scanf("%d", &number)

  start := time.Now()

  go producers(&l, c1, c2, number)
  go consumers(&l, c1, c2, c3)
  <-c3

  elapsed := time.Since(start)
  fmt.Println("Time taken is: ", elapsed)



}

