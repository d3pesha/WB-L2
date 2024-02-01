package main

import (
	"fmt"
	"sort"
	"strings"
)

/*Написать функцию поиска всех множеств анаграмм по словарю.


Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.


Требования:
1. Входные данные для функции: ссылка на массив, каждый элемент которого - слово на русском языке в кодировке utf8
2. Выходные данные: ссылка на мапу множеств анаграмм
3. Ключ - первое встретившееся в словаре слово из множества. Значение - ссылка на массив, каждый элемент которого,
   слово из множества.
4. Массив должен быть отсортирован по возрастанию.
5. Множества из одного элемента не должны попасть в результат.
6. Все слова должны быть приведены к нижнему регистру.
7. В результате каждое слово должно встречаться только один раз.
*/

func Anagram(str []string) map[string][]string {
	m := make(map[string][]string)

	for _, i := range str {
		i = strings.ToLower(i)
		key := sortString(i)
		if len(m[key]) == 0 {
			m[key] = []string{i}
		} else {
			m[key] = append(m[key], i)
		}
	}
	result := make(map[string][]string)
	for _, value := range m {
		if len(value) > 1 {
			sort.Strings(value)

			result[value[0]] = value
		}
	}

	return result
}

func sortString(s string) string {
	r := []rune(s)
	sort.Slice(r, func(i, j int) bool {
		return r[i] < r[j]
	})
	return string(r)
}

func main() {
	fmt.Println(Anagram([]string{"Апельсин", "Лимон", "Спаниель", "Мелисса", "Лиса"}))
}
