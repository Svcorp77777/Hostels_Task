package main

import "fmt"

// Разработайте функцию, которая принимает целое число в качестве аргумента и возвращает строку,
// содержащую это число и слово "компьютер" в нужном склонении по падежам в зависимости от числа.
//  Например, при вводе числа 25 функция должна возвращать "25 компьютеров", для числа 41 — "41 компьютер",
//  а для числа 1048 — "1048 компьютеров".

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 13, 14, 20, 21, 22, 30, 25, 41, 1000, 1048}

	for _, val := range arr {
		result := declension(val)
		fmt.Println(result)
	}
}

func declension(num int) string {
	value := num % 100
	if value > 20 {
		value = value % 10
	}

	switch {
	case value == 1:
		return fmt.Sprintf("%d компьютер", num)
	case value > 1 && value < 5:
		return fmt.Sprintf("%d компьютера", num)
	case value > 4 && value <= 20 || value == 0:
		return fmt.Sprintf("%d компьютеров", num)
	default:
		return fmt.Sprintln("Что то пошло не так")
	}
}
