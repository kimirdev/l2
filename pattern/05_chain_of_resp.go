package main

import "fmt"

/*
	Цепочка обязанностей — это поведенческий паттерн проектирования,
	который позволяет передавать запросы последовательно по цепочке обработчиков.
	Каждый последующий обработчик решает, может ли он обработать запрос сам и стоит ли передавать запрос дальше по цепи.

	Применять:
	Когда программа должна обрабатывать разнообразные запросы несколькими способами,
	но заранее неизвестно, какие конкретно запросы будут приходить и какие обработчики для них понадобятся.
	Когда важно, чтобы обработчики выполнялись один за другим в строгом порядке.
	Когда набор объектов, способных обработать запрос, должен задаваться динамически.

	Плюсы:
	Уменьшает зависимость между клиентом и обработчиками.
	Реализует принцип единственной обязанности.
	Реализует принцип открытости/закрытости.

	Минусы:
	Запрос может остаться никем не обработанным.
*/

// Интерфейс обработчика
type department interface {
	execute(*patient)
	setNext(department)
}

// Конкретный обработчик
type reception struct {
	next department
}

func (r *reception) execute(p *patient) {
	if p.registrationDone {
		fmt.Println("Patient registration already done")
		r.next.execute(p)
		return
	}
	fmt.Printf("Reception registering patient[%s]\n", p.name)
	p.registrationDone = true
	r.next.execute(p)
}

func (r *reception) setNext(next department) {
	r.next = next
}

// Конкретный обработчик
type doctor struct {
	next department
}

func (d *doctor) execute(p *patient) {
	if p.doctorCheckUpDone {
		fmt.Println("Doctor checkup already done")
		d.next.execute(p)
		return
	}
	fmt.Printf("Doctor checking patient[%s]\n", p.name)
	p.doctorCheckUpDone = true
	d.next.execute(p)
}

func (d *doctor) setNext(next department) {
	d.next = next
}

type medical struct {
	next department
}

func (m *medical) execute(p *patient) {
	if p.medicineDone {
		fmt.Printf("Medicine already given to patient[%s]\n", p.name)
		m.next.execute(p)
		return
	}
	fmt.Printf("Medical giving medicine to patient[%s]\n", p.name)
	p.medicineDone = true
	m.next.execute(p)
}

func (m *medical) setNext(next department) {
	m.next = next
}

type cashier struct {
	next department
}

func (c *cashier) execute(p *patient) {
	if p.paymentDone {
		fmt.Println("Payment Done")
	}
	fmt.Printf("Cashier getting money from patient[%s]\n", p.name)
	p.paymentDone = true
}

func (c *cashier) setNext(next department) {
	c.next = next
}

type patient struct {
	name              string
	registrationDone  bool
	doctorCheckUpDone bool
	medicineDone      bool
	paymentDone       bool
}

func (p patient) checkStatus() {
	fmt.Printf(`
	 Patient %q
registration status: %v
checkup status: %v
medicine status: %v
payment status: %v

`, p.name,
		p.registrationDone, p.doctorCheckUpDone,
		p.medicineDone, p.paymentDone)
}

func main() {

	cashier := &cashier{}

	//Set next for medical department
	medical := &medical{}
	medical.setNext(cashier)

	//Set next for doctor department
	doctor := &doctor{}
	doctor.setNext(medical)

	//Set next for reception department
	reception := &reception{}
	reception.setNext(doctor)

	patient := &patient{name: "Igor"}
	patient.checkStatus()
	//Patient visiting
	reception.execute(patient)
	patient.checkStatus()
}
