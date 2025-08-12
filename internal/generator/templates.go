package generator

const apiGoTemplate = `
package api

import (
	"github.com/go-chi/chi/v5"
)

type Api struct {
	Router *chi.Mux
}

func New() *Api {
	api := &Api{}
	api.Init()
	return api
}

func (api *Api) Init() {
	api.Router = chi.NewRouter()
	api.setupMiddleware()
	api.setupRoutes()
}
`

const jsonutilsGoTemplate = `
package jsonutils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func EncodeJson[T any](w http.ResponseWriter, r *http.Request, statusCode int, data T) error {
	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		return fmt.Errorf("failed to encode json %w", err)
	}

	return nil
}

func DecodeJson[T any](r *http.Request) (T, error) {
	var data T
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return data, fmt.Errorf("decode json failed: %w", err)
	}

	return data, nil
}

func SendError(w http.ResponseWriter, r *http.Request, statusCode int, message string) {
	response := map[string]string{
		"error":  message,
		"status": "error",
	}
	EncodeJson(w, r, statusCode, response)
}
`

const mainGoTemplate = `
package main

import (
	"log"
	"net/http"
	"{{.Module}}/internal/api"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	api := api.New()

	addr := ":3050"

	(log.Printf("Server is running on  http://localhost%s\n", addr))
	if err := http.ListenAndServe(addr, api.Router); err != nil {
		log.Fatalf("Error to initialize server: %v", err)
	}
}
`

const routerGoTemplate = `
package api

import (
	"net/http"
	"{{.Module}}/internal/jsonutils"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

func (api *Api) setupMiddleware() {
	api.Router.Use(middleware.Logger)
	api.Router.Use(middleware.Recoverer)
	api.Router.Use(middleware.RequestID)

	api.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		MaxAge:         300,
	}))
}

func (api *Api) setupRoutes() {
	api.Router.Get("/health", api.handleHealth)

}

func (api *Api) handleHealth(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"status":  "OK",
		"service": "api",
	}
	if err := jsonutils.EncodeJson(w, r, http.StatusOK, response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
`
