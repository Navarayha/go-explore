package main

import (
	"errors"
)

// Stack Обобщённый стек (Generics)
type Stack[T any] struct {
	items []T
}

// Push Добавить элемент в стек
func (s *Stack[T]) Push(value T) {
	s.items = append(s.items, value)
}

// Pop Удалить верхний элемент и вернуть его
func (s *Stack[T]) Pop() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("стек пуст")
	}
	top := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return top, nil
}

// Peek Посмотреть верхний элемент без удаления
func (s *Stack[T]) Peek() (T, error) {
	if s.IsEmpty() {
		var zero T
		return zero, errors.New("стек пуст")
	}
	return s.items[len(s.items)-1], nil
}

// IsEmpty Проверить, пуст ли стек
func (s *Stack[T]) IsEmpty() bool {
	return len(s.items) == 0
}

// Size Получить размер стека
func (s *Stack[T]) Size() int {
	return len(s.items)
}
