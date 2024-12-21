package internal

import (
	"context"
	"fmt"
	"gorm/dao/query"
	"gorm/model"
	"log"
)

func RunSave() {
	students := &model.Students{
		Name:   "hello world",
		Age:    20,
		Course: model.Courses{"english", "math", "computer"},
	}

	err := query.Q.Students.WithContext(context.Background()).Save(students)
	if err != nil {
		log.Fatal("save error")
		return
	}
}

func RunFind(id uint) {
	student, err := query.Students.WithContext(context.Background()).Where(query.Students.ID.Eq(id)).First()
	if err != nil {
		log.Fatal("find error")
	}
	fmt.Println(student)
}
