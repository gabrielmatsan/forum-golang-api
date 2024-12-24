package main

import (
	"fmt"

	"github.com/gabrielmatsan/forum-golang-api/internal/domain/forum/enterprise/models"
)

func main() {
	student := models.NewStudent(models.StudentProps{
		Name:     "Gabriel",
		Email:    "gabrielmatsan@hotmail.com",
		Password: "123456",
	})

	fmt.Println(student.GetID())
	fmt.Println(student.GetName())
	fmt.Println(student.GetEmail())
	fmt.Println(student.GetPassword())
	fmt.Println(student.ID().ToString())
}
