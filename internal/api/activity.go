package api

import (
	"context"
	"net/http"
	"time"

	"github.com/bayuuat/go-sprint-2/dto"
	"github.com/bayuuat/go-sprint-2/internal/middleware"
	"github.com/bayuuat/go-sprint-2/internal/service"
	"github.com/gofiber/fiber/v2"
)

type activityApi struct {
	activityService service.ActivityService
}

func NewActivity(app *fiber.App,
	activityService service.ActivityService) {

	da := activityApi{
		activityService: activityService,
	}

	user := app.Group("/v1/activity")

	user.Use(middleware.JWTProtected)
	user.Post("/", da.CreateActivity)
	user.Get("/", da.GetActivitys)
	user.Patch("/:id?", da.UpdateActivity)
	user.Delete("/:id?", da.DeleteActivity)
}

func (da activityApi) GetActivitys(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var filter dto.ActivityFilter
	if err := ctx.QueryParser(&filter); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResponse{Message: "Invalid query parameters"})
	}

	if filter.Limit == 0 {
		filter.Limit = 5
	}
	if filter.Offset == 0 {
		filter.Offset = 0
	}

	res, code, err := da.activityService.GetActivitysWithFilter(c, filter)
	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(code).JSON(res)
}

func (da activityApi) CreateActivity(ctx *fiber.Ctx) error {
	return ctx.Status(201).JSON(fiber.Map{})
}

func (da activityApi) UpdateActivity(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{})
}

func (da activityApi) DeleteActivity(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{})
}
