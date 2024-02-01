package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

/*
Создать Go-функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
"a4bc2d5e" => "aaaabccddddde"
"abcd" => "abcd"
"45" => "" (некорректная строка)
"" => ""

Дополнительно
Реализовать поддержку escape-последовательностей.
Например:
qwe\4\5 => qwe45 (*)
qwe\45 => qwe44444 (*)
qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка, функция должна возвращать ошибку. Написать unit-тесты.
*/

func main() {
	str, err := Unpack("qwe\\\\5")
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: : %v", err)
	}
	fmt.Println(str)
}

func Unpack(input string) (string, error) {
	//проверка на пустую строку
	if input == "" {
		return "", nil
	}

	runes := []rune(input)
	//проверка на валидность строки
	isDig := true
	for i := 0; i < len(runes); i++ {
		if isDig && unicode.IsDigit(runes[i]) {
			return "", errors.New("incorrect input")
		}
		isDig = unicode.IsDigit(runes[i])
	}

	res := make([]rune, 0)
	var prev rune
	//алгоритм распаковки
	for _, r := range input {
		//если буква - пишем
		if unicode.IsLetter(r) {
			res = append(res, r)
			prev = r
		} else if unicode.IsDigit(r) { //если цифра, то пишем предыдущую n раз
			num, err := strconv.Atoi(string(r))
			if err != nil {
				return "", errors.New("ERROR: converting to int")
			}
			for i := 0; i < num-1; i++ {
				res = append(res, prev)
			}
		}
	}
	return string(res), nil
}
