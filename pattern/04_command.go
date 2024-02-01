package main

import "fmt"

// Command - интерфейс команды
type Command interface {
	Execute()
}

// CommandOn - конкретная команда для включения светильника
type CommandOn struct {
	Receiver *Light
}

func (c *CommandOn) Execute() {
	c.Receiver.TurnOn()
}

// CommandOff - конкретная команда для выключения светильника
type CommandOff struct {
	Receiver *Light
}

func (c *CommandOff) Execute() {
	c.Receiver.TurnOff()
}

// Receiver - получатель, который знает, как выполнять операции
type Light struct{}

func (l *Light) TurnOn() {
	fmt.Println("Light is ON")
}

func (l *Light) TurnOff() {
	fmt.Println("Light is OFF")
}

// Invoker - инициатор, который запрашивает выполнение команды
type RemoteControl struct {
	Command Command
}

func (r *RemoteControl) PressButton() {
	r.Command.Execute()
}

func main() {
	// Создаем объекты светильника и команд
	light := &Light{}
	commandOn := &CommandOn{Receiver: light}
	commandOff := &CommandOff{Receiver: light}

	// Создаем пульт управления
	remoteControl := &RemoteControl{}

	// Устанавливаем команду для включения и нажимаем кнопку
	remoteControl.Command = commandOn
	remoteControl.PressButton() // Вывод: Light is ON

	// Устанавливаем команду для выключения и нажимаем кнопку
	remoteControl.Command = commandOff
	remoteControl.PressButton() // Вывод: Light is OFF
}

/*
Описание:
Это поведенческий паттерн проектирования, который превращает запросы в объекты,
позволяя передавать их как аргументы при вызове методов,
ставить запросы в очередь, логировать их, а также поддерживать отмену операций.
Преимущества:
1. Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
2. Позволяет реализовать простую отмену и повтор операций.
3. Позволяет реализовать отложенный запуск операций.
4. Позволяет собирать сложные команды из простых.
5. Реализует принцип открытости/закрытости.
Недостатки:
1. Усложняет код программы из-за введения множества дополнительных классов.
*/
