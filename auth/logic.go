package auth

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
)

type service struct {
	repo Repository
}

//NewService ...
func NewService(repo Repository) Service {
	return &service{
		repo: repo,
	}
}
func (s service) Login(ctx context.Context, req LoginRequest) (string, error) {
	user := User{
		Email:    req.Email,
		Password: req.Password,
	}
	err := s.repo.Login(ctx, user)
	if err != nil {
		return "error logging in", err
	}
	jwtToken, err := generateJWT()
	return jwtToken, err
}
func (s service) Register(ctx context.Context, req RegisterRequest) (string, error) {
	uuid, _ := uuid.NewV4()
	id := uuid.String()

	user := User{
		ID:       id,
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	err := s.repo.Register(ctx, user)
	if err != nil {
		return "failed to create user", err
	}
	return "Success", nil
}

func (s service) Token(ctx context.Context, req TokenRequest) (string, error) {
	return generateJWT()
}

func (s service) GetUser(ctx context.Context, id string) (string, error) {
	email, err := s.repo.GetUser(ctx, id)

	if err != nil {
		return "", err
	}

	return email, nil
}

func generateJWT() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "auth",
		"sub": "login",
		"aud": "any",
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})
	jwtToken, err := token.SignedString([]byte("SUPER_SECRET_KEY")) //could just return this func but this is more readable
	return jwtToken, err
}
