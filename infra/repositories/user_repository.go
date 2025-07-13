package repositories

import (
	"context"
	"database/sql"
)

type User struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserRepository struct {
	tx *sql.Tx
}

func NewUserRepository(tx *sql.Tx) *UserRepository {
	return &UserRepository{tx: tx}
}

func (u *UserRepository) Create(ctx context.Context, user *User) (*User, error) {
	_, err := u.tx.ExecContext(
		ctx,
		"INSERT INTO users (id, name, email, password) VALUES ($1, $2, $3, $4)",
		user.Id, user.Name, user.Email, user.Password,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) FindById(ctx context.Context, id string) (*User, error) {
	var user User
	err := u.tx.QueryRowContext(ctx,
		"SELECT id, name, email, password, created_at, updated_at FROM users WHERE id = $1",
		id,
	).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	err := u.tx.QueryRowContext(
		ctx,
		"SELECT id, name, email, password, created_at, updated_at FROM users WHERE email = $1",
		email,
	).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) Update(ctx context.Context, user *User) error {
	_, err := u.tx.ExecContext(
		ctx,
		"UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4",
		user.Name, user.Email, user.Password, user.Id,
	)
	return err
}

func (u *UserRepository) Delete(ctx context.Context, id string) error {
	_, err := u.tx.ExecContext(ctx, "DELETE FROM users WHERE id = $1", id)
	return err
}
