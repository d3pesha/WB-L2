Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
nil
false 
Функция Foo() возвращает значение nil типа *os.PathError. Однако, любое 
значение == nil только в случае,когда и значение и тип являются nil.
 Поэтому сравнение err == nil будет false, т.к. *os.PathError != nil

```