package main

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type Car struct {
	Model      string
	Color      string
	Horsepower int
}

type CarBuilderInterface interface {
	SetModel(model string) CarBuilderInterface
	SetColor(color string) CarBuilderInterface
	SetHorsepower(hp int) CarBuilderInterface
	Build() Car
}

type CarBuilder struct {
	model      string
	color      string
	horsepower int
}

func (cb *CarBuilder) SetModel(model string) CarBuilderInterface {
	cb.model = model
	return cb
}

func (cb *CarBuilder) SetColor(color string) CarBuilderInterface {
	cb.color = color
	return cb
}

func (cb *CarBuilder) SetHorsepower(hp int) CarBuilderInterface {
	cb.horsepower = hp
	return cb
}

func (cb *CarBuilder) Build() Car {
	return Car{
		Model:      cb.model,
		Color:      cb.color,
		Horsepower: cb.horsepower,
	}
}

func main() {
	builder := &CarBuilder{}
	car := builder.SetModel("Sedan").SetColor("Blue").SetHorsepower(2000).Build()

	fmt.Printf("Model: %s\nColor: %s\nHorsepower: %d\n", car.Model, car.Color, car.Horsepower)
}

/* Преимущества и недостатки
+ Позволяет создавать продукты пошагово.
+ Позволяет использовать один и тот же код для создания различных продуктов.
+ Изолирует сложный код сборки продукта от его основной бизнес-логики.
-Усложняет код программы из-за введения дополнительных классов
Паттерн Строитель также используется, когда нужный продукт сложный и требует нескольких шагов для построения.
В таких случаях несколько конструкторных методов подойдут лучше, чем один громадный конструктор.
При использовании пошагового построения объектов потенциальной проблемой является выдача клиенту частично построенного
нестабильного продукта. Паттерн "Строитель" скрывает объект до тех пор, пока он не построен до конца.
В этом примере мы можем создать Франкенштейна с разными параметрами */
