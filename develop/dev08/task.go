package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func command(strs ...string) {
	cmd := exec.Command(strs[0], strs[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func main() {
	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		cmd := scan.Text()

		args := strings.Split(cmd, " ")

		switch args[0] {
		case "cd":
			if len(args) == 1 {
				userDir, _ := os.UserHomeDir()
				os.Chdir(userDir)
			} else {
				os.Chdir(args[1])
			}
		default:
			if strings.ContainsRune(cmd, '|') {
				command("bash", "-c", cmd)
			} else {
				command(args...)
			}
		}
	}
}
