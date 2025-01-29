package main

import (
	"flag"
	"fmt"
	"math"
	"strings"
)

// Point представляет точку с координатами (долгота, широта)
type Point struct {
	X, Y float64
}

// DistanceTo вычисляет расстояние между двумя точками
func (p Point) DistanceTo(other Point) float64 {
	dx := p.X - other.X
	dy := p.Y - other.Y
	return math.Sqrt(dx*dx + dy*dy)
}

// WithinRadius проверяет, находится ли точка в радиусе N относительно другой точки
func (p Point) WithinRadius(center Point, radius float64) bool {
	return p.DistanceTo(center) <= radius
}

// Polygon представляет замкнутый многоугольник (список точек)
type Polygon struct {
	Points []Point
}

// Contains проверяет, находится ли точка внутри полигона (метод луча)
func (poly Polygon) Contains(pt Point) bool {
	n := len(poly.Points)
	if n < 3 {
		return false
	}
	inside := false
	j := n - 1
	for i := 0; i < n; i++ {
		xi, yi := poly.Points[i].X, poly.Points[i].Y
		xj, yj := poly.Points[j].X, poly.Points[j].Y
		if (yi > pt.Y) != (yj > pt.Y) &&
			(pt.X < (xj-xi)*(pt.Y-yi)/(yj-yi)+xi) {
			inside = !inside
		}
		j = i
	}
	return inside
}

// parsePoint парсит строку "X,Y" в структуру Point
func parsePoint(s string) (Point, error) {
	parts := strings.Split(s, ",")
	if len(parts) != 2 {
		return Point{}, fmt.Errorf("неверный формат точки: %s", s)
	}
	var x, y float64
	_, err := fmt.Sscanf(s, "%f,%f", &x, &y)
	if err != nil {
		return Point{}, err
	}
	return Point{X: x, Y: y}, nil
}

func main() {
	var pointFlags pointArray
	flag.Var(&pointFlags, "point", "Координаты точки (пример: 1.0,2.0)")

	distanceFlag := flag.Bool("distance", false, "Вычислить расстояние между первыми двумя точками")

	flag.Parse()

	if len(pointFlags) < 2 {
		fmt.Println("Необходимо передать хотя бы две точки через --point=X,Y")
		return
	}

	p1, p2 := pointFlags[0], pointFlags[1]

	if *distanceFlag {
		fmt.Printf("Расстояние: %.2f\n", p1.DistanceTo(p2))
	}
}

// pointArray — кастомный тип для обработки множества --point флагов
type pointArray []Point

func (p *pointArray) String() string {
	var points []string
	for _, pt := range *p {
		points = append(points, fmt.Sprintf("%.2f,%.2f", pt.X, pt.Y))
	}
	return strings.Join(points, " ")
}

func (p *pointArray) Set(value string) error {
	pt, err := parsePoint(value)
	if err != nil {
		return err
	}
	*p = append(*p, pt)
	return nil
}
