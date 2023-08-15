package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&num1)

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&num2)

	fmt.Print("Digite o operador (+, -, *, /): ")
	fmt.Scanln(&operator)

	result := 0.0

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Erro: Divisão por zero não é permitida.")
			return
		}
	default:
		fmt.Println("Operador inválido.")
		return
	}

	fmt.Printf("Resultado: %.2f\n", result)
}
