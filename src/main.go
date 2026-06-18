package main

import (
	"TaskFlow/internal/note"
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
	db.AutoMigrate(
		&task.Task{},
		&note.Note{},
	)

	reader := bufio.NewReader(os.Stdin)

	for {
		optionMenu := ui.ShowMenu()

		switch optionMenu {
		case 1:
			optionNote := ui.MenuNotes()

			switch optionNote {
			case 1:
				fmt.Println("\n=== Nova nota ===")
				fmt.Print("Título: ")
				title, _ := reader.ReadString('\n')
				title = strings.TrimSpace(title)

				fmt.Print("Descrição (opcional): ")
				description, _ := reader.ReadString('\n')
				description = strings.TrimSpace(description)

				fmt.Print("Conteúdo: ")
				content, _ := reader.ReadString('\n')
				content = strings.TrimSpace(content)

				note.Create(title, description, content)
				fmt.Println("===================")
			case 2:
				note.List()
			case 3:
				note.Update()
			case 4:
				note.Delete()
			case 0:
				fmt.Println("Saindo...")
				return

			default:
				fmt.Println("Opção inválida!")
			}

		case 2:
			optionTask := ui.MenuTasks()

			switch optionTask {
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
		case 0:
			return
		default:
			fmt.Println("Opção inválida")
		}
	}
}
