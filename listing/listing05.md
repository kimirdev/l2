Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
Вывод в консоль:
error
...
Как и в задании listing03:
error - interface
type - *customError
value - nil

interface == nil - будет истинной, только в случае, если type и value равны nil

```
