package dev01

import (
	"fmt"
	"github.com/beevik/ntp"
	"os"
	"time"
)

/*
Создать программу печатающую точное время с использованием NTP -библиотеки. Инициализировать как go module. Использовать библиотеку github.com/beevik/ntp. Написать программу печатающую текущее время / точное время с использованием этой библиотеки.

Требования:
Программа должна быть оформлена как go module
Программа должна корректно обрабатывать ошибки библиотеки: выводить их в STDERR и возвращать ненулевой код выхода в OS
*/

func getTimeNow() (time.Time, error) {
	t, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	return t, err
}

func main() {
	timeNow, err := getTimeNow()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %v", err)
		os.Exit(1)
	}
	fmt.Println(timeNow)
}
