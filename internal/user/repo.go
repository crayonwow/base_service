package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type (
	repo struct {
		db *sql.DB
	}
)

func newRepo(db *sql.DB) *repo {
	return &repo{db: db}
}

const (
	selectUserByID = `SELECT id, name, date_of_birth FROM users where id=?`
	insertUser     = `INSERT INTO users (name, date_of_birth) VALUES  (?, ?)`
	updateUserByID = `UPDATE users SET name=?, date_of_birth=? WHERE id=?`
	deleteUserByID = `DELETE FROM users WHERE id=?`
)

func (r *repo) Create(ctx context.Context, user User) (int64, error) {
	res, err := r.db.ExecContext(ctx, insertUser, user.Name, user.DateOfBirth)
	if err != nil {
		return 0, fmt.Errorf("exec context: %w", err)
	}
	return res.LastInsertId()
}

func (r *repo) Get(ctx context.Context, userID int64) (User, error) {
	user := User{}
	err := r.db.QueryRowContext(ctx, selectUserByID, userID).Scan(&user.ID, &user.Name, &user.DateOfBirth)
	if errors.Is(err, sql.ErrNoRows) {
		return user, ErrUserNotFound
	}

	return user, err
}

func (r *repo) Update(ctx context.Context, user User) error {
	_, err := r.db.ExecContext(ctx, updateUserByID, user.Name, user.DateOfBirth, user.ID)
	return err
}

func (r *repo) Delete(ctx context.Context, userID int64) error {
	_, err := r.db.ExecContext(ctx, deleteUserByID, userID)
	return err
}
