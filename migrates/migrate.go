package main

import (
	"github.com/bayunashr/restoran-api/initializers"
	"github.com/bayunashr/restoran-api/models"
)

func init() {
	initializers.LoadDb()
	initializers.LoadEnv()
}

func main() {
	initializers.DB.AutoMigrate(&models.Kategori{})
	initializers.DB.AutoMigrate(&models.Menu{})
}
