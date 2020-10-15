package model

type Shout struct {
	ID string
	UserID string
	Content string
	Parent string
	ReShoutUserID string
	CreatedAt int64
}
