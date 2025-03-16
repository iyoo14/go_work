package usecase

import (
	"context"
	"fmt"
	"go_work/go_tettei/model"
	"go_work/go_tettei/repository"
	"go_work/go_tettei/transaction"
)

type UserUsecase interface {
	GetByID(ctx context.Context, id string) (*model.User, error)
	Create(ctx context.Context, user *model.User) (string, error)
	Update(ctx context.Context, user *model.User) error
	Delete(ctx context.Context, id string) error
}

type userUsecase struct {
	ur repository.UserRepository
	mr repository.MessageRepository
	tr transaction.Transaction
}

func NewUserUsecase(ur repository.UserRepository, mr repository.MessageRepository, tr transaction.Transaction) UserUsecase {
	return &userUsecase{ur, mr, tr}
}

func (u *userUsecase) GetByID(ctx context.Context, id string) (*model.User, error) {
	user, err := u.ur.Read(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (u *userUsecase) Create(ctx context.Context, user *model.User) (string, error) {
	id, err := u.ur.Create(ctx, user)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (u *userUsecase) Update(ctx context.Context, user *model.User) error {
	err := u.ur.Update(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (u *userUsecase) Delete(ctx context.Context, id string) error {
	_, err := u.tr.DoInTx(ctx, func(ctx context.Context) (any, error) {
		fmt.Println("start delete user")
		err := u.ur.Delete(ctx, id)
		if err != nil {
			return nil, err
		}
		fmt.Println("start delete message")
		err = u.mr.Delete(ctx, id)
		if err != nil {
			return nil, err
		}
		return nil, nil
	})
	if err != nil {
		return err
	}

	return nil
}
