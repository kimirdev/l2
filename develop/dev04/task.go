package main

import (
	"fmt"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func getAnagrams(strs []string) map[string][]string {
	set := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		strs[i] = strings.ToLower(strs[i])
	}

	for _, str := range strs {
		runes := []rune(str)

		sort.SliceStable(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		key := string(runes)
		_, ok := set[key]
		if !ok {
			set[key] = make([]string, 0)
		}
		set[key] = append(set[key], str)
	}

	ret := make(map[string][]string)
	for _, el := range set {
		if len(el) > 1 {
			sort.SliceStable(el, func(i, j int) bool {
				return el[i] < el[j]
			})
			ret[el[0]] = el
		}
	}

	return ret
}

func main() {
	strs := []string{"пятак",
		"пятка",
		"тяпка",
		"листок",
		"слиток",
		"выфв",
		"ыввф",
		"столик",
	}

	anagrs := getAnagrams(strs)

	fmt.Println(anagrs, "test")
}
