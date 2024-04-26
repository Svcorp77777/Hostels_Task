package main

import "fmt"

// Написать метод, который в консоль выводит таблицу умножения.
// На вход метод получает число, до которого выводит таблицу умножения.
// В консоли должна появиться таблица.

func main() {
	creatMultiplicationTable(5)
}

func creatMultiplicationTable(n int) {
	for i := 0; i <= n; i++ {
		if i == 0 {
			fmt.Print("   ")

			continue
		}

		fmt.Printf("%d  ", i)
	}

	fmt.Println()

	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", i)

		for j := 1; j <= n; j++ {
			res := i * j
			sp := ""

			if res <= 9 {
				sp = " "
			}

			fmt.Printf("%s%d ", sp, res)
		}

		fmt.Println()
	}
}
