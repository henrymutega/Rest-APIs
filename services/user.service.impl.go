package services

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log.com/log-apis/models"
)

type UserServiceImpl struct {
	usercollection *mongo.Collection
	ctx            context.Context
}

// CreateUser implements UserService.
func (u *UserServiceImpl) CreateUser(user *models.User) error {
	_, err := u.usercollection.InsertOne(u.ctx, user)
	return err
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
func (u *UserServiceImpl) GetUser(name *string) (*models.User, error) {
	var user *models.User
	query := bson.D{bson.E{Key: "name", Value: name}}
	err := u.usercollection.FindOne(u.ctx, query).Decode(&user)
	return user, err
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
	return &UserServiceImpl{
		usercollection: usercollection,
		ctx:            ctx,
	}
}
