package main

import (
	"github.com/bookkyjung1221/echo-clean-architecture/db"
	"github.com/bookkyjung1221/echo-clean-architecture/repository"
	"github.com/bookkyjung1221/echo-clean-architecture/usecase"
	"github.com/bookkyjung1221/echo-clean-architecture/validator"
)

func main() {

	db := db.NewDB()

	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()

	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)

	userUsecase := usecase.NewTaskUsecase()

}
