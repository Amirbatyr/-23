package main

import "fmt"

func printInfo(done chan<- bool) {
	fmt.Println("Горутина запущена")
	done <- true
}

func sayHello(name string, done chan<- bool) {
	fmt.Printf("Привет, %s\n", name)
	done <- true
}

func heavyTask(done chan<- bool) {
	fmt.Println("Выполняю задачу...")
	done <- true
}

func count(id int, done chan<- bool) {
	for number := 1; number <= 3; number++ {
		fmt.Printf("Горутина %d: %d\n", id, number)
	}
	done <- true
}

func calculateSquare(number int, done chan<- bool) {
	fmt.Printf("Квадрат числа %d: %d\n", number, number*number)
	done <- true
}

func checkStatus(address string, done chan<- bool) {
	fmt.Printf("Сайт %s доступен\n", address)
	done <- true
}

func processData(id int, category string, delayMs int, done chan<- bool) {
	fmt.Printf("Запуск процесса %d в категории %s\n", id, category)
	fmt.Printf("Процесс %d завершен (заданная задержка %d мс не применена)\n", id, delayMs)
	done <- true
}

func waitFor(done <-chan bool, n int) {
	for i := 0; i < n; i++ {
		<-done
	}
}

func main() {
	done := make(chan bool)

	// 1
	go printInfo(done)
	waitFor(done, 1)

	// 2
	go sayHello("Анна", done)
	go sayHello("Иван", done)
	go sayHello("Мария", done)
	waitFor(done, 3)

	// 3
	go heavyTask(done)
	waitFor(done, 1)

	// 4
	for id := 1; id <= 5; id++ {
		go count(id, done)
	}
	waitFor(done, 5)

	// 5
	numbers := []int{2, 4, 6, 8}
	for _, number := range numbers {
		go calculateSquare(number, done)
	}
	waitFor(done, len(numbers))

	// 6
	addresses := []string{"google.com", "yandex.ru", "go.dev"}
	for _, address := range addresses {
		go checkStatus(address, done)
	}
	waitFor(done, len(addresses))

	// 7
	var categories [3]string
	var delays [3]int

	for i := 0; i < 3; i++ {
		fmt.Printf("Процесс %d — введите категорию (одним словом): ", i+1)
		if _, err := fmt.Scan(&categories[i]); err != nil {
			fmt.Println("Не удалось прочитать категорию:", err)
			return
		}

		fmt.Printf("Процесс %d — введите задержку в миллисекундах: ", i+1)
		if _, err := fmt.Scan(&delays[i]); err != nil || delays[i] < 0 {
			fmt.Println("Задержка должна быть целым неотрицательным числом миллисекунд.")
			return
		}
	}

	for i := 0; i < 3; i++ {
		go processData(i+1, categories[i], delays[i], done)
	}
	waitFor(done, 3)
}
