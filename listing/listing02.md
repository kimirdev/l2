Что выведет программа? Объяснить вывод программы. Объяснить как работают defer’ы и их порядок вызовов.

```go
package main

import (
	"fmt"
)


func test() (x int) {
	defer func() {
		x++
	}()
	x = 1
	return
}


func anotherTest() int {
	var x int
	defer func() {
		x++
	}()
	x = 1
	return x
}


func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
```

Ответ:
```
ф-ция test выдаст 2,
ф-ция anotherTest выдаст 1
...
Выражение defer добавляет вызов функции после ключевого слова defer в стеке приложения. 
Все вызовы в стеке вызываются при возврате функции, в которой они добавлены. 
Поскольку вызовы помещаются в стек, они производятся в порядке от последнего к первому.

В документации говорится (https://go.dev/ref/spec#Defer_statements):
For instance, if the deferred function is a function literal 
and the surrounding function has named result parameters that are in scope within the literal, 
the deferred function may access and modify the result parameters before they are returned. 
If the deferred function has any return values, they are discarded when the function completes

Т.е. если мы имеем ф-цию с "названным" возвращаемым значением, то defer имеет доступ к этому значению и может модифицировать его до возвращения.

```
