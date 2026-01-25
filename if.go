package main

import "fmt"

func main() {
	// if go run
	age := 40

	if age > 50 {
		fmt.Println("Ты старый")
	} else if age < 15 {
		fmt.Println("ты дибил как мой папа")
	} else {
		fmt.Println("Средний возраст")
	}
}
