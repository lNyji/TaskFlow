package ui

import (
	"TaskFlow/internal/task"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
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

func CreateTask() task.Task {

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n=== Nova Tarefa ===")
	fmt.Print("Título: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Print("Descrição (opcional): ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)
	fmt.Println("===================")

	return task.Task{
		Title:       title,
		Description: description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func ListTasks(tasks []task.Task) {
	fmt.Println("\n=== Lista de Tarefas ===")

	for _, t := range tasks {
		fmt.Printf("ID: %d | %s | %s | done: %t\n",
			t.ID, t.Title, t.Description, t.Completed)
	}
}

func UpdateTask(tasks []task.Task) task.UpdateTask {
	ListTasks(tasks)

	reader := bufio.NewReader(os.Stdin)

	var updateTask task.UpdateTask

	fmt.Print("\nDigite o ID da tarefa que deseja atualizar: ")
	fmt.Scan(&updateTask.ID)

	fmt.Print("Novo título: ")
	updateTask.Title, _ = reader.ReadString('\n')
	updateTask.Title = strings.TrimSpace(updateTask.Title)

	fmt.Print("Nova descrição: ")
	updateTask.Description, _ = reader.ReadString('\n')
	updateTask.Description = strings.TrimSpace(updateTask.Description)

	fmt.Print("Concluída? (s/n): ")
	updateTask.Completed, _ = reader.ReadString('\n')
	updateTask.Completed = strings.TrimSpace(strings.ToLower(updateTask.Completed))

	return updateTask
}

func DeleteTask(tasks []task.Task) (id uint) {

	ListTasks(tasks)
	fmt.Print("\nDigite o ID da tarefa que deseja remover: ")
	fmt.Scan(&id)

	return id
}
