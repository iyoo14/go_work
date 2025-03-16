package repository

import (
	"context"
	"database/sql"
	"fmt"
	"go_work/go_tettei/model"
	"strconv"
)

type UserRepository interface {
	Create(ctx context.Context, user *model.User) (id string, err error)
	Read(ctx context.Context, id string) (*model.User, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(ctx context.Context, user *model.User) (string, error) {
	stmt, err := r.db.Prepare("INSERT INTO users (name, email, age) values ($1, $2, $3) RETURNING id")
	if err != nil {
		return "", fmt.Errorf("query prepare faild: %w", err)
	}
	var id int
	err = stmt.QueryRow(user.Name, user.Email, user.Age).Scan(&id)
	if err != nil {
		return "", fmt.Errorf("query faild: %w", err)
	}
	idStr := strconv.FormatInt(int64(id), 10)
	return idStr, nil
}

func (r *userRepository) Read(ctx context.Context, id string) (*model.User, error) {
	var user model.User
	err := r.db.QueryRow("SELECT id, name, email, age FROM users WHERE id = $1", id).Scan(&user.ID, &user.Name, &user.Email, &user.Age)
	return &user, err
}

func (r *userRepository) Update(ctx context.Context, user *model.User) error {
	result, err := r.db.Exec("UPDATE users SET name = $1, age = $2 WHERE id = $3", user.Name, user.Age, user.ID)
	if err != nil {
		return err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected != 1 {
		return fmt.Errorf("expected 1 row to be affected, got %d, id %s", rowAffected, user.ID)
	}
	return nil
}

func (r *userRepository) Delete(ctx context.Context, id string) error {
	db, ok := GetTx(ctx)
	if !ok {
		return fmt.Errorf("transaction store doesn't exist")
	}
	fmt.Println("start delete user rep")
	query := "DELETE FROM users WHERE id = $1"
	result, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("query ExecContext faild: %w", err)
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected != 1 {
		fmt.Printf("expected 1 row to be affected, got %d, id %s, query %s", rowAffected, id, query)
		//return fmt.Errorf("expected 1 row to be affected, got %d, id %s, query %s", rowAffected, id, query)
	}
	fmt.Printf("user %s deleted: %d\n", id, rowAffected)
	return nil
}
