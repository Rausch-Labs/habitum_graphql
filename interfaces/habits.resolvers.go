package interfaces

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/suisuss/habitum_graphQL/generated"
	"github.com/suisuss/habitum_graphQL/models"
)

func (r *mutationResolver) InitHabit(ctx context.Context, request *models.InitHabitRequest) (*models.HabitResponse, error) {
	habit := models.Habit{
		Name: request.Name,
	}

	_, err := r.HabitService.CreateHabit(&habit)
	if err != nil {
		log.Printf("Habit Init Failed: %v", err)
		return nil, err
	}

	return &models.HabitResponse{
		Message: "success",
		Status:  http.StatusOK,
	}, nil
}

func (r *mutationResolver) UpdateHabit(ctx context.Context, request *models.UpdateHabitRequest) (*models.HabitResponse, error) {
	habit, err := r.HabitService.GetHabitByID(request.ID)
	if err != nil {
		log.Println("Error getting the Habit to update: ", err)
		return &models.HabitResponse{
			Message: "Error getting the Habit",
			Status:  http.StatusUnprocessableEntity,
		}, nil
	}

	habit.Name = request.Name

	/*
		if ok, errorString := helpers.ValidateInputs(*ans); !ok {
			return &models.Response{
				Message: errorString,
				Status:  http.StatusUnprocessableEntity,
			}, nil
		}
	*/

	habit, err = r.HabitService.UpdateHabit(habit)
	if err != nil {
		log.Println("Habit updating error: ", err)
		return &models.HabitResponse{
			Message: "Error updating Habit",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.HabitResponse{
		Message: "Successfully updated habit",
		Status:  http.StatusOK,
		Data:    habit,
	}, nil
}

func (r *mutationResolver) DeleteHabit(ctx context.Context, request *models.DeleteHabitRequest) (*models.HabitResponse, error) {
	log.Printf("Received: delete %v", request.ID)

	_, err := r.HabitService.DeleteHabit(request.ID)
	if err != nil {
		log.Printf("Habit Delete Failed: %v", err)
		return nil, err
	}

	return &models.HabitResponse{
		Message: "Habit " + request.ID + " Deleted",
		Status:  http.StatusOK,
	}, nil
}

func (r *mutationResolver) LogHabit(ctx context.Context, request *models.LogHabitRequest) (*models.HabitLogResponse, error) {
	log.Printf("Received: LogHabit %v, %v, %v", request.Type, request.Note, request.Date)

	_habitLog := models.HabitLog{
		Type: request.Type,
		Note: request.Note,
		Date: request.Date, // TODO: Have name conventions match
	}

	habitLog, err := r.HabitLogService.CreateHabitLog(&_habitLog)
	if err != nil {
		log.Printf("Habit Log Failed: %v", err)
		return nil, err
	}

	return &models.HabitLogResponse{
		Message: "Habit Logged",
		Status:  http.StatusOK,
		Data:    habitLog,
	}, nil
}

func (r *mutationResolver) UpdateHabitLog(ctx context.Context, request *models.UpdateHabitLogRequest) (*models.HabitLogResponse, error) {
	habitLog, err := r.HabitLogService.GetHabitLogByID(request.ID)
	if err != nil {
		log.Println("Error getting the HabitLog to update: ", err)
		return &models.HabitLogResponse{
			Message: "Error getting the HabitLog",
			Status:  http.StatusUnprocessableEntity,
		}, nil
	}

	habitLog.Type = request.Type
	habitLog.Date = request.Date
	habitLog.Note = request.Note

	/*
		if ok, errorString := helpers.ValidateInputs(*ans); !ok {
			return &models.Response{
				Message: errorString,
				Status:  http.StatusUnprocessableEntity,
			}, nil
		}
	*/

	habitLog, err = r.HabitLogService.UpdateHabitLog(habitLog)
	if err != nil {
		log.Println("HabitLog updating error: ", err)
		return &models.HabitLogResponse{
			Message: "Error updating HabitLog",
			Status:  http.StatusInternalServerError,
		}, nil
	}

	return &models.HabitLogResponse{
		Message: "Successfully updated habit",
		Status:  http.StatusOK,
		Data:    habitLog,
	}, nil
}

func (r *mutationResolver) DeleteHabitLog(ctx context.Context, request *models.DeleteHabitLogRequest) (*models.HabitLogResponse, error) {
	log.Printf("Received: DeleteHabitLog %v", request.ID)

	_, err := r.HabitLogService.DeleteHabitLog(request.ID)
	if err != nil {
		log.Printf("DeleteHabitLog Failed: %v", err)
		return nil, err
	}

	return &models.HabitLogResponse{
		Message: "Habit Log " + request.ID + " Deleted",
		Status:  http.StatusOK,
	}, nil
}

func (r *queryResolver) GetHabitLogIDByHabit(ctx context.Context, request *models.LogHabitRequest) (*models.HabitLogResponse, error) {
	log.Printf("Received: GetHabitLogIDByHabit %v", request.Date)

	habitLogs, err := r.HabitLogService.GetAllHabitLogsByDate(request.Date)
	if err != nil {
		log.Printf("Habit GetHabitLogIDByHabit Failed: %v", err)
		return nil, err
	}

	for _, oldHabitLog := range habitLogs {
		if oldHabitLog.Type == request.Type {
			return &models.HabitLogResponse{
				Message: string(oldHabitLog.ID),
				Status:  http.StatusOK,
				Data:    oldHabitLog,
			}, nil
		}
	}

	return nil, errors.New("Not habitlog found")
}

func (r *queryResolver) GetHabitIDByName(ctx context.Context, request *models.InitHabitRequest) (*models.HabitResponse, error) {
	habit, err := r.HabitService.GetAllHabitsByName(request.Name)
	if err != nil {
		fmt.Println("fail")
	}

	return &models.HabitResponse{
		Message: habit[0].ID,
		Status:  http.StatusOK,
		Data:    habit[0],
	}, nil
}

func (r *queryResolver) GetHabitLogByID(ctx context.Context, request *models.GetHabitLogByIDRequest) (*models.HabitLogResponse, error) {
	habitLog, err := r.HabitLogService.GetHabitLogByID(request.ID)
	if err != nil {
		return &models.HabitLogResponse{
			Message: "Error getting the HabitLog",
			Status:  http.StatusUnprocessableEntity,
		}, nil
	}

	return &models.HabitLogResponse{Message: string(habitLog.ID)}, nil
}

func (r *queryResolver) GetHabitByID(ctx context.Context, request *models.GetHabitByIDRequest) (*models.HabitResponse, error) {
	habit, err := r.HabitService.GetHabitByID(request.ID)
	if err != nil {
		return &models.HabitResponse{
			Message: "Error getting the Habit",
			Status:  http.StatusUnprocessableEntity,
		}, nil
	}

	return &models.HabitResponse{
		Message: string(habit.ID),
		Status:  http.StatusOK,
		Data:    habit,
	}, nil
}

func (r *queryResolver) GetAllHabitLogs(ctx context.Context) (*models.HabitLogResponse, error) {
	habitLogs, err := r.HabitLogService.GetAllHabitLogs()
	if err != nil {
		return &models.HabitLogResponse{
			Message: "Error getting the HabitLogs",
			Status:  http.StatusUnprocessableEntity,
		}, nil
	}

	return &models.HabitLogResponse{
		Message:  "Success",
		Status:   http.StatusOK,
		DataList: habitLogs,
	}, nil
}

func (r *queryResolver) GetAllHabits(ctx context.Context) (*models.HabitResponse, error) {
	habits, err := r.HabitService.GetAllHabits()
	if err != nil {
		return &models.HabitResponse{
			Message: "Error getting the Habits",
			Status:  http.StatusUnprocessableEntity,
		}, nil
	}

	return &models.HabitResponse{
		Message:  "Success",
		Status:   http.StatusOK,
		DataList: habits,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
