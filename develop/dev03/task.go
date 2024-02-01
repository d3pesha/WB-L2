package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
Отсортировать строки в файле по аналогии с консольной утилитой sort (man sort — смотрим описание и основные параметры): на входе подается файл из несортированными строками, на выходе — файл с отсортированными.

Реализовать поддержку утилитой следующих ключей:

-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок, по умолчанию разделитель — пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

# Дополнительно

Реализовать поддержку утилитой следующих ключей:

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учетом суффиксов
*/

func readingFile(name string) ([]string, error) {
	var res []string
	open, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer open.Close()

	scanner := bufio.NewScanner(open)

	for scanner.Scan() {
		str := scanner.Text()
		res = append(res, str)
	}
	return res, nil
}

func recordFile(file string, arr []string) error {
	outFile, err := os.Create(file)
	if err != nil {
		return err
	}
	defer outFile.Close()

	for i := 0; i < len(arr)-1; i++ {
		outFile.WriteString(arr[i] + "\n")
	}
	outFile.WriteString(arr[len(arr)-1])

	return nil
}

func UniqueString(noneUniqStr []string) []string {
	for i, str := range noneUniqStr {
		for j := i + 1; j < len(noneUniqStr); j++ {
			if str == noneUniqStr[j] {
				noneUniqStr = append(noneUniqStr[:i], noneUniqStr[j:]...)

			}
		}
	}
	return noneUniqStr
}

func SortByColumn(data []string, column int) []string {
	sort.Slice(data, func(i, j int) bool {
		wordsI := strings.Fields(data[i])
		wordsJ := strings.Fields(data[j])

		if column < len(wordsI) && column < len(wordsJ) {
			return wordsI[column] < wordsJ[column]
		}

		// Если колонка выходит за пределы строки, считаем, что строки равны
		return false
	})
	return data
}

func SortNum(data []string) []string {
	for i := range data {
		var count1 int
		l := strings.Split(data[i], " ")

		for j := range l {
			_, err := strconv.Atoi(l[j])
			if err == nil {
				count1++
			}
		}

		if count1 == len(l) {
			var result []int

			for j := range l {
				k, _ := strconv.Atoi(l[j])
				result = append(result, k)
			}
			sort.Ints(result)

			for m := range result {
				l = append(l[:m], strconv.Itoa(result[m]))
			}
		}
		data[i] = strings.Join(l, " ")
	}
	return data
}

func ReversSort(unsorted []string) []string {
	sort.Sort(sort.Reverse(sort.StringSlice(unsorted)))
	return unsorted
}

func main() {

	k := flag.Int("k", 0, "enter column")
	n := flag.Bool("n", false, "sort by num")
	r := flag.Bool("r", false, "revers sort")
	u := flag.Bool("u", false, "do not dublicate lines")

	flag.Parse()

	dataFile, _ := readingFile("test.txt")

	switch {
	case *u:
		dataFile = UniqueString(dataFile)
	case *k != 0:
		dataFile = SortByColumn(dataFile, *k)
	case *r == true:
		dataFile = ReversSort(dataFile)
	case *n == true:
		dataFile = SortNum(dataFile)
	}
	err := recordFile("outText.txt", dataFile)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
}
