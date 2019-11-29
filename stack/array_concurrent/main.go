package main

import (
	"fmt"
  "time"
)

type stack struct {
	data []interface{}
}

func (s *stack) Len() int {
	return len(s.data)
}

func (s *stack) Push(value interface{}) {
	s.data = append(s.data, value)
	//fmt.Println(s.data)
}

func (s *stack) Pop() interface{} {
	if s.Len() > 0 {
		rect := s.data[s.Len()-1]
		s.data = s.data[:s.Len()-1]
		return rect
	}
	return nil
}

func main() {
	var l stack
	var number int

	fmt.Println("Enter the number of producers you want: ")
	fmt.Scanf("%d", &number)

	fmt.Println("Now we Produce :")
  start := time.Now()
	for i := 1; i < number+1; i++ {
		l.Push(i)
	}
	//fmt.Println("Now we Consume:")
	//fmt.Println(l)
	for k := 1; k < number+1; k++ {
		l.Pop()
		//fmt.Println(l.data)

	}
  elapsed := time.Since(start)
  fmt.Println("Time taken is: ", elapsed)
}

