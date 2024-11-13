package validate

import (
	"github.com/go-playground/validator/v10"
	"log"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
	validate.SetTagName("v")
}

func outRes(tag string, err *error) {
	log.Println("----------start " + tag + "----------")
	log.Println(*err)
	log.Println("----------end " + tag + "----------")
	err = nil
}
