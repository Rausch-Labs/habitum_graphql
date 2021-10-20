package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"github.com/suisuss/habitum_graphQL/services"
	"github.com/suisuss/habitum_graphQL/generated"
	"github.com/suisuss/habitum_graphQL/db_test"
	"github.com/suisuss/habitum_graphQL/interfaces"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/rs/cors"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	var (
		defaultPort = "8080"
		database = os.Getenv("DATABASE")
	)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	log.Print(database)

	dbConn := db.OpenDB(database)

	var habitService services.HabitServiceI
	var habitLogService services.HabitLogServiceI
	var userService services.UserServiceI
	// var pingService services.PingServiceI

	habitService = services.NewHabitService(dbConn)
	habitLogService = services.NewHabitLogService(dbConn)
	userService = services.NewUserService(dbConn)
	// pingService = services.NewPingService()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &interfaces.Resolver{
		HabitService:           habitService,
		HabitLogService:       	habitLogService,
		UserService: 						userService,
		// PingService: 						pingService,
	}}))

	router := chi.NewRouter()

	// Add CORS middleware around every request
	// See https://github.com/rs/cors for full option listing
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)



	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
