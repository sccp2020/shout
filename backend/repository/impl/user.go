package impl

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/sccp2020/shout/backend/model"
)

type UserRepository struct {
	db *sqlx.DB
}

func (u UserRepository) GetUser(ctx context.Context, id string) (*model.User, error) {
	user := model.User{}
	err := u.db.GetContext(ctx,  &user, "SELECT * FROM users WHERE id = ?", id)
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
			) VALUES (
				:id,
				:screen_id,
				:name,
				:hash,
				:biography,
				:profile_image_url,
			)`

	_, err := u.db.NamedExecContext(ctx, sql, user)
	return err
}

func (u UserRepository) UpdateUser(ctx context.Context, user *model.User) (*model.User, error) {
	sql := `UPDATE users
            SET screen_id = :screen_id,
				name = :name,
				hash = :hash,
				biography = :biography,
				profile_image_url = :profile_image_url,
			WHERE id = :id`

	_, err := u.db.NamedExecContext(ctx, sql, user)
	return user, err
}

func (u UserRepository) DeleteUser(ctx context.Context, id string) error {
	sql := `DELETE FROM users WHERE id = ?`

	_, err := u.db.ExecContext(ctx, sql, id)
	return err
}

func (u UserRepository) GetFollowers(ctx context.Context, id string, limit int64, offset int64) ([]*model.User, error) {
	sql := `SELECT u2.*
			FROM users u1
			INNER JOIN follows
			ON u1.id = follows.followee
			INNER JOIN users u2
			ON u2.id = follows.follower
			WHERE u1.id = ?
			LIMIT ?
			OFFSET ?`

	var followers []*model.User
	err := u.db.SelectContext(ctx, &followers, sql, id, limit, offset)
	return followers, err
}

func (u UserRepository) GetFollowees(ctx context.Context, id string, limit int64, offset int64) ([]*model.User, error) {
	sql := `SELECT u2.*
			FROM users u1
			INNER JOIN follows
			ON u1.id = follows.follower
			INNER JOIN users u2
			ON u2.id = follows.followee
			WHERE u1.id = ?
			LIMIT ?
			OFFSET ?`

	var followers []*model.User
	err := u.db.SelectContext(ctx, &followers, sql, id, limit, offset)
	return followers, err
}

func (u UserRepository) FollowUser(ctx context.Context, id string, target string) error {
	sql := `SELECT * FROM follows WHERE followee = ? AND follower = ?`

	ok, err := rowExists(u.db, sql, target, id)
	if err != nil {
		return err
	}

	if ok {
		sql = `UPDATE follow
			   SET is_mutual = true
			   WHERE followee = ? AND follower = ?`

		_, err = u.db.ExecContext(ctx, sql, target, id)

		if err != nil {
			return err
		}
	}

	sql = fmt.Sprintf(`INSERT INTO follows (
    			followee,
        		follower,
            	is_mutual
		   ) VALUES (
		        $1,
		    	$2,
		        %t
		   )`, ok)

	_, err = u.db.ExecContext(ctx, sql, id, target)

	return nil
}

func (u UserRepository) UnFollowUser(ctx context.Context, id string, target string) error {
	sql := `SELECT * FROM follows WHERE followee = ? AND follower = ?`

	ok, err := rowExists(u.db, sql, target, id)
	if err != nil {
		return err
	}

	if ok {
		sql = `UPDATE follow
			   SET is_mutual = false
			   WHERE followee = ? AND follower = ?`

		_, err = u.db.ExecContext(ctx, sql, target, id)

		if err != nil {
			return err
		}
	}

	sql = `DELETE FROM follows WHERE followee = ? AND follower = ?`

	_, err = u.db.ExecContext(ctx, sql, id, target)

	return err
}
