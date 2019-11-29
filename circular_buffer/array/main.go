package main

import (
	"fmt"
	"time"
)

type CircularArray struct {
	Head  int
	Tail  int
	Size  int
	Queue []int
}

func Size(z int) CircularArray {
	fix_Size := make([]int, z)
	return CircularArray{-1, -1, z, fix_Size}
}

func (c *CircularArray) Push(value int) {
	c.Tail = (c.Tail + 1) % c.Size
	c.Queue[c.Tail] = value
}

func (c *CircularArray) Pop() {
	if c.Head == c.Tail {
		return
	}
	c.Head = (c.Head + 1) % c.Size
	c.Queue[c.Head] = 0
}


func producers(ca *CircularArray, number int){
  for i:=1; i<number+1; i++{
    ca.Push(i)
  }
}

func consumers(ca *CircularArray, number int){
  for i:=1; i<number+1; i++{
    ca.Push(i)
  }
}

func main() {
	var number int
	fmt.Println("Enter the number of producers you want: ")
	fmt.Scanf("%d", &number)

	obj := Size(number)
	start := time.Now()

  producers(&obj, number)
  consumers(&obj, number)

	elapsed := time.Since(start)

	fmt.Println("Time taken is: ", elapsed)
}
