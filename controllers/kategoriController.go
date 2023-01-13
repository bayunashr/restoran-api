package controllers

import (
	"github.com/bayunashr/restoran-api/initializers"
	"github.com/bayunashr/restoran-api/models"
	"github.com/gin-gonic/gin"
)

func CreateKategori(c *gin.Context) {
	var newKategori struct {
		Nama string
	}
	c.Bind(&newKategori)

	kategori := models.Kategori{Nama: newKategori.Nama}
	result := initializers.DB.Create(&kategori)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "failed to create new kategori",
		})
	} else {
		c.JSON(200, gin.H{
			"kategori": kategori,
		})
	}
}

func ReadAllKategori(c *gin.Context) {
	var kategori []models.Kategori
	result := initializers.DB.Find(&kategori)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "failed to read all kategori",
		})
	} else {
		c.JSON(200, gin.H{
			"kategori": kategori,
		})
	}
}

func ReadOneKategori(c *gin.Context) {
	id := c.Param("id")
	var kategori models.Kategori
	result := initializers.DB.First(&kategori, id)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "failed to read kategori",
		})
	} else {
		c.JSON(200, gin.H{
			"kategori": kategori,
		})
	}
}

func UpdateKategori(c *gin.Context) {
	id := c.Param("id")
	var updatedKategori struct {
		Nama string
	}
	c.Bind(&updatedKategori)

	var kategori models.Kategori
	initializers.DB.First(&kategori, id)
	result := initializers.DB.Model(&kategori).Updates(models.Kategori{
		Nama: updatedKategori.Nama,
	})

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "failed to update kategori",
		})
	} else {
		c.JSON(200, gin.H{
			"kategori": kategori,
		})
	}
}

func DeleteKategori(c *gin.Context) {
	id := c.Param("id")
	var kategori models.Kategori
	result := initializers.DB.Delete(&models.Kategori{}, id)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"error": "failed to delete kategori",
		})
	} else {
		c.JSON(200, gin.H{
			"kategori": kategori,
		})
	}
}
