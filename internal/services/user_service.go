package services

import (
	"context"
	"errors"
	"time"

	"github.com/beatrizrdgs/literalog/internal/models"
	"github.com/golang-jwt/jwt"
	"github.com/literalog/go-wise/wise"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo wise.MongoRepository[models.User]
}

func NewUserService(repo wise.MongoRepository[models.User]) (*UserService, error) {
	if repo == nil {
		return nil, errors.New("repository is required")
	}
	return &UserService{repo: repo}, nil
}

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func (s *UserService) RegisterUser(ctx context.Context, user models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	return s.repo.InsertOne(ctx, user)
}

func (s *UserService) LoginUser(ctx context.Context, email, password string) (string, error) {
	filters := map[string][]any{"email": {email}}

	user, err := s.repo.FindOne(ctx, filters)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	expirationTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}
