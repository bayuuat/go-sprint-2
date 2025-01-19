package main

import (
	"github.com/bayuuat/go-sprint-2/internal/api"
	"github.com/bayuuat/go-sprint-2/internal/config"
	"github.com/bayuuat/go-sprint-2/internal/connection"
	"github.com/bayuuat/go-sprint-2/internal/repository"
	"github.com/bayuuat/go-sprint-2/internal/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cnf := config.Get()
	app := fiber.New()
	dbConnection := connection.GetDatabase(cnf.Database)

	userRepository := repository.NewUser(dbConnection)
	authService := service.NewUser(cnf, userRepository)
	api.NewUser(app, authService)

	departmentRepository := repository.NewActivity(dbConnection)
	activityTypesRepository := repository.NewActivityType(dbConnection)
	departmentService := service.NewActivity(cnf, departmentRepository, activityTypesRepository)
	api.NewActivity(app, departmentService)

	api.NewAws(app)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)
}
