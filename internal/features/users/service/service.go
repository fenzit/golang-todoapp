package users_service

import (
	"context"

	"github.com/fenzit/golang-todoapp/internal/core/domain"
)

type UsersService struct {
	usersReposity UsersRepository
}

type UsersRepository interface {
	CreateUser(
		ctx context.Context,
		user domain.User,
	) (domain.User, error)
}

func NewUsersService(
	usersRepository UsersRepository,
) *UsersService {
	return &UsersService{
		usersReposity: usersRepository,
	}
}

