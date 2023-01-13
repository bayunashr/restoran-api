package main

import (
	"github.com/bayunashr/restoran-api/controllers"
	"github.com/bayunashr/restoran-api/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnv()
	initializers.LoadDb()
}

func main() {
	app := gin.Default()
	app.POST("/kategori", controllers.CreateKategori)
	app.GET("/kategori", controllers.ReadAllKategori)
	app.GET("/kategori/:id", controllers.ReadOneKategori)
	app.PUT("/kategori/:id", controllers.UpdateKategori)
	app.DELETE("/kategori/:id", controllers.DeleteKategori)
	app.Run()
}
