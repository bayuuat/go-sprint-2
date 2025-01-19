package api

import (
	"context"
	"net/http"
	"time"

	"github.com/bayuuat/go-sprint-2/dto"
	"github.com/bayuuat/go-sprint-2/internal/middleware"
	"github.com/bayuuat/go-sprint-2/internal/service"
	"github.com/bayuuat/go-sprint-2/internal/utils"
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
	return ctx.Status(200).JSON(fiber.Map{})
}

func (da activityApi) CreateActivity(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	var req dto.ActivityReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(dto.NewErrorResponse("Invalid request:" + err.Error()))
	}

	fails := utils.Validate(req)
	if len(fails) > 0 {
		var errMsg string
		for field, err := range fails {
			errMsg += field + ": " + err + "; "
		}
		return ctx.Status(http.StatusBadRequest).JSON(dto.NewErrorResponse("Validation error:  " + errMsg))
	}

	res, code, err := da.activityService.CreateActivity(c, req)
	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(http.StatusCreated).JSON(res)
}

func (da activityApi) UpdateActivity(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{})
}

func (da activityApi) DeleteActivity(ctx *fiber.Ctx) error {
	return ctx.Status(200).JSON(fiber.Map{})
}
