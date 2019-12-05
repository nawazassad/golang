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

func (s *Stack) Pop() {
	if s.Len() > 0 {
		s.Data = s.Data[:s.Len()-1]
	}
}

func producer(s *Stack, c1, c2 chan string, number int) {
  for i:=1; i<number+1; i++{
    s.Push(i)
    //fmt.Println("produced")
    c1 <- "produced"

    <-c2
  }
  close(c1)
}

/*
func consumer1(s *Stack){
  s.Pop()
}
*/

func consumer(s *Stack, c1, c2, c3 chan string){
  for _ = range c1{
    //println("received")
    //fmt.Println("consumed")
    c2 <- "consumed"
  }
  close(c2)
  c3 <- "completed"
}

func main() {
	var s Stack
  var number int
  var c1 = make(chan string)
  var c2 = make(chan string)
  var c3 = make(chan string)

	fmt.Println("Enter the number of producers you want: ")
	fmt.Scanf("%d", &number)


	fmt.Println("Now we Produce :")
  start := time.Now()

  go producer(&s, c1, c2, number)
  go consumer(&s, c1, c2, c3)

  <-c3

  elapsed := time.Since(start)
  fmt.Println("Time taken is: ", elapsed)
}

