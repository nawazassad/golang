package main

import (
	"fmt"
	"sync"
	"time"
)

var mutex sync.Mutex

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

func (c *circularArray) push(value int, wg *sync.WaitGroup) {
	defer wg.Done()
	mutex.Lock()
	defer mutex.Unlock()
	c.tail = (c.tail + 1) % c.size
	c.queue[c.tail] = value
	fmt.Println("moving push index to-->", c.tail)
	//fmt.Println(c.queue)
}

func (c *circularArray) pop(wg *sync.WaitGroup) {
	//defer wg.Done()
	mutex.Lock()
	defer mutex.Unlock()
	if c.head == c.tail {
		fmt.Println("Nothing to pop at this time")
		wg.Done()
		return
	}
	c.head = (c.head + 1) % c.size
	c.queue[c.head] = 0
	fmt.Println("moving pop index to-->", c.head)
	//fmt.Println(c.queue)
	wg.Done()
}

func main() {
	var number int
	var wg sync.WaitGroup
	fmt.Println("Enter the number of producers you want: ")
	fmt.Scanf("%d", &number)
	obj := size(number)
	wg.Add(number)
	start := time.Now()
	for i := 1; i <= number; i++ {
		// for checking if there is race condition we use below go l.Push()
		//go l.Push(i)
		go obj.push(i, &wg)
		//fmt.Println(obj)
		//l.Print()
	}
	wg.Wait()

	wg.Add(number)
	for k := 1; k <= number; k++ {
		// for checking if there is race condition we use below go l.Pop()
		//go l.Pop()
		go obj.pop(&wg)
		//fmt.Println(obj)
		//l.Print()
	}
	wg.Wait()

	elapsed := time.Since(start)

	fmt.Println("Time taken is: ", elapsed)
	// obj.push(1)
	// obj.push(2)
	// obj.push(3)

	// fmt.Println(obj)
	// obj.pop()
	// obj.pop()
	// fmt.Println(obj)
	// obj.pop()
	// obj.pop()
	// fmt.Println(obj)
}