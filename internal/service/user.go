package service

import (
	"context"
	"database/sql"
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
	GetUser(ctx context.Context, email string) (dto.UserData, int, error)
	PatchUser(ctx context.Context, req dto.UpdateUserReq, email string) (dto.UserData, int, error)
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

func (a userService) GetUser(ctx context.Context, email string) (dto.UserData, int, error) {
	return dto.UserData{}, 400, nil
}

func (a userService) PatchUser(ctx context.Context, req dto.UpdateUserReq, id string) (dto.UserData, int, error) {
	return dto.UserData{}, 400, nil
}
