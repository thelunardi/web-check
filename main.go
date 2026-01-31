package main

import (
	"fmt"
	"os"
)

func main() {
	start()
	showMenu()
	command := getOption()

	switch command {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo logs...")
	case 0:
		fmt.Println("Saindo...")
		os.Exit(0)
	default:
		fmt.Println("Comando inexistente")
		os.Exit(-1)
	}
}

func start() {
	name := "Alexandre"
	age := 33
	version := 1.1
	fmt.Println("Olá sr.", name, "sua idade é", age)
	fmt.Println("Este programa está na versão", version)
}

func showMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func getOption() int {
	var command int
	fmt.Scan(&command)
	fmt.Println("O comando escolhido foi", command)

	return command
}
