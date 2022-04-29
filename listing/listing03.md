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
Вывод в консоль:
<nil>
false
...
interface в голанг под капотом имеет 2 элемента: type и value

interface == nil - будет истинной, только в случае, если type и value равны nil

ф-ция Foo возвращает интерфейс error со значениями:
type = *os.PathError
value = nil
```
