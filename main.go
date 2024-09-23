package main

import (
	"github.com/bookkyjung1221/echo-clean-architecture/controller"
	"github.com/bookkyjung1221/echo-clean-architecture/db"
	"github.com/bookkyjung1221/echo-clean-architecture/repository"
	"github.com/bookkyjung1221/echo-clean-architecture/router"
	"github.com/bookkyjung1221/echo-clean-architecture/usecase"
	"github.com/bookkyjung1221/echo-clean-architecture/validator"
)

func main() {

	db := db.NewDB()

	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()

	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)

	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)

	e := router.NewRouter(userController, taskController)

	e.Logger.Fatal(e.Start(":8080"))

}
