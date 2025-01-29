package main

import (
	"testing"
)

// Тест Enqueue и Size
func TestQueueEnqueue(t *testing.T) {
	queue := Queue[int]{}
	queue.Enqueue(10)
	queue.Enqueue(20)
	queue.Enqueue(30)

	if queue.Size() != 3 {
		t.Errorf("Ожидался размер 3, получено %d", queue.Size())
	}
}

// Тест Dequeue
func TestQueueDequeue(t *testing.T) {
	queue := Queue[int]{}
	queue.Enqueue(10)
	queue.Enqueue(20)

	val, err := queue.Dequeue()
	if err != nil {
		t.Errorf("Ошибка при Dequeue: %v", err)
	}
	if val != 10 {
		t.Errorf("Ожидалось 10, получено %d", val)
	}
	if queue.Size() != 1 {
		t.Errorf("Ожидался размер 1, получено %d", queue.Size())
	}

	// Проверка Dequeue из пустой очереди
	_, err = queue.Dequeue()
	_, err = queue.Dequeue() // Должен быть пуст
	if err == nil {
		t.Errorf("Ожидалась ошибка при Dequeue из пустой очереди")
	}
}

// Тест Front
func TestQueueFront(t *testing.T) {
	queue := Queue[int]{}
	queue.Enqueue(100)
	queue.Enqueue(200)

	val, err := queue.Front()
	if err != nil {
		t.Errorf("Ошибка при Front: %v", err)
	}
	if val != 100 {
		t.Errorf("Ожидалось 100, получено %d", val)
	}

	// Проверка Front на пустой очереди
	_, _ = queue.Dequeue()
	_, _ = queue.Dequeue()
	_, err = queue.Front()
	if err == nil {
		t.Errorf("Ожидалась ошибка при Front на пустой очереди")
	}
}

// Тест IsEmpty
func TestQueueIsEmpty(t *testing.T) {
	queue := Queue[int]{}
	if !queue.IsEmpty() {
		t.Errorf("Ожидалось true, получено false")
	}

	queue.Enqueue(50)
	if queue.IsEmpty() {
		t.Errorf("Ожидалось false, получено true")
	}

	_, _ = queue.Dequeue()
	if !queue.IsEmpty() {
		t.Errorf("Ожидалось true после Dequeue, получено false")
	}
}

// Тест работы со строками (Generic)
func TestQueueString(t *testing.T) {
	queue := Queue[string]{}
	queue.Enqueue("hello")
	queue.Enqueue("world")

	val, _ := queue.Dequeue()
	if val != "hello" {
		t.Errorf("Ожидалось 'hello', получено '%s'", val)
	}

	val, _ = queue.Front()
	if val != "world" {
		t.Errorf("Ожидалось 'world', получено '%s'", val)
	}
}
