package heightConverter

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func HeightConverter() {
	fmt.Println("Certo converter altura.")

	fmt.Print("\n")

	fmt.Println("Qual conversão você quer fazer?")

	fmt.Println("1 - Converter Metros para Pes:")
	fmt.Println("2 - Converter Pes para Metros:")

	fmt.Print("Selecione uma opção: ")
	option := bufio.NewScanner(os.Stdin)
	option.Scan()

	switch option.Text() {
	case "1":
		fmt.Println("1 - Converter Metros para Pes:")
		convertMetersToFeet()
	case "2":
		fmt.Println("2 - Converter Pes para Metros:")
		convertFeetToMeters()
	}
}

func convertMetersToFeet() {
	fmt.Print("Insira a altura em metros: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	heightInMeters, _ := strconv.ParseFloat(scanner.Text(), 64)
	heightInFeet := heightInMeters * 3.281
	fmt.Printf("Altura em pes equivale a: %.2f pes", heightInFeet)
	fmt.Print("\n")
}

func convertFeetToMeters() {
	fmt.Print("Insira a altura em pes: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	heightInFeet, _ := strconv.ParseFloat(scanner.Text(), 64)
	heightInMeters := heightInFeet * 0.3048
	fmt.Printf("Altura em metros equivale a: %.2f metros", heightInMeters)
	fmt.Print("\n")
}
