package pattern

import "fmt"

/*
Строитель — это порождающий паттерн проектирования, который позволяет создавать сложные объекты пошагово.
Строитель даёт возможность использовать один и тот же код строительства для получения разных представлений объектов.

Строитель не позволяет посторонним объектам иметь доступ к конструируемому объекту,
пока тот не будет полностью готов. Это предохраняет клиентский код от получения незаконченных «битых» объектов.

Применять:
Когда вы хотите избавиться от «телескопического конструктора».
Когда ваш код должен создавать разные представления какого-то объекта.


Плюсы:
Позволяет создавать продукты пошагово.
Позволяет использовать один и тот же код для создания различных продуктов.
Изолирует сложный код сборки продукта от его основной бизнес-логики.

Минусы:
Усложняет код программы из-за введения дополнительных сущностей.
Клиент будет привязан к конкретным классам строителей, так как в интерфейсе директора может не быть метода получения результата.
*/

// Объект постройки
type house struct {
	windowType string
	doorType   string
	floor      int
}

func (h *house) printDetails() {
	fmt.Printf("Normal House Door Type: %s\n", h.doorType)
	fmt.Printf("Normal House Window Type: %s\n", h.windowType)
	fmt.Printf("Normal House Num Floor: %d\n", h.floor)
}

// Интерфейс cтроителя
type houseBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() house
}

func getBuilder(builderType string) houseBuilder {
	if builderType == "normal" {
		return &normalBuilder{}
	}

	if builderType == "igloo" {
		return &iglooBuilder{}
	}
	return nil
}

// Конкретный строитель обычного дома
type normalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newNormalBuilder() *normalBuilder {
	return &normalBuilder{}
}

func (b *normalBuilder) setWindowType() {
	b.windowType = "Wooden Window"
}

func (b *normalBuilder) setDoorType() {
	b.doorType = "Wooden Door"
}

func (b *normalBuilder) setNumFloor() {
	b.floor = 2
}

func (b *normalBuilder) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// Конкретный строитель иглу
type iglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func newIglooBuilder() *iglooBuilder {
	return &iglooBuilder{}
}

func (b *iglooBuilder) setWindowType() {
	b.windowType = "Snow Window"
}

func (b *iglooBuilder) setDoorType() {
	b.doorType = "Snow Door"
}

func (b *iglooBuilder) setNumFloor() {
	b.floor = 1
}

func (b *iglooBuilder) getHouse() house {
	return house{
		doorType:   b.doorType,
		windowType: b.windowType,
		floor:      b.floor,
	}
}

// Директор (опционально)
type director struct {
	builder houseBuilder
}

func newDirector(b houseBuilder) *director {
	return &director{
		builder: b,
	}
}

func (d *director) setBuilder(b houseBuilder) {
	d.builder = b
}

func (d *director) buildHouse() house {
	d.builder.setDoorType()
	d.builder.setWindowType()
	d.builder.setNumFloor()
	return d.builder.getHouse()
}

func main() {
	normalBuilder := getBuilder("normal")
	iglooBuilder := getBuilder("igloo")

	director := newDirector(normalBuilder)
	normalHouse := director.buildHouse()

	normalHouse.printDetails()

	director.setBuilder(iglooBuilder)
	iglooHouse := director.buildHouse()

	iglooHouse.printDetails()
}
