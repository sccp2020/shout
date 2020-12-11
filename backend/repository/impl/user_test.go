package impl

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"regexp"
	"testing"
	"time"
)

func TestGetUser(t *testing.T) {
	db, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	defer db.Close()

	mock.ExpectQuery(regexp.QuoteMeta(`SELECT * FROM users WHERE id = ?`)).
		WithArgs("xxxx").
		WillReturnRows(
			sqlmock.NewRows(
				[]string{
					"id",
					"screen_id",
					"name",
					"hash",
					"biography",
					"profile_image_url",
					"created_at",
					"updated_at",
				},
			).AddRow(
				"xxxx",
				"uzimaru0000",
				"uzimaru",
				"xxxxxxxxxx",
				"piyopiyo",
				"https://example.com/image.png",
				time.Now(),
				time.Now(),
			),
		)

	userRepo := UserRepository{
		db: sqlx.NewDb(db, "mock"),
	}

	user, err := userRepo.GetUser(context.Background(), "xxxx")
	if err != nil {
		t.Fatalf("%s", err)
	}

	if user.ID != "xxxx" {
		t.Fatalf("%v", user)
	}
}
