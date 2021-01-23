package auth

import "time"

//User ...
type User struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Password    string    `json:"password"`
	CreatedDate time.Time `json:"createdDate,omitempty"`
	LastLogin   time.Time `json:"lastLogin,omitempty"`
	Rev         string    `json:"_rev,omitempty"`
}

//Query ...
type Query struct {
	Selector map[string]Selector `json:"selector"`
}

//Selector can be used in the selector field of the couch query
type Selector struct {
	Value interface{} `json:"$eq,omitempty"`
}
