package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type customFlags struct {
	fieldsFlag    []int
	delimiterFlag string
	separatedFlag bool
}

type customCut struct {
	strs  []string
	flags customFlags
}

func usage() {
	fmt.Printf(`./cut [flags]
	-f - выбрать поля (колонки)
	-d -  использовать другой разделитель
	-s - только строки с разделителем
`)
}

func readFlags(flgs *customFlags) error {
	var fieldsStr string
	flag.StringVar(&fieldsStr, "f", "", "указание колонки для сортировки")
	flag.StringVar(&flgs.delimiterFlag, "d", "", "указание колонки для сортировки")
	flag.BoolVar(&flgs.separatedFlag, "s", false, "указание колонки для сортировки")

	flag.Usage = usage

	flag.Parse()

	fields := strings.Split(fieldsStr, ",")

	flgs.fieldsFlag = make([]int, 0)
	for _, f := range fields {
		field, err := strconv.Atoi(strings.Trim(f, " "))
		if err != nil {
			return errors.New("invalid field")
		}
		flgs.fieldsFlag = append(flgs.fieldsFlag, field)
		sort.Ints(flgs.fieldsFlag)
	}

	if len(flgs.delimiterFlag) > 1 {
		return errors.New("invalid delimiter")
	}

	if flgs.delimiterFlag == "" {
		flgs.delimiterFlag = "\t"
	}
	return nil
}

func readInput() []string {
	ret := make([]string, 0)
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		ret = append(ret, s.Text())
	}
	return ret
}

func (c *customCut) Cut() []string {
	res := make([]string, 0)

	for _, l := range c.strs {
		if strings.Contains(l, c.flags.delimiterFlag) {
			parts := strings.Split(l, c.flags.delimiterFlag)
			sb := strings.Builder{}

			for i, f := range c.flags.fieldsFlag {
				if len(parts) >= f {
					if i != 0 {
						sb.WriteByte(' ')
					}
					sb.WriteString(parts[f-1])
				}
			}
			res = append(res, sb.String())
		} else {
			if !c.flags.separatedFlag {
				res = append(res, l)
			}
		}
	}

	return res
}

func main() {
	cut := &customCut{}

	err := readFlags(&cut.flags)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(2)
	}

	cut.strs = readInput()

	for _, str := range cut.Cut() {
		fmt.Println(str)
	}
}
