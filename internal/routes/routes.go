package routes

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/gustavoarendt/jobtracker/configs"
	"github.com/gustavoarendt/jobtracker/internal/controllers"
)

func HandleRequest(config *configs.Config) {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Route("/users", func(r chi.Router) {
		r.Post("/", controllers.CreateUser)
		r.Post("/login", controllers.Login)
	})

	r.Route("/companies", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", controllers.CreateCompany)
		r.Get("/", controllers.GetCompanies)
		r.Get("/{id}", controllers.GetCompany)
		r.Put("/{id}", controllers.UpdateCompany)
		r.Delete("/{id}", controllers.DeleteCompany)
	})

	r.Route("/jobs", func(r chi.Router) {
		r.Use(jwtauth.Verifier(config.TokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/", controllers.CreateJob)
		r.Get("/", controllers.GetJobs)
		r.Get("/{id}", controllers.GetJob)
		r.Put("/{id}", controllers.UpdateJob)
		r.Delete("/{id}", controllers.DeleteJob)
	})

	http.ListenAndServe(config.ServerPort, r)
}
