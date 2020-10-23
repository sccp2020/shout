package impl

import (
	"context"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sccp2020/shout/backend/model"
)

type UserRepository struct {
	db sqlx.DB
}

func (u UserRepository) GetUser(ctx context.Context, id string) (*model.User, error) {
	user := model.User{}
	err := u.db.GetContext(ctx,  &user, "SELECT * FROM users WHERE users.id = $1", id)
	return &user, err
}

func (u UserRepository) CreateUser(ctx context.Context, user *model.User) error {
	sql := `INSERT INTO users (
				id,
				screen_id,
				name,
				hash,
				biography,
				profile_image_url,
				created_at,
				updated_at
			) VALUES (
				:id,
				:screen_id,
				:name,
				:hash,
				:biography,
				:profile_image_url,
				:created_at,
				:updated_at
			)`

	_, err := u.db.NamedExecContext(ctx, sql, user)
	return err
}

func (u UserRepository) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	panic("implement me")
}

func (u UserRepository) DeleteUser(ctx context.Context, id string) error {
	panic("implement me")
}

func (u UserRepository) GetFollowers(ctx context.Context, id string, cursor string, limit int64) ([]*model.User, error) {
	panic("implement me")
}

func (u UserRepository) GetFollowees(ctx context.Context, id string, cursor string, limit int64) ([]*model.User, error) {
	panic("implement me")
}

func (u UserRepository) FollowUser(ctx context.Context, id string, target string) error {
	panic("implement me")
}

func (u UserRepository) UnFollowUser(ctx context.Context, id string, target string) error {
	panic("implement me")
}
