package main

import (
	"TaskFlow/internal/storage"
	"TaskFlow/internal/task"
	"TaskFlow/ui"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	Menu()
}

func Menu() {
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

			CreateTask(title, description)
			fmt.Println("===================")

		case 2:
			fmt.Println("\n=== Lista de Tarefas ===")
			ListTasks()

		case 3:
			UpdateTask()

		case 4:
			DeleteTask()

		case 0:
			fmt.Println("Saindo...")
			return

		default:
			fmt.Println("Opção inválida!")
		}
	}
}

func CreateTask(title, description string) {
	db, err := storage.Connect()
	if err != nil {
		panic(err)
	}

	task := task.Task{
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := db.Create(&task).Error; err != nil {
		panic(err)
	}
}

func ListTasks() {
	db, err := storage.Connect()
	if err != nil {
		panic(err)
	}

	var tasks []task.Task

	if err := db.Order("id asc").Find(&tasks).Error; err != nil {
		panic(err)
	}

	for _, t := range tasks {
		fmt.Printf("ID: %d | %s | %s | done: %t\n",
			t.ID, t.Title, t.Description, t.Completed)
	}
}

func UpdateTask() {
	db, err := storage.Connect()
	if err != nil {
		panic(err)
	}

	ListTasks()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\nDigite o ID da tarefa que deseja atualizar: ")
	var id uint
	fmt.Scan(&id)

	var task task.Task

	if err := db.First(&task, id).Error; err != nil {
		fmt.Println("Tarefa não encontrada!")
		return
	}

	fmt.Printf("Novo título (%s): ", task.Title)
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Printf("Nova descrição (%s): ", task.Description)
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	fmt.Printf("Concluída? (s/n) [%t]: ", task.Completed)
	completedStr, _ := reader.ReadString('\n')
	completedStr = strings.TrimSpace(strings.ToLower(completedStr))

	if title != "" {
		task.Title = title
	}

	if description != "" {
		task.Description = description
	}

	switch completedStr {
	case "s":
		task.Completed = true
	case "n":
		task.Completed = false
	}

	task.UpdatedAt = time.Now()

	if err := db.Save(&task).Error; err != nil {
		fmt.Println("Erro ao atualizar tarefa:", err)
		return
	}

	fmt.Println("Tarefa atualizada com sucesso!")
}

func DeleteTask() {
	db, err := storage.Connect()
	if err != nil {
		panic(err)
	}

	ListTasks()

	fmt.Print("\nDigite o ID da tarefa que deseja remover: ")
	var id uint
	fmt.Scan(&id)

	var task task.Task

	if err := db.First(&task, id).Error; err != nil {
		fmt.Println("Tarefa não encontrada!")
		return
	}

	if err := db.Delete(&task, id).Error; err != nil {
		fmt.Println("Erro ao remover a tarefa")
		return
	}
}
