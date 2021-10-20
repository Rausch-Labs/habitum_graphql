package interfaces

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/suisuss/habitum_graphQL/models"
)

func (r *mutationResolver) CreateUser(ctx context.Context, request *models.CreateUserRequest) (*models.UserResponse, error) {
	// Validation
	if request.Username == "" || request.Password == "" || request.Authority < 0 || request.Email == "" {
		return &models.UserResponse{
			Message: "'user', 'password', 'authority' and 'email' parameters are required",
			Status:  http.StatusBadRequest,
		}, nil
	}

	// Create user object
	hashedPassword, _ := r.UserService.HashPassword(request.Password)

	user := &models.User{
		Username:  request.Username,
		Password:  hashedPassword,
		Email:     request.Email,
		Authority: request.Authority,
	}

	user.CreatedAt = time.Now()

	//Save the message to database:
	user, err := r.UserService.CreateUser(user)
	if err != nil {
		fmt.Println("the error with this: ", err)
		return &models.UserResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
			Data:    user,
		}, nil
	}

	return &models.UserResponse{
		Message: "Successfully created User",
		Status:  http.StatusCreated,
		Data:    user,
	}, nil
}

func (r *mutationResolver) DeleteUserByID(ctx context.Context, request *models.DeleteUserRequest) (*models.UserResponse, error) {
	err := r.UserService.DeleteUserByID(request.ID)

	if err != nil {
		return &models.UserResponse{
			Message: "Something went wrong deleting that User.",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.UserResponse{
		Message: "Successfully deleted that User",
		Status:  http.StatusOK,
	}, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, request *models.UpdateUserRequest) (*models.UserResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetOneUser(ctx context.Context, request *models.GetOneUserRequest) (*models.UserResponse, error) {
	users, err := r.UserService.GetAllUsers()

	if err != nil {
		log.Println("getting all users error: ", err)
		return &models.UserResponse{
			Message: "Something went wrong getting all Users.",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.UserResponse{
		Message:  "Successfully retrieved all messages",
		Status:   http.StatusOK,
		DataList: users,
	}, nil
}

func (r *queryResolver) GetAllUsers(ctx context.Context) (*models.UserResponse, error) {
	users, err := r.UserService.GetAllUsers()
	if err != nil {
		return &models.UserResponse{
			Message: "Error getting the Users",
			Status:  http.StatusUnprocessableEntity,
		}, nil
	}

	return &models.UserResponse{
		Message:  "Success",
		Status:   http.StatusOK,
		DataList: users,
	}, nil
}

func (r *queryResolver) AuthorizeUser(ctx context.Context, request *models.AuthorizeUserRequest) (*models.JWTResponse, error) {
	// Get user
	user, err := r.UserService.GetUserByUsername(request.Username)
	if err != nil {
		log.Println("Failed username search for ", request.Username)
		return &models.JWTResponse{
			Message: "Invaild username or password",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	// Check password
	if r.UserService.CheckPasswordHash(user.Password, request.Password) != true {
		log.Println("Failed password check for", user.Username)

		return &models.JWTResponse{
			Message: "Invaild username or password",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	jwtString, err := r.UserService.GenerateJWT(user.ID, user.Username)

	jwt := &models.Jwt{
		Token: jwtString,
	}

	if err != nil {
		fmt.Println(err)
		return &models.JWTResponse{
			Message: "JWT generation error",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.JWTResponse{
		Message: "Authorization Successful",
		Status:  http.StatusOK,
		Jwt:     jwt,
	}, nil
}
