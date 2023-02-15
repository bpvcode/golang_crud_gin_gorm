package main

import (
	"github.com/bpvcode/golang_crud_gin_gorm/controllers"
	"github.com/bpvcode/golang_crud_gin_gorm/db"
	"github.com/bpvcode/golang_crud_gin_gorm/http"
	"github.com/bpvcode/golang_crud_gin_gorm/initializers"
)

// Runs once, before main() function
// Should load environment variables variable, invoking a function of initializer package
func init() {
	initializers.LoadEnvVariables()
	db.ConnectToDB()
	db.MigrateTables()
	//db.MigrateShipmentsMockData()
}

func main() {
	server := http.New()
	controllers.Router(server)
	http.Listen(server)
}
