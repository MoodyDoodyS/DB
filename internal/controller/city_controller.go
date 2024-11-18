package controller

import (
	"awesomeProject16/internal/models"
	"awesomeProject16/internal/services"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

type CityHandler struct {
	service services.CityService
}

func NewCityHandler(service services.CityService) *CityHandler {
	return &CityHandler{service: service}
}

func (h *CityHandler) CreateCity(w http.ResponseWriter, r *http.Request) {
	var city models.City
	if err := json.NewDecoder(r.Body).Decode(&city); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := h.service.CreateCity(&city); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func (h *CityHandler) DeleteCity(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	if err := h.service.DeleteCity(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
func (h *CityHandler) UpdateCity(w http.ResponseWriter, r *http.Request) {
	// Извлекаем ID из параметра URL
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Извлекаем данные для обновления из тела запроса
	var city models.City
	if err := json.NewDecoder(r.Body).Decode(&city); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Устанавливаем ID города из параметра URL
	city.ID = id

	// Вызываем метод сервиса для обновления города
	if err := h.service.UpdateCity(&city); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Возвращаем успешный ответ
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(city)
}

func (h *CityHandler) GetCities(w http.ResponseWriter, r *http.Request) {
	cities, err := h.service.GetCities()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(cities)
}
