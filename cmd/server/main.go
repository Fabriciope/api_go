package main

import (
	"fmt"
	"net/http"

	"github.com/Fabriciope/my-api/configs"
	"github.com/Fabriciope/my-api/internal/infra/webserver/handlers"
	"github.com/Fabriciope/my-api/pkg"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/jwtauth"
)

func init() {
	err := configs.LoadConfig(".")
	if err != nil {
		pkg.LogError("Error: load config", err)
	}
}

func main() {
	// TODO: estudar context para aplicar tempo limite às requisições
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	makeRoutes(r)

	http.ListenAndServe(fmt.Sprintf(":%d", configs.Cfg.WebServerPort), r)
}

// TODO: testar todas as requisições
func makeRoutes(r *chi.Mux) {
	h, err := handlers.LoadHandlers()
	if err != nil {
		pkg.LogError("Error: load handlers", err)
		return
	}

	r.Route("/product", func(r chi.Router) {
		r.Use(jwtauth.Verifier(configs.Cfg.JWTTokenAuth))
		r.Use(jwtauth.Authenticator)

		hp := h.Product
		r.Post("/create", hp.Create)
		r.Put("/update/{id}", hp.Update)
		r.Delete("/delete/{id}", hp.Delete)
		r.Get("/{id}", hp.GetByID)
		r.Get("/all/{page}/{limit}", hp.GetAll)
	})

	r.Route("/user", func(r chi.Router) {
		hu := h.User
		r.Post("/create", hu.Create)
		r.Post("/generate_jwt", hu.GetJWT)
	})
}
