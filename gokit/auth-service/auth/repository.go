package auth

import (
	"context"
	"errors"
	"time"

	"github.com/go-kivik/kivik"
	"golang.org/x/crypto/bcrypt"
)

//Repository consists of all abstract methods which interfaces with the database implementation
type Repository interface {
	Login(ctx context.Context, user User) error
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

func (r *repository) Login(ctx context.Context, user User) error {
	rows, err := r.db.Find(ctx, Query{
		Selector: map[string]Selector{
			"email": {
				user.Email,
			},
		},
	})
	storedUser := User{}
	for rows.Next() {
		if err := rows.ScanDoc(&storedUser); err != nil {
			return err
		}
	}
	defer rows.Close()

	errf := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password))
	if errf != nil && errf == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return errors.New("invalid login credentials, please try again")
	}

	storedUser.LastLogin = time.Now()
	_, err = r.db.Put(ctx, storedUser.ID, storedUser)

	return err
}

func (r *repository) Register(ctx context.Context, user User) error {
	rows, err := r.db.Find(ctx, Query{
		Selector: map[string]Selector{
			"email": {
				user.Email,
			},
		},
	})

	for rows.Next() {
		return errors.New("email already in persistence")
	}
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
