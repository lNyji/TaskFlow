package note

import (
	"TaskFlow/internal/storage"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func Create(title, description, content string) {
	db, err := storage.Connect()
	if err != nil {
		panic(err)
	}

	note := Note{
		Title:       title,
		Description: description,
		Content:     content,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := db.Create(&note).Error; err != nil {
		panic(err)
	}
}

func List() {
	db, err := storage.Connect()
	if err != nil {
		panic(err)
	}

	var notes []Note

	if err := db.Order("id asc").Find(&notes).Error; err != nil {
		panic(err)
	}

	for _, t := range notes {
		fmt.Printf("ID: %d | %s | %s | %s \n",
			t.ID, t.Title, t.Description, t.Content)
	}
}

func Update() {
	db, err := storage.Connect()
	if err != nil {
		panic(err)
	}

	List()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("\nDigite o ID da anotação que deseja atualizar: ")
	var id uint
	fmt.Scan(&id)

	var note Note

	if err := db.First(&note, id).Error; err != nil {
		fmt.Println("Anotação não encontrada!")
		return
	}

	fmt.Printf("Novo título (%s): ", note.Title)
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Printf("Nova descrição (%s): ", note.Description)
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	fmt.Printf("Novo conteúdo [%s]: ", note.Content)
	content, _ := reader.ReadString('\n')
	content = strings.TrimSpace(strings.ToLower(content))

	if title != "" {
		note.Title = title
	}

	if description != "" {
		note.Description = description
	}

	if content != "" {
		note.UpdatedAt = time.Now()
	}
	if err := db.Save(&note).Error; err != nil {
		fmt.Println("Erro ao atualizar anotação:", err)
		return
	}

	fmt.Println("Anotação atualizada com sucesso!")
}

func Delete() {
	db, err := storage.Connect()
	if err != nil {
		panic(err)
	}

	List()

	fmt.Print("\nDigite o ID da Anotação que deseja remover: ")
	var id uint
	fmt.Scan(&id)

	var note Note

	if err := db.First(&note, id).Error; err != nil {
		fmt.Println("Anotação não encontrada!")
		return
	}

	if err := db.Delete(&note, id).Error; err != nil {
		fmt.Println("Erro ao remover a Anotação")
		return
	}
}
