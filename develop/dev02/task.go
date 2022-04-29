package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var usage string = `type:
./main [string]
`

var errUnpack error = errors.New("invalid string")

func unpack(str string) (string, error) {
	var sb strings.Builder
	var repeatStart int
	var repeatChar byte

	for i := 0; i < len(str); i++ {
		if str[i] == '\\' {
			if i+1 == len(str) {
				return "", errUnpack
			}
			if i+2 >= len(str) || str[i+2] < '0' || str[i+2] > '9' {
				sb.WriteByte(str[i+1])
			}
			i++
		} else if str[i] >= '0' && str[i] <= '9' {
			if i == 0 {
				return "", errUnpack
			}
			repeatChar = str[i-1]
			repeatStart = i
			for i+1 < len(str) && str[i+1] >= '0' && str[i+1] <= '9' {
				i++
			}
			count, err := strconv.Atoi(str[repeatStart : i+1])
			if err != nil {
				return "", errUnpack
			}
			for ; count > 0; count-- {
				sb.WriteByte(repeatChar)
			}
		} else if i+1 >= len(str) || str[i+1] < '0' || str[i+1] > '9' {
			sb.WriteByte(str[i])
		}
	}
	return sb.String(), nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println(usage)
		return
	}

	str := os.Args[1]

	newStr, err := unpack(str)

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(2)
	}

	fmt.Printf("%q\n", newStr)
}
