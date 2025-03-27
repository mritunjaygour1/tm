package main

import (
	"log"
	"net/http"
	"regexp"
	"task-manager/internal/model"
	"task-manager/internal/router"
	"task-manager/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var (
	nameValidatePattern        = regexp.MustCompile(`^[A-Za-zÀ-ÿ]+([ -][A-Za-zÀ-ÿ]+)*$`)
	phoneNumberValidatePattern = regexp.MustCompile(`^\d{10}$`)
)

func main() {
	// db, err := gorm.Open("postgres", "postgres:password@/task_manager?charset=utf8&parseTime=True&loc=Local&sslmode=disable")
	db, err := gorm.Open("postgres", "user=postgres dbname=task_manager sslmode=disable host=localhost port=5432 password=password")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	db.AutoMigrate(&model.Task{})

	// taskService := &service.TaskService{DB: db}
	taskService := service.NewTaskService(db)

	r := gin.Default()

	// Register the custom validation function
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("name", NameValidator)   // for Name regex validation
		v.RegisterValidation("phone", PhoneValidator) // for Phone regex validation
	}

	router.SetupRouter(r, taskService)

	log.Fatal(http.ListenAndServe(":8080", r))
}

// NameValidator is a custom validation function for the 'otp' tag
func NameValidator(fl validator.FieldLevel) bool {
	name := fl.Field().String()
	return nameValidatePattern.MatchString(name)
}

// PhoneValidator is a custom validation function for the 'phone' tag
func PhoneValidator(fl validator.FieldLevel) bool {
	phone := fl.Field().String()
	return phoneNumberValidatePattern.MatchString(phone)
}
