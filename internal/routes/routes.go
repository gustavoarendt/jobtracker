package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gustavoarendt/jobtracker/configs"
	"github.com/gustavoarendt/jobtracker/internal/controllers"
)

func HandleRequest(config *configs.Config) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/users", controllers.CreateUser)
	r.Post("/users/login", controllers.Login)

	http.ListenAndServe(config.ServerPort, r)
}
