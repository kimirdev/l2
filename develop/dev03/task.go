package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func usage() {
	fmt.Printf(`./sort [options] [filename]
	-k указание колонки для сортировки
	-n сортировать в обратном порядке
	-r сортировать в обратном порядке
	-u сортировать в обратном порядке
`)
}

func readFlags(columnFlag *int, numFlag, reverseFlag, uniqFlag *bool) {
	flag.IntVar(columnFlag, "k", 0, "указание колонки для сортировки")
	flag.BoolVar(numFlag, "n", false, "сортировать по числовому значению")
	flag.BoolVar(reverseFlag, "r", false, "сортировать в обратном порядке")
	flag.BoolVar(uniqFlag, "u", false, "не выводить повторяющиеся строки")

	flag.Usage = usage

	flag.Parse()
}

func readFile() []string {
	args := flag.Args()

	if len(args) != 1 {
		flag.Usage()
		os.Exit(2)
	}

	filename := args[len(args)-1]

	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	strs := strings.Split(string(bytes), "\n")

	return strs
}

func sortUtil(strs []string, columnFlag int, numFlag, reverseFlag, uniqFlag bool) []string {
	var uniqSet map[string]struct{}
	if uniqFlag {
		sliceSet := make([]string, 0)
		uniqSet = make(map[string]struct{})
		for _, str := range strs {
			_, ok := uniqSet[str]
			if !ok {
				sliceSet = append(sliceSet, str)
				uniqSet[str] = struct{}{}
			}
		}
		strs = sliceSet
	}

	if columnFlag != 0 {
		sort.Slice(strs, func(i, j int) bool {
			iStr := strings.Split(strs[i], " ")
			jStr := strings.Split(strs[j], " ")

			if len(iStr) < columnFlag && len(jStr) >= columnFlag {
				return true
			} else if len(jStr) < columnFlag && len(iStr) >= columnFlag {
				return false
			} else if len(iStr) < columnFlag && len(jStr) < columnFlag {
				return strs[i] < strs[j]
			}
			if numFlag {
				iNum, iErr := strconv.Atoi(iStr[columnFlag-1])
				jNum, jErr := strconv.Atoi(jStr[columnFlag-1])

				if iErr != nil || jErr != nil {
					return iStr[columnFlag-1] < jStr[columnFlag-1]
				}
				return iNum < jNum
			}
			return iStr[columnFlag-1] < jStr[columnFlag-1]
		})
	} else {
		sort.Strings(strs)
	}

	if reverseFlag {
		for i, j := 0, len(strs)-1; i < j; i, j = i+1, j-1 {
			strs[i], strs[j] = strs[j], strs[i]
		}
	}

	return strs
}

func main() {
	var (
		columnFlag  int
		numFlag     bool
		reverseFlag bool
		uniqFlag    bool
	)

	readFlags(&columnFlag, &numFlag, &reverseFlag, &uniqFlag)

	strs := readFile()

	strs = sortUtil(strs, columnFlag, numFlag, reverseFlag, uniqFlag)

	for _, str := range strs {
		fmt.Println(str)
	}
}
