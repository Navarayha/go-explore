package main

import (
	"errors"
)

// Queue Очередь (Queue)
type Queue[T any] struct {
	items []T
}

// Enqueue Добавить элемент в очередь
func (q *Queue[T]) Enqueue(value T) {
	q.items = append(q.items, value)
}

// Dequeue Удалить элемент из очереди и вернуть его
func (q *Queue[T]) Dequeue() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("очередь пуста")
	}
	front := q.items[0]
	q.items = q.items[1:]
	return front, nil
}

// Front Посмотреть элемент в начале очереди без удаления
func (q *Queue[T]) Front() (T, error) {
	if q.IsEmpty() {
		var zero T
		return zero, errors.New("очередь пуста")
	}
	return q.items[0], nil
}

// IsEmpty Проверить, пуста ли очередь
func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

// Size Получить размер очереди
func (q *Queue[T]) Size() int {
	return len(q.items)
}
