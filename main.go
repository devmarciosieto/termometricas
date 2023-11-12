package main

import (
	"fmt"
)

func main() {
	var temperatura float64
	var escolha int

	fmt.Print("Digite a temperatura: ")
	fmt.Scanln(&temperatura)

	fmt.Println("Escolha a escala para conversão:")
	fmt.Println("1: Celsius para Fahrenheit")
	fmt.Println("2: Celsius para Kelvin")
	fmt.Println("3: Fahrenheit para Celsius")
	fmt.Println("4: Fahrenheit para Kelvin")
	fmt.Println("5: Kelvin para Celsius")
	fmt.Println("6: Kelvin para Fahrenheit")
	fmt.Scanln(&escolha)

	switch escolha {
	case 1:
		fmt.Printf("Temperatura em Fahrenheit: %.2f\n", celsiusParaFahrenheit(temperatura))
	case 2:
		fmt.Printf("Temperatura em Kelvin: %.2f\n", celsiusParaKelvin(temperatura))
	case 3:
		fmt.Printf("Temperatura em Celsius: %.2f\n", fahrenheitParaCelsius(temperatura))
	case 4:
		fmt.Printf("Temperatura em Kelvin: %.2f\n", fahrenheitParaKelvin(temperatura))
	case 5:
		fmt.Printf("Temperatura em Celsius: %.2f\n", kelvinParaCelsius(temperatura))
	case 6:
		fmt.Printf("Temperatura em Fahrenheit: %.2f\n", kelvinParaFahrenheit(temperatura))
	default:
		fmt.Println("Escolha inválida")
	}
}

func celsiusParaFahrenheit(c float64) float64 {
	return (c * 9 / 5) + 32
}

func celsiusParaKelvin(c float64) float64 {
	return c + 273.15
}

func fahrenheitParaCelsius(f float64) float64 {
	return (f - 32) * 5 / 9
}

func fahrenheitParaKelvin(f float64) float64 {
	return fahrenheitParaCelsius(f) + 273.15
}

func kelvinParaCelsius(k float64) float64 {
	return k - 273.15
}

func kelvinParaFahrenheit(k float64) float64 {
	return celsiusParaFahrenheit(kelvinParaCelsius(k))
}
