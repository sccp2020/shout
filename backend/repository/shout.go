package repository

import (
	"context"
	"github.com/sccp2020/shout/backend/model"
)

type ShoutRepository interface {
	GetTimeline(ctx context.Context, id string, cursor string, limit int64) ([]*model.Shout, error)
	GetUserShouts(ctx context.Context, id string, cursor string, limit int64) ([]*model.Shout, error)
	CreateShout(ctx context.Context, shout *model.Shout) error
	DeleteShout(ctx context.Context, id string) error
	ReShout(ctx context.Context, id string, reShouterID string) error
	GetReShoutUsers(ctx context.Context, id string) ([]*model.User, error)
	DeleteReShout(ctx context.Context, id string, reShouterID string) error
	Like(ctx context.Context, id string, reShouterID string) error
	GetLikeUsers(ctx context.Context, id string) ([]*model.User, error)
	DeleteLike(ctx context.Context, id string, reShouterID string) error
}
