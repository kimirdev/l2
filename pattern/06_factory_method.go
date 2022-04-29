package pattern

import "fmt"

/*
	Фабричный метод — это порождающий паттерн проектирования,
	который определяет общий интерфейс для создания объектов в суперклассе,
	позволяя подклассам изменять тип создаваемых объектов.

	Применять:
	Когда заранее неизвестны типы и зависимости объектов, с которыми должен работать ваш код.
	Когда вы хотите дать возможность пользователям расширять части вашего фреймворка или библиотеки.
	Когда вы хотите экономить системные ресурсы, повторно используя уже созданные объекты, вместо порождения новых.

	Плюсы:
	Избавляет класс от привязки к конкретным объектам продуктов.
	Выделяет код производства продуктов в одно место, упрощая поддержку кода.
	Упрощает добавление новых продуктов в программу.
	Реализует принцип открытости/закрытости.

	Минусы:
	Может привести к созданию больших параллельных иерархий объектов,
	так как для каждого продукта надо создать своего создателя.
*/

// Интерфейс продукта
type weapon interface {
	setName(name string)
	setPower(power int)
	getName() string
	getPower() int
}

// Конкретный продукт
type gun struct {
	name  string
	power int
}

func (g *gun) setName(name string) {
	g.name = name
}

func (g *gun) getName() string {
	return g.name
}

func (g *gun) setPower(power int) {
	g.power = power
}

func (g *gun) getPower() int {
	return g.power
}

type ak47 struct {
	gun
}

// Конкретный продукт
func newAk47() weapon {
	return &ak47{
		gun: gun{
			name:  "AK47 gun",
			power: 4,
		},
	}
}

// Конкретный продукт
type glock struct {
	gun
}

func newGlock() weapon {
	return &glock{
		gun: gun{
			name:  "Glock pistol",
			power: 1,
		},
	}
}

// Фабрика
func getGun(gunType string) (weapon, error) {
	if gunType == "ak47" {
		return newAk47(), nil
	}
	if gunType == "musket" {
		return newGlock(), nil
	}
	return nil, fmt.Errorf("Wrong gun type passed")
}

func main() {
	ak47, _ := getGun("ak47")
	musket, _ := getGun("musket")

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g weapon) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
