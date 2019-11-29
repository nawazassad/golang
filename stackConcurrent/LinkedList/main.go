package main

import (
	"fmt"
	"sync"
  "time"
)

// Creating Struct for Node
type node struct {
	Next  *node
	Value interface{}
}

//Creating Struct For link list
type linkList struct {
	sync.Mutex
	Length int // will be used for counting the length
	Head   *node
	Tail   *node
}

// this will set the next node as current node , if its present
func (n *node) Set(next *node) {
	n.Next = next
}

// to push a new node  at the end of the link list
func (l *linkList) Push(val interface{}, wg *sync.WaitGroup) {
	defer wg.Done()
	l.Lock()
	defer l.Unlock()
	n := &node{Value: val}

	if l.Head == nil {
		l.Head = n
	} else {
		l.Tail.Set(n)
	}
	l.Tail = n
	l.Length = l.Length + 1
	//l.Print()
}

// this will pop the last element in the link list
func (l *linkList) Pop(wg *sync.WaitGroup) {
	defer wg.Done()
	l.Lock()
	defer l.Unlock()
	node := l.Head
	if l.Length == 1 {
		l.Head = nil
	} else {
		for i := 1; i < l.Length-1; i++ {
			node = node.Next
		}
		node.Set(node.Next.Next)
	}
	l.Length = l.Length - 1
	//l.Print()
}

// for printing the link list
func (l *linkList) Print() {
	if l.Head != nil {
		for node := l.Head; node != nil; node = node.Next {
			fmt.Printf("%+v ", node.Value)
		}
		fmt.Println()
	}
}

func main() {
	var l linkList
	var wg sync.WaitGroup
	var number int

	fmt.Println("Enter the number of producers you want: ")
	fmt.Scanf("%d", &number)

  start := time.Now()
	wg.Add(number)
	fmt.Println("Now we Produce :")
	for i := 1; i < number+1; i++ {
		go l.Push(i, &wg)
		//l.Print()
	}
	wg.Wait()

	wg.Add(number)
	fmt.Println("Now we Consume:")
	for k := 1; k < number+1; k++ {
		go l.Pop(&wg)
		//l.Print()
	}
	wg.Wait()

  elapsed := time.Since(start)
  fmt.Println("Time taken is: ", elapsed)

}
