package main

import (
	"fmt"
)

type stack struct {
	data []interface{}
}

//return the number of items in stack
func (s *stack) Len() int {
	return len(s.data)
}

func (s *stack) Push(value interface{}) {
	s.data = append(s.data, value)
	fmt.Println(s.data)
}

//pop the top item out, if stack is empty, will return ErrEmptyStack decleared above
func (s *stack) Pop() interface{} {
	if s.Len() > 0 {
		rect := s.data[s.Len()-1]
		s.data = s.data[:s.Len()-1]
		return rect
		//fmt.Println(s.data)
	}
	return nil
	//return fmt.Println("Empty")
}

func main() {
	var l stack
	var number int
	fmt.Println("Enter the number of producers you want: ")
	fmt.Scanf("%d", &number)
	fmt.Println("Now we Produce :")
	for i := 1; i < number+1; i++ {
		l.Push(i)

	}
	fmt.Println("Now we Consume:")
	fmt.Println(l)
	for k := 1; k < number+1; k++ {
		l.Pop()
		fmt.Println(l)

		//l.Print()
	}
}

//-------------------------------------

// func main() {
// 	var l stack
// 	l.Push(1)
// 	l.Push(2)
// 	l.Pop()
// 	fmt.Println(l)
// }
