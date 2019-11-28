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
}

func (c *circularArray) pop(index int) {
	c.queue[index] = 0
}

func producer(obj *circularArray, c chan int) {
	for i := 0; i < obj.size; i++ {
		obj.push(i)
		c <- (i)
	}
	close(c)
}

func consume(obj *circularArray, i int) {
	obj.pop(i)
}

func main() {
	obj := size(100000000)
	var c = make(chan int, 10)
	start := time.Now()
	go producer(&obj, c)

	for msg := range c {
		//fmt.Println("Popping---->", msg)
		consume(&obj, msg)
	}
	elapsed := time.Since(start)

	fmt.Println("Time taken is: ", elapsed)
}
