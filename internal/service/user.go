package service

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/bayuuat/go-sprint-2/domain"
	"github.com/bayuuat/go-sprint-2/dto"
	"github.com/bayuuat/go-sprint-2/internal/config"
	"github.com/bayuuat/go-sprint-2/internal/repository"
	"github.com/bayuuat/go-sprint-2/internal/utils"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, req dto.AuthReq) (dto.AuthResponse, int, error)
	Login(ctx context.Context, req dto.AuthReq) (dto.AuthResponse, int, error)
	GetUser(ctx context.Context, email string) (dto.UserPreferences, int, error)
	PatchUser(ctx context.Context, req dto.UpdateUserPreferences, email string) (dto.UserPreferences, int, error)
}

type userService struct {
	cnf            *config.Config
	userRepository repository.UserRepository
}

func NewUser(cnf *config.Config,
	userRepository repository.UserRepository) UserService {
	return &userService{
		cnf:            cnf,
		userRepository: userRepository,
	}
}

func (a userService) Register(ctx context.Context, req dto.AuthReq) (dto.AuthResponse, int, error) {
	user, err := a.userRepository.FindByEmail(ctx, req.Email)
	if err != nil && err != sql.ErrNoRows {
		slog.ErrorContext(ctx, err.Error())
		return dto.AuthResponse{}, http.StatusInternalServerError, err
	}

	if user.Id != "" {
		return dto.AuthResponse{}, http.StatusConflict, domain.ErrEmailExists
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.AuthResponse{}, http.StatusInternalServerError, err
	}

	newUser := domain.User{
		Id:       uuid.New().String(),
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	err = a.userRepository.Save(ctx, &newUser)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.AuthResponse{}, http.StatusInternalServerError, err
	}

	user = newUser

	token, err := utils.GenerateToken(user)

	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.AuthResponse{}, http.StatusInternalServerError, err
	}

	return dto.AuthResponse{
		Email: user.Email,
		Token: token,
	}, http.StatusCreated, nil
}

func (a userService) Login(ctx context.Context, req dto.AuthReq) (dto.AuthResponse, int, error) {
	user, err := a.userRepository.FindByEmail(ctx, req.Email)
	if err != nil && err != sql.ErrNoRows {
		slog.ErrorContext(ctx, err.Error())
		return dto.AuthResponse{}, http.StatusInternalServerError, err
	}

	if user.Id == "" {
		return dto.AuthResponse{}, http.StatusNotFound, domain.ErrNotFound
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return dto.AuthResponse{}, http.StatusUnauthorized, domain.ErrInvalidCredential
	}

	token, err := utils.GenerateToken(user)

	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.AuthResponse{}, http.StatusInternalServerError, err
	}

	return dto.AuthResponse{
		Email: user.Email,
		Token: token,
	}, http.StatusOK, nil
}

func (a userService) GetUser(ctx context.Context, email string) (dto.UserPreferences, int, error) {
	fmt.Print(email)
	user, err := a.userRepository.FindByEmail(ctx, email)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.UserPreferences{}, http.StatusInternalServerError, err
	}

	return dto.UserPreferences{
		Email:      &user.Email,
		Name:       user.Name,
		ImageUri:   user.ImageURI,
		Preference: user.Preference,
		WeightUnit: user.WeightUnit,
		HeightUnit: user.HeightUnit,
		Weight:     floatPtrToIntPtr(user.Weight),
		Height:     floatPtrToIntPtr(user.Height),
	}, http.StatusOK, nil
}

func (a userService) PatchUser(ctx context.Context, req dto.UpdateUserPreferences, id string) (dto.UserPreferences, int, error) {
	user, err := a.userRepository.FindById(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.UserPreferences{}, http.StatusInternalServerError, err
	}

	if user.Email == "" {
		return dto.UserPreferences{}, http.StatusNotFound, domain.ErrUserNotFound
	}

	if req.Name != nil {
		user.Name = req.Name
	}
	if req.ImageUri != nil {
		user.ImageURI = req.ImageUri
	}

	user.Preference = req.Preference
	user.WeightUnit = req.WeightUnit
	user.HeightUnit = req.HeightUnit
	user.Weight = req.Weight
	user.Height = req.Height
	user.Name = req.Name

	err = a.userRepository.Update(ctx, &user)
	if err != nil {
		slog.ErrorContext(ctx, err.Error())
		return dto.UserPreferences{}, http.StatusInternalServerError, err
	}

	return dto.UserPreferences{
		Email:      &user.Email,
		Name:       user.Name,
		ImageUri:   user.ImageURI,
		Preference: user.Preference,
		WeightUnit: user.WeightUnit,
		HeightUnit: user.HeightUnit,
		Weight:     floatPtrToIntPtr(user.Weight),
		Height:     floatPtrToIntPtr(user.Height),
	}, 200, nil
}

func floatPtrToIntPtr(floatPtr *float64) *int {
	if floatPtr == nil {
		return nil
	}

	result := int(*floatPtr)
	return &result
}
