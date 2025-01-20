package api

import (
	"context"
	"fmt"
	"github.com/bayuuat/go-sprint-2/dto"
	"github.com/bayuuat/go-sprint-2/internal/middleware"
	"github.com/bayuuat/go-sprint-2/internal/service"
	"github.com/bayuuat/go-sprint-2/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"time"
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

	user := ctx.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

	res, code, err := da.activityService.GetActivitysWithFilter(c, filter, userId)
	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(code).JSON(res)
}

func (da activityApi) CreateActivity(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	user := ctx.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)

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

	res, code, err := da.activityService.CreateActivity(c, req, userId)
	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(code).JSON(res)
}

func (da activityApi) UpdateActivity(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	user := ctx.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)
	id := ctx.Params("id")

	var req dto.UpdateActivityReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(http.StatusBadRequest)
	}

	if err := utils.Validate(req); err != nil {
		fmt.Println(err)
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err})
	}

	res, code, err := da.activityService.PatchActivity(c, req, userId, id)

	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(code).JSON(res)
}

func (da activityApi) DeleteActivity(ctx *fiber.Ctx) error {
	c, cancel := context.WithTimeout(ctx.Context(), 10*time.Second)
	defer cancel()

	user := ctx.Locals("jwt").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["id"].(string)
	id := ctx.Params("id")

	res, code, err := da.activityService.DeleteActivity(c, userId, id)

	if err != nil {
		return ctx.Status(code).JSON(dto.ErrorResponse{Message: err.Error()})
	}

	return ctx.Status(code).JSON(res)
}
