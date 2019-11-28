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

func size(z int) circularArray {
	fix_size := make([]int, z)
	return circularArray{-1, -1, z, fix_size}
}

func (c *circularArray) push(value int) {
	c.head = (c.head + 1) % c.size
	c.queue[c.head] = value
	//fmt.Println("moving push index to-->", c.head)
}

func (c *circularArray) pop(index int) {
	c.queue[index] = 0
}

func producer(obj *circularArray, c chan int) {
	for i := 0; i < obj.size; i++ {
		fmt.Println("Pushing--->", i+1)
		obj.push(i + 1)
		c <- (i + 1)
	}
	close(c)
}

func consume(obj *circularArray, i int) {
	obj.pop(i - 1)
}

func main() {
	obj := size(1000000)
	var c = make(chan int)
	start := time.Now()
	go producer(&obj, c)

	for msg := range c {
		fmt.Println("Popping---->", msg)
		consume(&obj, msg)
	}
	elapsed := time.Since(start)

	fmt.Println("Time taken is: ", elapsed)
}
