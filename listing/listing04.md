Что выведет программа? Объяснить вывод программы.

```go
package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		// close(ch)
	}()

	for n := range ch {
		println(n)
	}
}
```

Ответ:
```
Вывод в консоль:
числа от 0 до 9,
Ошибка: deadlock
...
Для того, чтобы избежать дедлока, нужно закрыть канал после завершения цикла for [close(ch)]

```
