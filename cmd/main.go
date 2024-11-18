package main

import (
	"awesomeProject16/config"
	controller "awesomeProject16/internal/controller"
	repository "awesomeProject16/internal/repository"
	services "awesomeProject16/internal/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
)

func main() {
	db, err := config.GetDBConnection()
	if err != nil {
		log.Fatal(err)
	}

	repo := repository.NewMySQLCityRepository(db)
	service := services.NewMySQLCityService(repo)
	handler := controller.NewCityHandler(service)

	/*
		err = repository.InitDB(db)
			if err != nil {
				log.Fatal(err)
			}
	*/
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Post("/cities", handler.CreateCity)
	r.Get("/cities", handler.GetCities)
	r.Delete("/cities/{id}", handler.DeleteCity)
	r.Put("/cities/{id}", handler.UpdateCity)
	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
