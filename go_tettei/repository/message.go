package repository

import (
	"context"
	"database/sql"
	"fmt"
)

type MessageRepository interface {
	Delete(ctx context.Context, id string) error
}

type messageRepository struct {
	db *sql.DB
}

func NewMessageRepository(db *sql.DB) MessageRepository {
	return &messageRepository{db}
}

func (r *messageRepository) Delete(ctx context.Context, id string) error {
	db, ok := GetTx(ctx)
	if !ok {
		return fmt.Errorf("transaction store doesn't exist")
	}
	fmt.Println("start delete message rep")
	query := "DELETE FROM message WHERE user_id = $1"
	result, err := db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowAffected != 1 {
		fmt.Printf("expected 1 row to be affected, got %d, id %s, query %s", rowAffected, id, query)
		return fmt.Errorf("expected 1 row to be affected, got %d, id %s, query %s", rowAffected, id, query)
	}
	fmt.Printf("message %s deleted: %d\n", id, rowAffected)
	return nil
}
