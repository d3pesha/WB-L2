package main

import "fmt"

// SortingStrategy - интерфейс стратегии сортировки
type SortingStrategy interface {
	Sort(numbers []int)
}

// BubbleSort - конкретная стратегия сортировки пузырьком
type BubbleSort struct{}

func (bs *BubbleSort) Sort(numbers []int) {
	n := len(numbers)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
}

// QuickSort - конкретная стратегия быстрой сортировки
type QuickSort struct{}

func (qs *QuickSort) Sort(numbers []int) {
	// Реальная реализация быстрой сортировки
	quickSort(numbers, 0, len(numbers)-1)
	fmt.Println("Using QuickSort")
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)

		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] < pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

// Context - контекст, использующий стратегию
type Context struct {
	strategy SortingStrategy
}

// SetStrategy - метод для установки стратегии сортировки
func (c *Context) SetStrategy(strategy SortingStrategy) {
	c.strategy = strategy
}

// Sort - метод для вызова сортировки через текущую стратегию
func (c *Context) Sort(numbers []int) {
	c.strategy.Sort(numbers)
}

func main() {
	// Создаем контекст
	context := &Context{}

	// Используем стратегию сортировки пузырьком
	bubbleSort := &BubbleSort{}
	context.SetStrategy(bubbleSort)

	// Создаем список чисел
	numbers := []int{4, 2, 7, 1, 9, 5}

	// Вызываем сортировку через текущую стратегию
	context.Sort(numbers)
	fmt.Println("Sorted numbers:", numbers)

	// Изменяем стратегию на быструю сортировку
	quickSort := &QuickSort{}
	context.SetStrategy(quickSort)

	// Вызываем сортировку через новую стратегию
	context.Sort(numbers)
	// Вывод: Using QuickSort
}

/*
	Стратегия — это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает
	каждый из них в собственный класс, после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.
Плюсы:
	1. Горячая замена алгоритмов на лету.
	2. Изолирует код и данные алгоритмов от остальных классов.
	3. Уход от наследования к делегированию.
	4. Реализует принцип открытости/закрытости.
Минусы:
	1. Усложняет программу за счёт дополнительных классов.
	2. Клиент должен знать, в чём состоит разница между стратегиями, чтобы выбрать подходящую.
*/
