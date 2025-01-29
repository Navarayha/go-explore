package main

import (
	"fmt"
)

func main() {
	stack := Stack[int]{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	fmt.Println("Размер стека:", stack.Size()) // 3

	top, _ := stack.Pop()
	fmt.Println("Pop:", top) // 30

	peek, _ := stack.Peek()
	fmt.Println("Peek:", peek) // 20

	fmt.Println("Стек пуст?", stack.IsEmpty()) // false

	queue := Queue[int]{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	fmt.Println("Размер очереди:", queue.Size()) // 3

	front, _ := queue.Front()
	fmt.Println("Front:", front) // 10

	dequeued, _ := queue.Dequeue()
	fmt.Println("Dequeue:", dequeued) // 10

	fmt.Println("Размер очереди после Dequeue:", queue.Size()) // 2
}
