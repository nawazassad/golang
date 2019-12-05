package main

import (
	"fmt"
	"sync"
)

type stack struct {
	sync.Mutex
	data []interface{}
}

//return the number of items in stack
func (s *stack) Len() int {
	return len(s.data)
}

func (s *stack) Push(value interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	s.Lock()
	defer s.Unlock()
	s.data = append(s.data, value)
	fmt.Println(s.data)
}

//pop the top item out
func (s *stack) Pop(wg *sync.WaitGroup) interface{} {
	defer wg.Done()
	s.Lock()
	defer s.Unlock()
	if s.Len() > 0 {
		s.data = s.data[:s.Len()-1]
		fmt.Println(s.data)
	}
	return nil
}

func main() {
	var l stack
	var number int
	var wg sync.WaitGroup

	fmt.Println("Enter the number of producers you want: ")
	fmt.Scanf("%d", &number)

	wg.Add(number)
	fmt.Println("Now we Produce :")
	for i := 1; i <= number; i++ {
		go l.Push(i, &wg)
	}
	wg.Wait()

	wg.Add(number)
	for k := 1; k <= number; k++ {
		go l.Pop(&wg)
		//fmt.Println(l)
		//l.Print()
	}
	wg.Wait()
}
