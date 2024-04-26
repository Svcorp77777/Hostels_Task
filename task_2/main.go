package main

import "fmt"

// Написать функцию/метод, которая на вход получает массив положительных
// целых чисел произвольной длины.
// Например [42, 12, 18]. На выходе возвращает массив чисел, которые
// являются общими делителями для всех указанных числе. В примере это будет [2, 3, 6].

func main() {
	array := []int{42, 12, 18}

	result := commonDivisor(array)

	fmt.Println(result)
}

func commonDivisor(arr []int) []int {
	res := make([]int, 0)
	result := make([]int, 0)

	for _, val := range arr {
		for i := 1; i < int(float64(val)*0.5+1); i++ {
			if val%i == 0 {
				res = append(res, i, val/i)
			}
		}
	}

	res = removDuplicat(res)

	for _, val := range res {
		if val == 1 {
			continue
		}

		flag := false
		for _, num := range arr {
			if num%val != 0 {
				flag = false

				break
			}

			flag = true
		}

		if flag {
			result = append(result, val)
		}
	}

	return result
}

func removDuplicat(arr []int) (res []int) {
	var chek map[int]struct{} = make(map[int]struct{})

	for _, val := range arr {
		if _, exist := chek[val]; !exist {
			res = append(res, val)
			chek[val] = struct{}{}

		}
	}

	return res
}
