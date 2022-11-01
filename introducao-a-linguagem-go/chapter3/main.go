package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	hc "github.com/RafisElias/golang-book-exercises/introducao-a-linguagem-go/chapter3/src/heightConverter"
	tc "github.com/RafisElias/golang-book-exercises/introducao-a-linguagem-go/chapter3/src/temperatureConverter"
)

func main() {
	fmt.Println("O que você deseja fazer:")
	fmt.Println("1 - Converter temperaturas:")
	fmt.Println("2 - Converter altura:")
	fmt.Println("0 - Para sair:")
	fmt.Print("\n")

	fmt.Print("Selecione uma opção: ")
	option := bufio.NewScanner(os.Stdin)
	option.Scan()

	switch option.Text() {
	case "1":
		fmt.Print("\n")
		tc.TemperatureConverter()
	case "2":
		fmt.Print("\n")
		hc.HeightConverter()
	case "0":
		fmt.Println("Saindo do sistema...")
		os.Exit(1)
	default:
		fmt.Println("Comando invalido")
	}

}

func convertPesToMeters() {
	fmt.Println("Conversor de Pes para Metros")
	fmt.Print("Insira a altura em pes: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	heightInFoot, _ := strconv.ParseFloat(scanner.Text(), 64)

	heightInMeters := heightInFoot * 0.3048
	fmt.Printf("Altura em metros equivale a: %.2f metros", heightInMeters)
	fmt.Print("\n")
}
