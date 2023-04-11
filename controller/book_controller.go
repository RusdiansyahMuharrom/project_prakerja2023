package controller

import (
	"net/http"
	"project_prakerja2023/configuration"
	"project_prakerja2023/model"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func CreateBook(e echo.Context) error {
	bookRequest := model.BookRequest{}

	//Proses penginputan request
	errBind := e.Bind(&bookRequest)
	if errBind != nil {
		return e.JSON(http.StatusBadRequest, model.ResponseError{
			Message: "Periksa kembali inputan anda!",
			Error:   errBind.Error(),
		})
	}

	//Validasi data yang diinputkan
	var validate = validator.New()
	errValidate := validate.Struct(bookRequest)
	if errValidate != nil {
		return e.JSON(http.StatusBadRequest, model.ResponseError{
			Message: "Periksa kembali inputan anda!",
			Error:   errValidate.Error(),
		})
	}

	//Inisialisasi data baru
	new_book := model.Book{
		No_ISBN:            bookRequest.No_ISBN,
		Judul:              bookRequest.Judul,
		Penerbit:           bookRequest.Penerbit,
		Penulis:            bookRequest.Penulis,
		Tanggal_penerimaan: bookRequest.Tanggal_penerimaan,
		Tahun_terbit:       bookRequest.Tahun_terbit,
		Jumlah_halaman:     bookRequest.Jumlah_halaman,
	}
	if bookRequest.Keterangan != "" {
		new_book.Keterangan = &bookRequest.Keterangan
	}

	//Proses input data ke database
	errCreate := configuration.DB.Create(&new_book).Error
	if errCreate != nil {
		return e.JSON(http.StatusBadRequest, model.ResponseError{
			Message: "Input Data Gagal!",
			Error:   errCreate.Error(),
		})
	}

	return e.JSON(http.StatusOK, model.ResponseOK{
		Message: "Sukses!",
		Data:    new_book,
	})

}

func GetAllBook(e echo.Context) error {
	books := []model.Book{}

	//Proses get data dari database
	err := configuration.DB.Find(&books).Error
	if err != nil {
		return e.JSON(http.StatusInternalServerError, model.ResponseError{
			Message: "Get Data Gagal!",
			Error:   err.Error(),
		})
	}

	return e.JSON(http.StatusOK, model.ResponseOK{
		Message: "Sukses!",
		Data:    books,
	})

}
