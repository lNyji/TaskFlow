package main

import (
	"TaskFlow/internal/storage"
	"TaskFlow/internal/task"
	"TaskFlow/ui"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	db, err := storage.Connect()
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&task.Task{})

	reader := bufio.NewReader(os.Stdin)

	for {
		choiceStr := ui.ShowMenu()

		switch choiceStr {
		case 1:
			fmt.Println("\n=== Nova Tarefa ===")
			fmt.Print("Título: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			fmt.Print("Descrição (opcional): ")
			description, _ := reader.ReadString('\n')
			description = strings.TrimSpace(description)

			task.Create(title, description)
			fmt.Println("===================")

		case 2:
			fmt.Println("\n=== Lista de Tarefas ===")
			task.List()

		case 3:
			task.Update()

		case 4:
			task.Delete()

		case 0:
			fmt.Println("Saindo...")
			return

		default:
			fmt.Println("Opção inválida!")
		}
	}
}
