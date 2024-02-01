package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

/*Реализовать утилиту фильтрации по аналогии с консольной утилитой (man grep — смотрим описание и основные параметры).


Реализовать поддержку утилитой следующих ключей:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", напечатать номер строки
*/

func readFile(name string) ([]string, error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	byteFile, _ := io.ReadAll(file)
	str := string(byteFile)
	return strings.Split(str, "\n"), nil
}

func main() {
	after := flag.Int("A", 0, "after key")
	before := flag.Int("B", 0, "before key")
	context := flag.Int("C", 0, "context key")

	count := flag.Bool("c", false, "count key")
	ignoreCase := flag.Bool("i", false, "ignore-case key")
	inverter := flag.Bool("v", false, "iverter key")
	fixed := flag.Bool("fixed", false, "fixed key")
	lineNum := flag.Bool("n", false, "live num key")

	flag.Parse()
	pattern := flag.Arg(0)
	files := flag.Args()[1:]

	// Вызов функции фильтрации для каждого файла
	for _, file := range files {
		matchLines := filterFile(file, pattern, *after, *before, *context, *count, *ignoreCase, *inverter, *fixed, *lineNum)
		printMatchLines(matchLines)
	}
}

func filterFile(file, pattern string, after, before, context int, count, ignoreCase, inverter, fixed, lineNum bool) []string {
	matchLines := make([]string, 0)

	// Открытие файла для чтения
	f, err := os.Open(file)
	if err != nil {
		fmt.Printf("Cannot open file: %s\n", err)
		return matchLines
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lineNumber := 0
	prevLines := make([]string, 0)
	found := false

	for scanner.Scan() {
		line := scanner.Text()
		lineNumber++

		match := false

		// Проверка на совпадение в зависимости от настроек фильтрации
		if fixed {
			// Точное совпадение со строкой (без использования регулярных выражений)
			match = strings.Contains(line, pattern)
		} else {
			// Совпадение с паттерном (с учетом или без учета регистра)
			if ignoreCase {
				match = strings.Contains(strings.ToLower(line), strings.ToLower(pattern))
			} else {
				match = strings.Contains(line, pattern)
			}
		}

		if match && !inverter {
			if count {
				// Если указан флаг -c, увеличиваем счетчик
				matchLines = append(matchLines, line)
			} else {
				// Сохраняем строки до совпадения
				if before > 0 {
					for _, prevLine := range prevLines {
						matchLines = append(matchLines, prevLine)
					}
				}

				// Сохраняем текущую строку
				if lineNum {
					line = fmt.Sprintf("%d:%s", lineNumber, line)
				}
				matchLines = append(matchLines, line)

				// Сохраняем строки после совпадения
				if after > 0 {
					afterLines := make([]string, 0)
					for i := 0; i < after; i++ {
						if scanner.Scan() {
							afterLine := scanner.Text()
							matchLines = append(matchLines, afterLine)
							afterLines = append(afterLines, afterLine)
						} else {
							break
						}
					}
					prevLines = afterLines
				} else {
					prevLines = make([]string, 0)
				}

				// Сохраняем контекстные строки
				if context > 0 {
					contextLines := make([]string, 0)
					for i := 0; i < context; i++ {
						if scanner.Scan() {
							contextLine := scanner.Text()
							matchLines = append(matchLines, contextLine)
							contextLines = append(contextLines, contextLine)
						} else {
							break
						}
					}
					prevLines = contextLines
				}
			}

			found = true
		} else if !match && inverter {
			if count {
				// Если указан флаг -c, увеличиваем счетчик
				matchLines = append(matchLines, line)
			} else {
				// Сохраняем строки до и после совпадения
				if (before > 0 || after > 0) && (prevLines != nil || after > 0) {
					if lineNum {
						line = fmt.Sprintf("%d:%s", lineNumber, line)
					}
					matchLines = append(matchLines, line)
				}

				// Сохраняем контекстные строки
				if context > 0 {
					contextLines := make([]string, 0)
					for i := 0; i < context; i++ {
						if scanner.Scan() {
							contextLine := scanner.Text()
							matchLines = append(matchLines, contextLine)
							contextLines = append(contextLines, contextLine)
						} else {
							break
						}
					}
					prevLines = contextLines
				} else {
					prevLines = make([]string, 0)
				}
			}

			found = true
		} else {
			// Сохраняем строки до и после совпадения
			if (before > 0 || after > 0) && (prevLines != nil || after > 0) {
				prevLines = append(prevLines, line)

				if len(prevLines) > before+after {
					prevLines = prevLines[1:]
				}
			} else if context > 0 {
				prevLines = append(prevLines, line)

				if len(prevLines) > context {
					prevLines = prevLines[1:]
				}
			}
		}
	}

	if count && !found {
		matchLines = []string{"0"}
	}

	return matchLines
}

// Функция вывода совпадающих строк
func printMatchLines(matchLines []string) {
	for _, line := range matchLines {
		fmt.Println(line)
	}
}
