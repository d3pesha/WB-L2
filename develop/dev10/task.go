package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
Реализовать простейший telnet-клиент.

Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123


Требования:
Программа должна подключаться к указанному хосту (ip или доменное имя + порт) по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s)
При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout
*/
func Exit(ch <-chan os.Signal) {
	for {
		switch <-ch {
		case syscall.SIGQUIT:
			os.Exit(0)
		default:
			
		}
	}
}

func main() {
	exitCh := make(chan os.Signal, 1)

	signal.Notify(exitCh)

	go Exit(exitCh)

	timeout := flag.String("timeout", "10s", "timeout for a connection")

	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Println("Example: task.go host port")
		return
	}
	timeoutDuration, err := time.ParseDuration(*timeout)
	if err != nil {
		return
	}
	host := flag.Arg(0)
	port := flag.Arg(1)

	hostPort := host + ":" + port

	start := time.Now()
	start = start.Add(timeoutDuration)

	var conn net.Conn

	fmt.Printf("Trying to coonect for %s:%s", host, port)
	for start.After(time.Now()) {
		conn, _ = net.DialTimeout("tcp", hostPort, timeoutDuration)
		break
	}

	if err != nil {
		fmt.Printf("Connection error: %s", err)
		return
	}

	fmt.Println("Connected")

	defer conn.Close()

	go func() {
		reader := bufio.NewReader(conn)
		for {
			msg, err := reader.ReadString('\n')
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Print("Message from server: " + msg)
		}
	}()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		_, err := fmt.Fprintf(conn, input+"\n")
		if err != nil {
			log.Fatal("Connection close")
		}
	}
}
