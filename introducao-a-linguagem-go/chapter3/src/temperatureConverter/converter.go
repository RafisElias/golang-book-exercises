package temperatureConverter

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TemperatureConverter() {
	fmt.Println("Certo converter temperatura.")
	fmt.Print("\n")

	fmt.Println("Qual conversão você quer fazer?")

	fmt.Println("1 - Converter Fahrenheit para Celsius:")
	fmt.Println("2 - Converter Fahrenheit para Kelvin:")
	fmt.Println("3 - Converter Celsius para Fahrenheit:")
	fmt.Println("4 - Converter Celsius para Kelvin:")
	fmt.Println("5 - Converter Kelvin para Celsius:")
	fmt.Println("6 - Converter Kelvin para Fahrenheit:")
	fmt.Println("0 - Para sair do sistema")

	fmt.Print("\n")

	fmt.Print("Selecione uma opção: ")
	option := bufio.NewScanner(os.Stdin)
	option.Scan()

	switch option.Text() {
	case "1":
		fmt.Println("1 - Converter Fahrenheit para Celsius:")
		convertFahrenheitToCelsius()
	case "2":
		fmt.Println("2 - Converter Fahrenheit para Kelvin:")
		convertFahrenheitToKelvin()
	case "3":
		fmt.Println("3 - Converter Celsius para Fahrenheit:")
		convertCelsiusToFahrenheit()
	case "4":
		fmt.Println("4 - Converter Celsius para Kelvin:")
		convertCelsiusToKelvin()
	case "5":
		fmt.Println("5 - Converter Kelvin para Celsius:")
		convertKelvinToCelsius()
	case "6":
		fmt.Println("6 - Converter Kelvin para Fahrenheit:")
		convertKelvinToFahrenheit()
	case "0":
		fmt.Println("Saindo do sistema")
		os.Exit(1)
	default:
		fmt.Println("Comando invalido")
		os.Exit(1)
	}
}

func convertFahrenheitToCelsius() {
	fmt.Print("Insira a temperatura em Fahrenheit: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	temperatureConverted, _ := strconv.ParseFloat(scanner.Text(), 64)
	temperature := (temperatureConverted - 32) * (5.0 / 9.0)
	fmt.Printf("Temperatura em Celsius equivale a: %.2f ºC", temperature)
	fmt.Print("\n")
}

func convertFahrenheitToKelvin() {
	fmt.Print("Insira a temperatura em Fahrenheit: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	temperatureConverted, _ := strconv.ParseFloat(scanner.Text(), 64)
	temperature := (temperatureConverted-32)*(5.0/9.0) + 273.15
	fmt.Printf("Temperatura em Kelvin equivale a: %.2f K", temperature)
	fmt.Print("\n")
}

func convertCelsiusToFahrenheit() {
	fmt.Print("Insira a temperatura em Celsius: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	temperatureConverted, _ := strconv.ParseFloat(scanner.Text(), 64)
	temperature := (temperatureConverted * (9.0 / 5.0)) + 32
	fmt.Printf("Temperatura em Fahrenheit equivale a: %.2f ºF", temperature)
	fmt.Print("\n")
}

func convertCelsiusToKelvin() {
	fmt.Print("Insira a temperatura em Celsius: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	temperatureConverted, _ := strconv.ParseFloat(scanner.Text(), 64)
	temperature := temperatureConverted + 273.15
	fmt.Printf("Temperatura em Kelvin equivale a: %.2f K", temperature)
	fmt.Print("\n")
}

func convertKelvinToCelsius() {
	fmt.Print("Insira a temperatura em Kelvin: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	temperatureConverted, _ := strconv.ParseFloat(scanner.Text(), 64)
	temperature := temperatureConverted - 273.15
	fmt.Printf("Temperatura em Celsius equivale a: %.2f ºC", temperature)
	fmt.Print("\n")
}

// (10 K − 273,15) × 9/5 + 32
func convertKelvinToFahrenheit() {
	fmt.Print("Insira a temperatura em Kelvin: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	temperatureConverted, _ := strconv.ParseFloat(scanner.Text(), 64)
	temperature := (temperatureConverted-273.15)*(9.0/5.0) + 32
	fmt.Printf("Temperatura em Fahrenheit equivale a: %.2f ºF", temperature)
	fmt.Print("\n")
}
