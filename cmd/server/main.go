package main

import (
	"net/http"

	"github.com/Fabriciope/my-api/configs"
	"github.com/Fabriciope/my-api/internal/infra/webserver/handlers"
	"github.com/Fabriciope/my-api/pkg"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func init() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	configs.Cfg = config
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

	http.ListenAndServe(":8000", r)
}

func makeRoutes(r *chi.Mux) {
	h, err := handlers.LoadHandlers()
	if err != nil {
		pkg.LogError("Error: load handlers", err)
		return
	}

	r.Route("/product", func(r chi.Router) {
		r.Post("/create", h.Product.Create)
		r.Put("/update/{id}", h.Product.Update)
		r.Delete("/delete/{id}", h.Product.Delete)
		r.Get("/{id}", h.Product.GetByID)
		r.Get("/all/{page}/{limit}", h.Product.GetAll)
	})
}
