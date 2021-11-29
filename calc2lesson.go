package main

import "fmt"

func main() {

	var number1, number2 float64
	var operator string

	fmt.Print("Введите первое число :")
	fmt.Scanln(&number1)

	fmt.Print("Введите второе число :")
	fmt.Scanln(&number2)

	fmt.Print("Введите операцию (+-*/): ")
	fmt.Scanln(&operator)

	switch operator {
	case "+":
		fmt.Printf("%f %s %f = %f", number1, operator, number2, number1+number2)

	case "-":
		fmt.Printf("%f %s %f = %f", number1, operator, number2, number1-number2)

	case "*":
		fmt.Printf("%f %s %f = %f", number1, operator, number2, number1*number2)

	case "/":
		if number2 == 0.0 {
			fmt.Println(" ошибка деления на ноль")
		} else {
			fmt.Printf("%f %s %f = %f", number1, operator, number2, number1/number2)

		}

	}
}
