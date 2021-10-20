package interfaces

import (
	"github.com/suisuss/habitum_graphQL/services"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	HabitService						services.HabitServiceI
	HabitLogService					services.HabitLogServiceI
	UserService 						services.UserServiceI
	// PingService							services.PingServiceI			
}


