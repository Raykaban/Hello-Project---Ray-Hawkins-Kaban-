package main

import "fmt"

func main() {
	var a, i int
	var topi, tulis, tali, pisau, hasil bool
	i = 0
	hasil = true
	fmt.Scanln(&a)
	for i < a && hasil {
		fmt.Scanln(&topi, &tulis, &tali, &pisau)
		hasil = topi && tulis && tali && pisau
		i++
	}
	fmt.Println(hasil)
}
