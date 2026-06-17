package ui

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ShowMenu() int {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n=== Menu ===")
	fmt.Println("1. Anotações")
	fmt.Println("2. Tarefas")
	fmt.Println("0. Sair")
	fmt.Print("\nEscolha uma opção: ")

	optionStr, err := reader.ReadString('\n')

	if err != nil {
		return -1
	}

	option, err := strconv.Atoi(strings.TrimSpace(optionStr))

	if err != nil {
		return -1
	}

	return option
}

func MenuNotes() int {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n=== Menu Anotações ===")
	fmt.Println("1. Criar Nova Anotação")
	fmt.Println("2. Ver Anotações")
	fmt.Println("3. Atualizar Anotação")
	fmt.Println("4. Remover Anotação")
	fmt.Println("0. Sair")
	fmt.Print("Escolha uma opção: ")

	optionStr, err := reader.ReadString('\n')

	if err != nil {
		return -1
	}

	option, err := strconv.Atoi(strings.TrimSpace(optionStr))

	if err != nil {
		return -1
	}

	return option
}

func MenuTasks() int {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n=== Menu Tarefas ===")
	fmt.Println("1. Criar Nova Tarefa")
	fmt.Println("2. Ver Tarefas")
	fmt.Println("3. Atualizar Tarefa")
	fmt.Println("4. Remover Tarefa")
	fmt.Println("0. Sair")
	fmt.Print("Escolha uma opção: ")

	optionStr, err := reader.ReadString('\n')

	if err != nil {
		return -1
	}

	option, err := strconv.Atoi(strings.TrimSpace(optionStr))

	if err != nil {
		return -1
	}

	return option
}
