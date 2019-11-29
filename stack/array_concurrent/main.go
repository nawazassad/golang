package main

import (
	"fmt"
  "time"
)

type Stack struct {
	Data []interface{}
  Size int
}

func (s *Stack) Len() int {
	return len(s.Data)
}

func (s *Stack) Push(value interface{}) {
	s.Data = append(s.Data, value)
}

func (s *Stack) Pop() interface{} {
	if s.Len() > 0 {
		rect := s.Data[s.Len()-1]
		s.Data = s.Data[:s.Len()-1]
		return rect
	}
	return nil
}

func producer(s *Stack, c1, c2 chan string){

  for i:=1; i<s.Size+1; i++{
    s.Push(i)
    fmt.Println("pushing data-->", i)
    c1 <- "produced"
    <-c2
  }
  close(c1)
}

func consumer(s *Stack, c1, c2 chan string){
  for _ = range c1{
    s.Pop()
    c2 <- "consumed"
  }
}

func main() {
	var s Stack
  var c1 = make(chan string)
  var c2 = make(chan string)

	fmt.Println("Enter the number of producers you want: ")
	fmt.Scanf("%d", &s.Size)

	fmt.Println("Now we Produce :")
  start := time.Now()

  go producer(&s, c1, c2)
  go consumer(&s, c1, c2)
/*
  for msg := range c{
    fmt.Println("Popping", msg)
    consumer(&s)
  }
*/

  elapsed := time.Since(start)
  fmt.Println("Time taken is: ", elapsed)
}

