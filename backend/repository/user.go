package repository

import (
	"context"
	"github.com/sccp2020/shout/backend/model"
)

type UserRepository interface {
	GetUser(ctx context.Context, id string) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) error
	UpdateUser(ctx context.Context, user *model.User) (*model.User, error)
	DeleteUser(ctx context.Context, id string) error
	GetFollowers(ctx context.Context, id string, cursor string, limit int64) ([]*model.User, error)
	GetFollowees(ctx context.Context, id string, cursor string, limit int64) ([]*model.User, error)
	FollowUser(ctx context.Context, id string, target string) error
	UnFollowUser(ctx context.Context, id string, target string) error
}
