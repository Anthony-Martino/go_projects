package auth

import (
	"context"

	"github.com/go-kivik/kivik"
	"golang.org/x/crypto/bcrypt"
)

//Repository consists of all abstract methods which interfaces with the database implementation
type Repository interface {
	Register(ctx context.Context, user User) error
	GetUser(ctx context.Context, id string) (string, error)
}

type repository struct {
	db *kivik.DB
}

//NewRepository ...
func NewRepository(db *kivik.DB) (Repository, error) {
	return &repository{
		db: db,
	}, nil
}

func (r *repository) Register(ctx context.Context, user User) error {
	pass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(pass)
	_, err = r.db.Put(ctx, user.ID, user)
	return err
}

func (r *repository) GetUser(ctx context.Context, id string) (string, error) {
	row := r.db.Get(ctx, id)
	var user User
	if err := row.ScanDoc(&user); err != nil {
		return "", err
	}
	return user.Email, nil
}
