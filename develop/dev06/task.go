package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

/*Реализовать утилиту аналог консольной команды cut (man cut). Утилита должна принимать строки через STDIN, разбивать по разделителю (TAB) на колонки и выводить запрошенные.

Реализовать поддержку утилитой следующих ключей:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
*/

type flags struct {
	field     int
	delimiter string
	separated bool
}

func main() {
	field := flag.Int("f", 0, "выбор поля (колонки)")
	delimiter := flag.String("d", " ", "использовать другой разделитьель")
	separated := flag.Bool("s", false, "только строки с разделителем")

	flag.Parse()
	flags := flags{*field, *delimiter, *separated}

	scanner := bufio.NewScanner(os.Stdin)
	for {
		ok := scanner.Scan()
		if !ok || scanner.Err() != nil {
			log.Fatal("Scanner error")
		}
		if scanner.Text() != "" {
			flags.printString(scanner.Text())
		}
	}
}

func (f *flags) printString(str string) {
	split := strings.Split(str, f.delimiter)

	if f.separated {
		if f.field > 0 && len(split) > 1 {
			if len(split) >= f.field {
				fmt.Println(split[f.field-1])
			} else {
				fmt.Println("")
			}
		}
	} else {
		if f.field > 0 && len(split) >= f.field {
			fmt.Println(split[f.field-1])
		} else {
			fmt.Println(str)
		}
	}
}
