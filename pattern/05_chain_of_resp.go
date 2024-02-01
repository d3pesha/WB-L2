package main

import "fmt"

type Request struct {
	content string
}

type Handler interface {
	handleRequest(request Request)
	setNext(handler Handler)
}

type ConcreteHandler struct {
	nextHandler Handler
	name        string
}

func (c *ConcreteHandler) handleRequest(request Request) {
	fmt.Printf("%s is handling the request: %s\n", c.name, request.content)

	if c.nextHandler != nil {
		c.nextHandler.handleRequest(request)
	}
}

func (c *ConcreteHandler) setNext(handler Handler) {
	c.nextHandler = handler
}

func main() {
	// Создаем обработчики
	manager := &ConcreteHandler{name: "Manager"}
	supervisor := &ConcreteHandler{name: "Supervisor"}
	hr := &ConcreteHandler{name: "HR"}

	// Устанавливаем цепочку обработчиков
	manager.setNext(supervisor)
	supervisor.setNext(hr)

	// Создаем запрос
	request := Request{content: "Need a day off"}

	// Начинаем обработку запроса с первого обработчика в цепочке
	manager.handleRequest(request)
}

/*
Описание:
Цепочка обязанностей — это поведенческий паттерн проектирования,
который позволяет передавать запросы последовательно по цепочке обработчиков.
Каждый последующий обработчик решает, может ли он обработать запрос сам
и стоит ли передавать запрос дальше по цепи.
Преимущества:
1. Уменьшает зависимость между клиентом и обработчиками.
2. Реализует принцип единственной обязанности.
3. Реализует принцип открытости/закрытости.
Недостатки:
1. Запрос может остаться никем не обработанным.

Цепочка ответственностей может быть как линейной, так и разветвленной.
*/
