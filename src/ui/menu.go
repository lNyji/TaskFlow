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
