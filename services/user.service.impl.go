package services

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"log.com/log-apis/models"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

// CreateUser implements UserService.
func (u UserServiceImpl) CreateUser(*models.User) error {
	panic("unimplemented")
}

// DeleteUser implements UserService.
func (u UserServiceImpl) DeleteUser(*string) error {
	panic("unimplemented")
}

// GetAll implements UserService.
func (u UserServiceImpl) GetAll() ([]*models.User, error) {
	panic("unimplemented")
}

// GetUser implements UserService.
func (u UserServiceImpl) GetUser(*string) (*models.User, error) {
	panic("unimplemented")
}

// UpdateUser implements UserService.
func (u UserServiceImpl) UpdateUser(*models.User) {
	panic("unimplemented")
}

// func (u *UserServiceImpl) CreateUser(user *models.User) error {
// 	return nil
// }

// func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {
// 	return nil, nil
// }

// func (u *UserServiceImpl) GetAll() ([]*models.User, error) {
// 	return nil, nil
// }

// func (u *UserServiceImpl) UpdateUser(user *models.User) error {
// 	return nil
// }

// func (u *UserServiceImpl) DeleteUser(name *string) (*models.User, error) {
// 	return nil, nil
// }

func NewUserService(usercollection *mongo.Collection, ctx context.Context) UserService {
	return UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}
