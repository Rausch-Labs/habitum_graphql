package services

import (
	"github.com/suisuss/habitum_graphQL/models"
)

type PingServiceI interface {
	Ping() (*models.PingResponse, error)
}


type PingServiceS struct {
}

func NewPingService() *PingServiceS {
	return &PingServiceS{}
}