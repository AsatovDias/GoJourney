package main

import (
	"fmt"
	"os"
)

func main() {
	filename := "example.txt"
	text := "Привет! Это пример работы с файлами в Go."

	// 1. Запись в файл
	err := os.WriteFile(filename, []byte(text), 0644)
	if err != nil {
		fmt.Println("Ошибка при записи файла:", err)
		return
	}

	fmt.Println("Файл успешно создан и записан.")

	// 2. Чтение файла
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Ошибка при чтении файла:", err)
		return
	}

	fmt.Println("Содержимое файла:")
	fmt.Println(string(data))
}
