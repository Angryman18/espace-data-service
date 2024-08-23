package Env

import (
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	MongoURI string
	Port     string
}

func GetEnv() *Env {
	if err := godotenv.Load(); err != nil {
		panic("could not load the environment variables")
	}
	port := os.Getenv("PORT")
	mongoURI := os.Getenv("MONGODB_URI")
	return &Env{
		MongoURI: mongoURI,
		Port:     port,
	}
}
