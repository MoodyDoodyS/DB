package services

import (
	"awesomeProject16/internal/models"
	"awesomeProject16/internal/repository"
)

type CityService interface {
	CreateCity(city *models.City) error
	UpdateCity(city *models.City) error
	DeleteCity(id int) error
	GetCities() ([]models.City, error)
}

type MySQLCityService struct {
	repo repository.CityRepository
}

func NewMySQLCityService(repo repository.CityRepository) *MySQLCityService {
	return &MySQLCityService{
		repo: repo,
	}
}

func (s *MySQLCityService) CreateCity(city *models.City) error {
	return s.repo.Create(city)
}

func (s *MySQLCityService) UpdateCity(city *models.City) error {
	return s.repo.Update(city)
}

func (s *MySQLCityService) DeleteCity(id int) error {
	return s.repo.Delete(id)
}

func (s *MySQLCityService) GetCities() ([]models.City, error) {
	return s.repo.List()
}
