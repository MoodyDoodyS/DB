package repository

import (
	"awesomeProject16/internal/models"
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
)

type CityRepository interface {
	Create(city *models.City) error
	Update(city *models.City) error
	Delete(id int) error
	List() ([]models.City, error)
}

func InitDB(db *sqlx.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS cities (
		id SERIAL  PRIMARY KEY ,
		name VARCHAR(30) NOT NULL,
		state VARCHAR(30) NOT NULL
	);
	`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("ошибка создания таблицы: %v", err)
	}

	log.Println("Таблица cities успешно создана или уже существует.")
	return nil
}
func DropTable(db *sqlx.DB) error {
	query := `DROP TABLE IF EXISTS cities;`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("ошибка при удалении таблицы: %v", err)
	}

	log.Println("Таблица cities успешно удалена.")
	return nil
}

type MySQLCityRepository struct {
	db *sqlx.DB
}

func NewMySQLCityRepository(db *sqlx.DB) *MySQLCityRepository {
	return &MySQLCityRepository{
		db: db,
	}
}
func (r *MySQLCityRepository) Create(city *models.City) error {
	query := `INSERT INTO cities (id, name, state) VALUES ( DEFAULT,$1, $2)`
	_, err := r.db.Exec(query, city.Name, city.State)
	if err != nil {
		return fmt.Errorf("ошибка при создании записи: %v", err)
	}
	return nil
}
func (r *MySQLCityRepository) Update(city *models.City) error {
	query := `UPDATE cities SET name = $1, state = $2 WHERE id = $3`
	_, err := r.db.Exec(query, city.Name, city.State, city.ID)
	if err != nil {
		return fmt.Errorf("ошибка при обновлении записи: %v", err)
	}
	return nil
}

func (r *MySQLCityRepository) Delete(id int) error {
	query := `DELETE FROM cities WHERE id = $1`
	_, err := r.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("ошибка при удалении записи: %v", err)
	}
	return nil
}
func (r *MySQLCityRepository) List() ([]models.City, error) {
	var cities []models.City
	query := `SELECT id, name, state FROM cities`
	err := r.db.Select(&cities, query)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении списка городов: %v", err)
	}
	return cities, nil
}
