package main

import "fmt"

type Transport interface {
	Drive()
}

type Motocycle struct{}

func (c *Motocycle) Drive() {
	fmt.Println("Driving a car")
}

type Bicycle struct{}

func (b *Bicycle) Drive() {
	fmt.Println("Riding a bicycle")
}

type TransportFactory interface {
	CreateTransport() Transport
}

type MotocycleFactory struct {
}

func (cf *MotocycleFactory) CreateTransport() Transport {
	return &Motocycle{}
}

type BicycleFactory struct{}

func (bf *BicycleFactory) CreateTransport() Transport {
	return &Bicycle{}
}

func main() {
	// Используем фабричный метод для создания машины
	motocycleFactory := &MotocycleFactory{}
	car := motocycleFactory.CreateTransport()
	car.Drive() // Вывод: Driving a car

	// Используем фабричный метод для создания велосипеда
	bicycleFactory := &BicycleFactory{}
	bicycle := bicycleFactory.CreateTransport()
	bicycle.Drive() // Вывод: Riding a bicycle
}

/*
Краткое описание:
	Фабричный метод(виртуальный конструктор) — это порождающий паттерн проектирования, который определяет общий интерфейс
	для создания объектов в супер-классе, позволяя подклассам изменять тип создаваемых объектов.
Плюсы:
	1. Избавляет класс от привязки к конкретным классам продуктов.
	2. Выделяет код производства продуктов в одно место, упрощая поддержку кода.
	3. Упрощает добавление новых продуктов в программу.
	4. Реализует принцип открытости/закрытости.
Минусы:
	1. Может привести к созданию больших параллельных иерархий классов,
		так как для каждого класса продукта надо создать свой подкласс создателя.
*/
