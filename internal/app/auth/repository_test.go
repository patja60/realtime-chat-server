package auth

import (
	"database/sql"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	database "github.com/patja60/realtime-chat-server/pkg"
)

func TestCreateUser(t *testing.T) {
	t.Run("SuccessCreateUser", func(t *testing.T) {
		sqlDB, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer sqlDB.Close()

		db := database.DB{Conn: sqlDB}

		repo := NewAuthRepository(db)

		email := "user@example.com"
		hashedPassword := "hashedpassword"

		mock.ExpectExec("INSERT INTO users").
			WithArgs(email, hashedPassword).
			WillReturnResult(sqlmock.NewResult(1, 1))

		err = repo.CreateUser(email, hashedPassword)

		if err != nil {
			t.Errorf("expected no error, but got %v", err)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}

func GetUserByEmail(t *testing.T) {
	t.Run("SuccessGetUserByEmail", func(t *testing.T) {
		sqlDB, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer sqlDB.Close()

		db := database.DB{Conn: sqlDB}

		repo := NewAuthRepository(db)

		email := "user@example.com"
		hashedPassword := "hashedpassword"

		rows := sqlmock.NewRows([]string{"id", "email", "password_hash"}).
			AddRow("1", email, hashedPassword)

		mock.ExpectQuery("SELECT id, email, passwordHash FROM user WHERE email = ?").
			WithArgs(email).
			WillReturnRows(rows)

		user, err := repo.GetUserByEmail(email)

		if err != nil {
			t.Errorf("expected no error, but got %v", err)
		}

		if user.Email != email {
			t.Errorf("expected email to be %v, but got %v", email, user.Email)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("NotExistUser", func(t *testing.T) {
		sqlDB, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
		}
		defer sqlDB.Close()

		db := database.DB{Conn: sqlDB}

		repo := NewAuthRepository(db)

		email := "user@example.com"

		mock.ExpectQuery("SELECT id, email, passwordHash FROM user WHERE email = ?").
			WithArgs(email).
			WillReturnRows(nil)

		user, err := repo.GetUserByEmail(email)

		if err == nil {
			t.Errorf("expected an error, got nil")
		}

		if !errors.Is(err, sql.ErrNoRows) {
			t.Errorf("expected sql.ErrNoRows, got %v", err)
		}

		if user != nil {
			t.Errorf("expected user to be nil, got %v", user)
		}

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
