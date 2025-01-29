package main

import (
	"math"
	"testing"
)

// Тест DistanceTo
func TestDistanceTo(t *testing.T) {
	p1 := Point{X: 0, Y: 0}
	p2 := Point{X: 3, Y: 4} // Должно быть 5 (теорема Пифагора)

	dist := p1.DistanceTo(p2)
	expected := 5.0

	if math.Abs(dist-expected) > 1e-6 {
		t.Errorf("Ожидалось %.2f, получено %.2f", expected, dist)
	}
}

// Тест WithinRadius
func TestWithinRadius(t *testing.T) {
	center := Point{X: 0, Y: 0}
	target := Point{X: 3, Y: 4} // Расстояние = 5

	if !target.WithinRadius(center, 5) {
		t.Errorf("Точка должна быть в радиусе 5")
	}

	if target.WithinRadius(center, 4.9) {
		t.Errorf("Точка не должна быть в радиусе 4.9")
	}
}

// Тест Contains для полигона (треугольник)
func TestPolygonContains(t *testing.T) {
	poly := Polygon{
		Points: []Point{
			{X: 0, Y: 0},
			{X: 5, Y: 0},
			{X: 2.5, Y: 5},
		},
	}

	tests := []struct {
		pt       Point
		expected bool
	}{
		{Point{X: 2.5, Y: 2}, true},  // Внутри
		{Point{X: 5, Y: 5}, false},   // Вне
		{Point{X: 2.5, Y: 0}, true},  // На ребре
		{Point{X: 0, Y: 0}, true},    // Вершина
		{Point{X: -1, Y: -1}, false}, // Вне
	}

	for _, test := range tests {
		result := poly.Contains(test.pt)
		if result != test.expected {
			t.Errorf("Ошибка: точка %+v, ожидалось %v, получено %v", test.pt, test.expected, result)
		}
	}
}
