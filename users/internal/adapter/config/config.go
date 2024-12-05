package config

import (
	"os"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		App       *App
		DB        *DB
		Redis     *Redis
		Transport *Transport
	}
	App struct {
		Name string
		Env  string
	}
	DB struct {
		Connection string
		Host       string
		Port       string
		Name       string
		User       string
		Password   string
	}
	Redis struct {
		Host     string
		Port     string
		Password string
	}
	Transport struct {
		Env  string
		Host string
		Port string
	}
)

func New() (*Container, error) {
	if os.Getenv("APP_ENV") != "prod" {
		if err := godotenv.Load(); err != nil {
			return nil, err
		}
	}

	app := &App{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
	}
	db := &DB{
		Connection: os.Getenv("DB_CONNECTION"),
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		Name:       os.Getenv("DB_NAME"),
		User:       os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASSWORD"),
	}
	redis := &Redis{
		Host: os.Getenv("REDIS_HOST"),
		Port: os.Getenv("REDIS_PORT"),
	}
	transport := &Transport{
		Env:  os.Getenv("APP_ENV"),
		Host: os.Getenv("TRANSPORT_HOST"),
		Port: os.Getenv("TRANSPORT_PORT"),
	}

	return &Container{
		App:       app,
		DB:        db,
		Redis:     redis,
		Transport: transport,
	}, nil
}
