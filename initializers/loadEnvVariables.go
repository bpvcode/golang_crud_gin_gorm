package initializers

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	fmt.Println(os.Getenv("DB_USER"))
	fmt.Println(os.Getenv("DB_PASS"))
	fmt.Println(os.Getenv("PORT"))
	fmt.Println(os.Getenv("DB_HOST"))
	fmt.Println(os.Getenv("DB_USER"))
	fmt.Println(os.Getenv("DB_PASS"))
	fmt.Println(os.Getenv("DB_NAME"))
	fmt.Println(os.Getenv("DB_PORT"))
	fmt.Println(os.Getenv("DB_SSL_MODE"))
}
