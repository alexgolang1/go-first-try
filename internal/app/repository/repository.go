package repository

import (
	"back-api/internal/app/types"
	"context"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	dsn := "any param"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	err = db.AutoMigrate(&types.Model{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func InitRDB() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

type Repository struct {
	db  *gorm.DB
	rdb *redis.Client
}

func NewRepository(database *gorm.DB, rdb *redis.Client) *Repository {
	return &Repository{
		db:  database,
		rdb: rdb,
	}
}

func (repo *Repository) GetUser(ctx context.Context, id int) (*types.Model, error) {
	var person types.Model

	// redis
	convId := strconv.Itoa(id)
	cacheKey := "user:" + convId

	val, err := repo.rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		if err := json.Unmarshal([]byte(val), &person); err == nil {
			return &person, nil
		}
		return nil, err
	}

	// database
	if err := repo.db.First(&person, id).Error; err != nil {
		return nil, err
	}

	data, _ := json.Marshal(person)
	repo.rdb.Set(ctx, cacheKey, data, 10*time.Minute)
	return &person, nil
}

func (repo *Repository) CreateUser(name, surname string) error {
	new_person := types.Model{
		Name:    name,
		Surname: surname,
	}

	return repo.db.Create(&new_person).Error
}

func (repo *Repository) DeleteUser(id int) error {
	return repo.db.Delete(&types.Model{}, id).Error
}

func (repo *Repository) UpdateUser(ctx context.Context, person types.Model) error {
	if err := repo.db.Save(person).Error; err != nil {
		return err
	}

	convId := strconv.Itoa(person.ID)
	repo.rdb.Del(ctx, "person:"+convId)

	return nil
}
