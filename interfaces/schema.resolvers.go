package interfaces

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"net/http"

	"github.com/suisuss/habitum_graphQL/models"
)

func (r *queryResolver) Ping(ctx context.Context) (*models.PingResponse, error) {
	return &models.PingResponse{
		Status: http.StatusOK,
	}, nil
}
