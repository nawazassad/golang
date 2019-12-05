package main

import (
	"fmt"
	"sync"
	"time"
)

type Node struct {
	Next  *Node
	Value interface{}
}

type circularBuffer struct {
	Length int
	Head   *Node
	Tail   *Node
	sync.Mutex
}

func (l *circularBuffer) Set(new_node *Node) {
	head := l.Head
	tail := l.Tail

	new_node.Next = head
	tail.Next = new_node

}

func (l *circularBuffer) Push(val interface{}, wg *sync.WaitGroup) {
	//var wg sync.WaitGroup
	defer wg.Done()
	l.Lock()
	defer l.Unlock()
	n := &Node{Value: val}

	if l.Head == nil {
		l.Head = n
	} else {
		l.Set(n)
		// l.Unlock()
	}

	l.Tail = n
	l.Length = l.Length + 1
	fmt.Printf("%+v ", n.Value)
	//fmt.Print("\n", n.Value)
	//wg.Done()
}

func (l *circularBuffer) Pop(wg *sync.WaitGroup) {
	//var wg sync.WaitGroup
	defer wg.Done()
	l.Lock()
	defer l.Unlock()

	if l.Length == 0 {
		return
	}

	l.Head = l.Head.Next
	l.Tail.Next = l.Head
	l.Length = l.Length - 1

	//l.Print()
	// for node := l.Head; node != nil; node = node.Next {
	// 	fmt.Printf("%+v ", node.Value)
	// }
}

func (l *circularBuffer) Print() {
	if l.Head != nil {
		//fmt.Printf("List elements: ")
		for node := l.Head; node != nil; node = node.Next {
			fmt.Printf("%+v ", node.Value)
		}
		fmt.Println()
	}
}

// func Producers(cb *circularBuffer, number int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	cb.Lock()
// 	defer cb.Unlock()
// 	for i := 0; i < number+1; i++ {
// 		fmt.Println("pushing", i)
// 		cb.Push(i)
// 	}

// }

// func Consumers(cb *circularBuffer, number int) {
// 	for i := 0; i < number+1; i++ {

// 		cb.Pop()
// 	}
// }

func main() {
	var l circularBuffer
	var number int
	var wg sync.WaitGroup
	fmt.Println("Enter the number of producers you want: ")
	fmt.Scanf("%d", &number)

	start := time.Now()
	wg.Add(number)

	fmt.Println("Now we push:")
	//go Producers(&l, number, &wg)
	for i := 1; i <= number; i++ {
		//fmt.Println(i)
		go l.Push(i, &wg)
		//l.Print()
	}
	wg.Wait()

	wg.Add(number)
	fmt.Println("Now we Consume:")
	//l.Print()
	for k := 1; k <= number; k++ {
		// for checking if there is race condition we use below go l.Pop()
		//go l.Pop()
		// fmt.Println(k)
		go l.Pop(&wg)

		//l.Print()
	}
	// fmt.Println(l.Head.Value)
	// wg.Wait()
	// wg.Add(number)
	//go consumers(&l, number)
	//wg.Wait()
	//l.Print()
	elapsed := time.Since(start)
	fmt.Println("Time taken is: ", elapsed)

}
