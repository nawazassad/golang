package main

import (
	"fmt"
  "time"
  "sync"
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

func (s *Stack) Pop() {
	if s.Len() > 0 {
		s.Data = s.Data[:s.Len()-1]
	}
}

func producer(s *Stack, c1 chan string, number int, wg *sync.WaitGroup){
  for i:=1; i<number+1; i++{
    s.Push(i)
    c1 <- "produced"
    wg.Done()
  }
  close(c1)
}

func consumer(s *Stack){
  s.Pop()
}

func main() {
	var s Stack
  var number int
  var c1 = make(chan string)
  //var c2 = make(chan string)

	fmt.Println("Enter the number of producers you want: ")
	fmt.Scanf("%d", &number)

  var wg sync.WaitGroup
  wg.Add(number)

	fmt.Println("Now we Produce :")
  start := time.Now()

  go producer(&s, c1, number, &wg)
  for _ =range c1{
    consumer(&s)
  }
  wg.Wait()
  elapsed := time.Since(start)
  fmt.Println("Time taken is: ", elapsed)
}

