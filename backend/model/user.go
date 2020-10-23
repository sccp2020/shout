package model

import "time"

type User struct {
	ID string `db:"id"`
	ScreenID string `db:"screen_id"`
	Name string `db:"name"`
	Hash string	`db:"hash"`
	Biography string `db:"biography"`
	ProfileImageUrl string `db:"profile_image_url"`
	CreatedAt *time.Time `db:"created_at"`
	UpdatedAt *time.Time `db:"updated_at"`
}
