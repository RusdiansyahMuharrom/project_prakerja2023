package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	Id                 uint           `json:"id" gorm:"primaryKey"`
	No_ISBN            string         `json:"no_isbn" gorm:"type:varchar(100);not null"`
	Judul              string         `json:"judul" gorm:"type:varchar(255);not null"`
	Penerbit           string         `json:"penerbit" gorm:"type:varchar(255);not null"`
	Penulis            string         `json:"penulis" gorm:"type:varchar(255);not null"`
	Tanggal_penerimaan string         `json:"tanggal_penerimaan" gorm:"type:date;not null"`
	Tahun_terbit       int            `json:"tahun_terbit" gorm:"not null"`
	Jumlah_halaman     int            `json:"jumlah_halaman" gorm:"not null"`
	Keterangan         *string        `json:"keterangan"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `json:"-" gorm:"index"`
}

type BookRequest struct {
	No_ISBN            string `json:"no_isbn" validate:"required"`
	Judul              string `json:"judul" validate:"required"`
	Penerbit           string `json:"penerbit" validate:"required"`
	Penulis            string `json:"penulis" validate:"required"`
	Tanggal_penerimaan string `json:"tanggal_penerimaan" validate:"required" example:"YYYY-MM-DD"`
	Tahun_terbit       int    `json:"tahun_terbit" validate:"required"`
	Jumlah_halaman     int    `json:"jumlah_halaman" validate:"required"`
	Keterangan         string `json:"keterangan"`
}
