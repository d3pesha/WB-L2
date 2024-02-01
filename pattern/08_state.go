package main

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

// State - интерфейс состояния
type State interface {
	Handle()
}

// TrafficLight - контекст, который использует состояния
type TrafficLight struct {
	state State
}

// SetState - метод для установки текущего состояния
func (t *TrafficLight) SetState(state State) {
	t.state = state
}

// Request - метод для обработки внешнего запроса
func (t *TrafficLight) Request() {
	t.state.Handle()
}

// RedState - конкретное состояние "Красный"
type RedState struct{}

func (r *RedState) Handle() {
	fmt.Println("Traffic Light is RED. Stop!")
}

// YellowState - конкретное состояние "Желтый"
type YellowState struct{}

func (y *YellowState) Handle() {
	fmt.Println("Traffic Light is YELLOW. Prepare to stop.")
}

// GreenState - конкретное состояние "Зеленый"
type GreenState struct{}

func (g *GreenState) Handle() {
	fmt.Println("Traffic Light is GREEN. Go!")
}

func main() {
	// Создаем объект светофора
	trafficLight := &TrafficLight{}

	// Устанавливаем начальное состояние "Красный"
	trafficLight.SetState(&RedState{})

	// Обрабатываем запрос
	trafficLight.Request() // Вывод: Traffic Light is RED. Stop!

	// Меняем состояние на "Желтый"
	trafficLight.SetState(&YellowState{})
	trafficLight.Request() // Вывод: Traffic Light is YELLOW. Prepare to stop.

	// Меняем состояние на "Зеленый"
	trafficLight.SetState(&GreenState{})
	trafficLight.Request() // Вывод: Traffic Light is GREEN. Go!
}

/*
Краткое описание:
	Состояние — это поведенческий паттерн проектирования,
	который позволяет объектам менять поведение в зависимости от своего состояния.
	Извне создаётся впечатление, что изменился класс объекта.
Плюсы:
	1. Избавляет от множества больших условных операторов машины состояний.
	2. Концентрирует в одном месте код, связанный с определённым состоянием.
	3. Упрощает код контекста.
Минусы:
	1. Может неоправданно усложнить код, если состояний мало и они редко меняются.
*/
