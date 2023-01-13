package models

import "gorm.io/gorm"

type Kategori struct {
	gorm.Model
	Nama string
}

type Menu struct {
	gorm.Model
	IDKategori int
	Nama       string
	Harga      int
	Gambar     string
}
