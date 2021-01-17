package account

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var errInvalidRequest = errors.New("invalid request")

//MakeHTTPHandler mounts all of the service endpoints into a http.Handler
func MakeHTTPHandler(s Service) http.Handler {
	r := mux.NewRouter()
	e := MakeEndpoints(s)

	r.Methods("POST").Path("/user").Handler(httptransport.NewServer(
		e.CreateUser,
		decodeCreateUserRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/user/{id}").Handler(httptransport.NewServer(
		e.GetUser,
		decodeGetUserRequest,
		encodeResponse,
	))

	return r
}

type (
	//CreateUserRequest ...
	CreateUserRequest struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	//CreateUserResponse ...
	CreateUserResponse struct {
		Ok string `json:"ok"`
	}
	// GetUserRequest ...
	GetUserRequest struct {
		ID string `json:"id"`
	}
	// GetUserResponse ...
	GetUserResponse struct {
		Email string `json:"email"`
	}
)

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

func decodeCreateUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	req := CreateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	return req, err
}

func decodeGetUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)

	req := GetUserRequest{
		ID: vars["id"],
	}
	return req, nil
}
