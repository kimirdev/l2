package main

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/
import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func usage() {
	fmt.Println("./telnet [flag] host port\n	flag: -timeout")
}

func main() {
	timeout := flag.Duration("timeout", 10*time.Second, "timeout")

	flag.Usage = usage
	flag.Parse()

	args := flag.Args()

	if len(args) != 2 {
		usage()
		os.Exit(2)
	}

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(args[0], args[1]), *timeout)

	if err != nil {
		time.Sleep(*timeout)
		fmt.Fprintf(os.Stderr, "timeout")
		os.Exit(2)
	}

	sigChan := make(chan os.Signal, 1)
	errChan := make(chan error, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func(conn net.Conn, errChan chan<- error, sigChan chan<- os.Signal) {
		for {
			reader := bufio.NewReader(os.Stdin)
			text, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					sigChan <- syscall.Signal(syscall.SIGQUIT)
					return
				}
				errChan <- err
				return
			}

			fmt.Fprintf(conn, text+"\n")
		}
	}(conn, errChan, sigChan)

	go func(conn net.Conn, errChan chan<- error, sigChan chan<- os.Signal) {
		for {
			reader := bufio.NewReader(conn)
			text, err := reader.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					sigChan <- syscall.Signal(syscall.SIGQUIT)
					return
				}
				errChan <- err
				return
			}

			fmt.Println(text)
		}
	}(conn, errChan, sigChan)

	select {
	case <-sigChan:
		conn.Close()
	case err = <-errChan:
		fmt.Fprintf(os.Stderr, err.Error())
		os.Exit(2)
	}
}
