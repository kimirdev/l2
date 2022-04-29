package main

import "fmt"

/*
	Команда — это поведенческий паттерн проектирования,
	который превращает запросы в объекты, позволяя передавать их как аргументы
	при вызове методов, ставить запросы в очередь, логировать их, а также поддерживать отмену операций.

	Применять:
	Когда вы хотите параметризовать объекты выполняемым действием.
	Когда вы хотите ставить операции в очередь, выполнять их по расписанию или передавать по сети.
	Когда вам нужна операция отмены.

	Плюсы:
	Убирает прямую зависимость между объектами, вызывающими операции, и объектами, которые их непосредственно выполняют.
	Позволяет реализовать простую отмену и повтор операций.
	Позволяет реализовать отложенный запуск операций.
	Позволяет собирать сложные команды из простых.
	Реализует принцип открытости/закрытости.
*/

// Отправитель
type button struct {
	command command
}

func (b *button) press() {
	b.command.execute()
}

// Интерфейс команды
type command interface {
	execute()
}

// Конкретная команда
type onCommand struct {
	device device
}

func (c *onCommand) execute() {
	c.device.on()
}

// Конкретная команда
type offCommand struct {
	device device
}

func (c *offCommand) execute() {
	c.device.off()
}

// Интерфейс получателя
type device interface {
	on()
	off()
}

// Конкретный получатель
type tv struct {
	isRunning bool
}

func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turning tv on")
}

func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}

func main() {
	tv := &tv{}

	onCommand := &onCommand{
		device: tv,
	}

	offCommand := &offCommand{
		device: tv,
	}

	onButton := &button{
		command: onCommand,
	}
	onButton.press()

	offButton := &button{
		command: offCommand,
	}
	offButton.press()
}
