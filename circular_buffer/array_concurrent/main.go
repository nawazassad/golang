package main

import (
	"fmt"
	"time"
)

type circularArray struct {
	head  int
	tail  int
	size  int
	queue []int
}

func (c *circularArray) Len() int {
  return len(c.queue)
}

func (c *circularArray) push(value int) {
	c.head = (c.head + 1) % c.size
  fmt.Println("head--->", c.head)
  fmt.Println("Queue--->", c.queue)
	c.queue[c.head] = value
  c.queue = append(c.queue, value)
}

func (c *circularArray) pop() {
  if c.Len() >0{
    c.queue = c.queue[:c.Len()-1]
  }
  return
}

func producer(obj *circularArray, c1, c2 chan string) {
	for i := 0; i < obj.size; i++ {
		obj.push(i)
		c1 <- "produced"
    <- c2
	}
	close(c1)
}

func consumer(obj *circularArray, c1, c2, c3 chan string) {
  for _ = range c1{
    obj.pop()
    c2 <- "consumed"
  }
  close(c2)
  c3 <- "completed"
}

func main() {
  var number int
  obj := circularArray{head: -1, tail: -1}
  fmt.Println("Enter the number of producers you want: ")
  fmt.Scanf("%d", &number)


	var c1 = make(chan string)
	var c2 = make(chan string)
	var c3 = make(chan string)

	start := time.Now()
	go producer(&obj, c1, c2)
	go consumer(&obj, c1, c2, c3)

  <- c3
  close(c3)
	elapsed := time.Since(start)
	fmt.Println("Time taken is: ", elapsed)
}
