package main

import (
	"testing"
)

// Тест Push и Size
func TestStackPush(t *testing.T) {
	stack := Stack[int]{}
	stack.Push(10)
	stack.Push(20)
	stack.Push(30)

	if stack.Size() != 3 {
		t.Errorf("Ожидалось 3, получено %d", stack.Size())
	}
}

// Тест Pop
func TestStackPop(t *testing.T) {
	stack := Stack[int]{}
	stack.Push(10)
	stack.Push(20)

	val, err := stack.Pop()
	if err != nil || val != 20 {
		t.Errorf("Pop() вернул %d, ожидалось 20", val)
	}

	if stack.Size() != 1 {
		t.Errorf("Ожидался размер 1, получено %d", stack.Size())
	}
}

// Тест Pop на пустом стеке
func TestStackPopEmpty(t *testing.T) {
	stack := Stack[int]{}

	_, err := stack.Pop()
	if err == nil {
		t.Errorf("Ожидалась ошибка при Pop() из пустого стека")
	}
}

// Тест Peek
func TestStackPeek(t *testing.T) {
	stack := Stack[int]{}
	stack.Push(5)
	stack.Push(15)

	val, err := stack.Peek()
	if err != nil || val != 15 {
		t.Errorf("Peek() вернул %d, ожидалось 15", val)
	}

	if stack.Size() != 2 {
		t.Errorf("После Peek() размер должен остаться 2, получено %d", stack.Size())
	}
}

// Тест IsEmpty
func TestStackIsEmpty(t *testing.T) {
	stack := Stack[int]{}
	if !stack.IsEmpty() {
		t.Errorf("Ожидалось, что стек будет пустым")
	}

	stack.Push(1)
	if stack.IsEmpty() {
		t.Errorf("Ожидалось, что стек НЕ будет пустым")
	}
}
