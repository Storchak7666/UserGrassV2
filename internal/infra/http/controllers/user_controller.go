package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/test_server/internal/domain/user"
	"net/http"
	"strconv"
)

type UserController struct {
	service *user.Service
}

func NewUserController(s *user.Service) *UserController {
	return &UserController{
		service: s,
	}
}

func (c *UserController) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		users, err := (*c.service).FindAll()
		if err != nil {
			fmt.Printf("UserController.FindAll(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("UserController.FindAll(): %s", err)
			}
			return
		}
		err = success(w, users)
		if err != nil {
			fmt.Printf("UserController.FindAll(): %s", err)
		}
	}
}

func (c *UserController) FindById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("UserController.FindById(): %s", err)
			err = badRequest(w, err)
			if err != nil {
				fmt.Printf("UserController.FindById(): %s", err)
			}
			return
		}
		user, err := (*c.service).FindById(id)
		if err != nil {
			fmt.Printf("UserController.FindById(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("UserController.FindById(): %s", err)
			}
			return
		}

		err = success(w, user)
		if err != nil {
			fmt.Printf("UserController.FindById(): %s", err)
		}
	}
}

func (c *UserController) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user user.User

		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			fmt.Printf("UserController.CreateUser(): %s", err)
			err = badRequest(w, err)
			if err != nil {
				fmt.Printf("UserController.CreateUser(): %s", err)
			}
			return
		}

		// Do something with the Person struct...

		userService, err := (*c.service).CreateUser(user)
		if err != nil {
			fmt.Printf("UserController.CreateUser(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("UserController.CreateUser(): %s", err)
			}
			return
		}

		err = success(w, userService)
		if err != nil {
			fmt.Printf("UserController.CreateUser(): %s", err)
		}
	}
}

func (c *UserController) UpdateById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user user.User
		// Try to decode the request body into the struct. If there is an error,
		// respond to the client with the error message and a 400 status code.
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			fmt.Printf("UserController.UpdateById(): %s", err)
			err = badRequest(w, err)
			if err != nil {
				fmt.Printf("UserController.UpdateById(): %s", err)
			}
			return
		}

		// Do something with the Person struct...
		//fmt.Fprintf(w, "Person: %+v", user)

		_, err = (*c.service).UpdateById(user)
		if err != nil {
			fmt.Printf("UserController.UpdateById(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("UserController.UpdateById(): %s", err)
			}
			return
		}
	}
}
