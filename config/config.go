package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB      DB
	OpenApi OpenApi
	Server  Server
}

type DB struct {
	DBName string
	DBPort string
	DBHost string
	DBPwd  string
	DBUser string
}

type OpenApi struct {
	ApiNation string
	ApiGender string
	ApiAge    string
}

type Server struct {
	Port string
}

func NewServer() *Server {
	Port, exist := os.LookupEnv("SERVER_PORT")
	if !exist {
		return nil
	}
	return &Server{Port: Port}
}

func NewOpenApi() *OpenApi {
	ApiAge, exist := os.LookupEnv("API_AGE")
	if !exist {
		return nil
	}
	ApiGender, exist := os.LookupEnv("API_GENDER")
	if !exist {
		return nil
	}
	ApiNation, exist := os.LookupEnv("API_NATION")
	if !exist {
		return nil
	}
	return &OpenApi{ApiAge: ApiAge, ApiGender: ApiGender, ApiNation: ApiNation}
}

func NewDB() *DB {
	DBHost, exist := os.LookupEnv("DB_HOST")
	if !exist {
		return nil
	}
	DBPort, exist := os.LookupEnv("DB_PORT")
	if !exist {
		return nil
	}
	DBUser, exist := os.LookupEnv("DB_USER")
	if !exist {
		return nil
	}
	DBPwd, exist := os.LookupEnv("DB_PWD")
	if !exist {
		return nil
	}
	DBName, exist := os.LookupEnv("DB_NAME")
	if !exist {
		return nil
	}
	return &DB{DBName: DBName, DBPort: DBPort, DBHost: DBHost, DBPwd: DBPwd, DBUser: DBUser}
}
func NewConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	DB := NewDB()
	if DB == nil {
		return nil
	}
	OpenApi := NewOpenApi()
	if OpenApi == nil {
		return nil
	}
	Server := NewServer()
	if Server == nil {
		return nil
	}

	return &Config{DB: *DB, OpenApi: *OpenApi, Server: *Server}
}
