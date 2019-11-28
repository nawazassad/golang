package main

import (
  "fmt"
)

type node struct{
  next *node
  value interface{}
}

type circularBuffer struct{
  length int
  head *node
  tail *node
}

func (n *node) Value() interface{} {
  return n.value
}

func (n *node) Next() *node{
  return n.next
}

func (l *circularBuffer) Set(new_node *node){
  head := l.head
  tail := l.tail

  new_node.next = head
  tail.next     = new_node

}

func (l *circularBuffer) Length() int{
  return l.length
}

func (l *circularBuffer) Push(val interface{}){
  n := &node{value: val}

  if l.head == nil {
    l.head = n
  }else{
    //l.head.Set(n)
    l.Set(n)
  }
  l.tail = n
  l.length = l.length + 1
  //fmt.Println("pushing--->", val)
}

func (l *circularBuffer) Pop(){
  if l.length ==0{
    return
  }

  l.head = l.head.next
  l.tail.next = l.head
  l.length = l.length - 1
  l.Print()
}

func (l *circularBuffer) Pop1(){
  node := l.head
  if l.Length() == 1 {
    l.head = nil
  } else {
    for i:=1; i < l.Length()-1; i++{
      node = node.Next()
    }
    node.next = l.head
  }
  l.length = l.length - 1
}

func (l *circularBuffer) Print() {
  if l.length ==0{
    return
  }
  node := l.head
  fmt.Println("length-->", l.length)
  //fmt.Println("first_value", node.value)
  for i:=(l.length-1); i!=0; i--{
    fmt.Printf("%+v ", node.value)
    node = node.next
  }
  fmt.Println()
}

func main(){
  var l circularBuffer
  var number int
  fmt.Println("Enter the number of producers you want: ")
  fmt.Scanf("%d", &number)

  fmt.Println("Now we Produce :")
  for i := 1; i < number+1; i++ {
    l.Push(i)

  }

  fmt.Println("Length of the list: ", l.Length())
  l.Print()

  fmt.Println("Now we Consume:")

  for k := 1; k < number+1; k++ {
    l.Pop()

  }
    l.Print()

  fmt.Println("Linked List is empty ")



}

