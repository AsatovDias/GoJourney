package main

import "fmt"

func main() {
	fmt.Println("1. Обычный for")
	for i := 1; i <= 5; i++ {
		fmt.Println("i =", i)
	}

	fmt.Println("\n2. for как while")
	n := 1
	for n <= 3 {
		fmt.Println("n =", n)
		n++
	}

	fmt.Println("\n3. Бесконечный for + break")
	counter := 1
	for {
		fmt.Println("counter =", counter)
		if counter == 3 {
			fmt.Println("Выходим из цикла через break")
			break
		}
		counter++
	}

	fmt.Println("\n4. continue")
	for i := 1; i <= 5; i++ {
		if i == 3 {
			fmt.Println("Пропускаем i =", i)
			continue
		}
		fmt.Println("Обрабатываем i =", i)
	}

	fmt.Println("\n5. break")
	for i := 1; i <= 10; i++ {
		if i == 6 {
			fmt.Println("Останавливаем цикл на i =", i)
			break
		}
		fmt.Println("i =", i)
	}

	fmt.Println("\n6. range по срезу")
	numbers := []int{10, 20, 30, 40}

	for index, value := range numbers {
		fmt.Println("index =", index, "value =", value)
	}

	fmt.Println("\n7. range только по значениям")
	for _, value := range numbers {
		fmt.Println("value =", value)
	}

	fmt.Println("\n8. range только по индексам")
	for index := range numbers {
		fmt.Println("index =", index)
	}

	fmt.Println("\n9. range по строке")
	word := "Go Язык"

	for index, symbol := range word {
		fmt.Printf("index: %d, symbol: %c\n", index, symbol)
	}

	fmt.Println("\n10. range по map")
	studentScores := map[string]int{
		"Aruzhan": 95,
		"Ali":     88,
		"Daniyar": 91,
	}

	for name, score := range studentScores {
		fmt.Println("Студент:", name, "Балл:", score)
	}
}
