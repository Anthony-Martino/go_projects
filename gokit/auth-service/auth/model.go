package auth

import "time"

//User ...
type User struct {
	ID        string    `json:"id,omitempty"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	LastLogin time.Time `json:"lastLogin,omitempty"`
}
