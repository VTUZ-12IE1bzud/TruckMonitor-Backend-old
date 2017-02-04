package env

import "os"

var (
	ServiceHost = os.Getenv("SERVICE_HOST")
	DBHost      = os.Getenv("DB_HOST")
	DBPort      = os.Getenv("DB_PORT")
	DBDatabase  = os.Getenv("DB_DATABASE")
	DBUser      = os.Getenv("DB_USER")
	DBPassword  = os.Getenv("DB_PASSWORD")
)
