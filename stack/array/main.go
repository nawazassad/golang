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

func producer(s *Stack){

  for i:=1; i<s.Size+1; i++{
    s.Push(i)
  }
}

func consumer(s *Stack){
  for i:=1;i<s.Size+1;i++{
    s.Pop()
  }
}

func main() {
	var s Stack

	fmt.Println("Enter the number of producers you want: ")
	fmt.Scanf("%d", &s.Size)

	fmt.Println("Now we Produce :")
  start := time.Now()

  producer(&s)
  consumer(&s)

  elapsed := time.Since(start)
  fmt.Println("Time taken is: ", elapsed)
}

