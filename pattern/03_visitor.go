package main

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/
// интерефейс представляющий геометрическую фигуру
type Shape interface {
	accept(visitor Visitor)
}

// интерфейс представляющий посетителя
type Visitor interface {
	visitCircle(circle *Circle)
	visitRectangle(rectangle *Rectangle)
}

// конкретная реализация интерфейса Shape - круг
type Circle struct {
	Radius float64
}

// метод принятия посетителя для круга
func (c *Circle) accept(visitor Visitor) {
	visitor.visitCircle(c)
}

// конкретная реализация Shape - прямоугольник
type Rectangle struct {
	Width, Height float64
}

func (r *Rectangle) accept(visitor Visitor) {
	visitor.visitRectangle(r)
}

// конкретная реализация интерфейса Visitor для вычисления площади
type AreaVisitor struct {
	TotalArea float64
}

func (v *AreaVisitor) visitCircle(circle *Circle) {
	v.TotalArea += 3.14 * circle.Radius * circle.Radius
}

func (v *AreaVisitor) visitRectangle(rectangle *Rectangle) {
	v.TotalArea += rectangle.Width * rectangle.Height
}

// объект структуры, содержащий коллекцию элементов
type ObjectStructure struct {
	shapes []Shape
}

// метод для добавления фигуры в коллекцию
func (o *ObjectStructure) attach(shape Shape) {
	o.shapes = append(o.shapes, shape)
}

// метод для принятия посетителя для каждой фигуры в коллекции
func (o *ObjectStructure) accept(visitor Visitor) {
	for _, shape := range o.shapes {
		shape.accept(visitor)
	}
}

func main() {
	// создание объекта структуры и добавление фигур
	objectStructure := &ObjectStructure{}
	objectStructure.attach(&Circle{Radius: 5})
	objectStructure.attach(&Rectangle{Width: 4, Height: 3})

	// создание посетителя для вычисления площади
	areaVisitor := &AreaVisitor{}
	objectStructure.accept(areaVisitor)

	fmt.Printf("Total area: %f\n", areaVisitor.TotalArea)
}

/*
	Паттерн "посетитель" позволяет нам добавить функционала к существующей структуре, не изменяя ее структуру.
Плюсы:
	1. Упрощает добавление операций, работающих со сложными структурами объектов.
	2. Объединяет родственные операции в одном классе.
	3. Посетитель может накапливать состояние при обходе структуры элементов.
Минусы:
	1. Паттерн не оправдан, если иерархия элементов часто меняется.
	2. Может привести к нарушению инкапсуляции элементов.
*/
