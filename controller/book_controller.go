package controller

import (
	"net/http"
	"project_prakerja2023/configuration"
	"project_prakerja2023/model"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// Post Data Buku
// @Summary     Insert Data Buku
// @Description Semua wajib diisi kecuali keterangan
// @ID          insert-buku
// @Accept      json
// @Produce     json
// @Param       Input body model.BookRequest true "Request Body"
// @Tags        Book
// @Success     200 {object} model.ResponseOK
// @Failure     400 {object} model.ResponseError
// @Router      /books [post]
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

// Get Data Buku All
// @Summary     Get Data Buku All
// @ID          get-buku-all
// @Produce     json
// @Tags        Book
// @Success     200 {object} model.ResponseOK
// @Failure     400 {object} model.ResponseError
// @Router      /books [get]
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

// Get Data Buku By Id
// @Summary     Get Data Buku By Id
// @ID          get-buku-by-id
// @Produce     json
// @Tags        Book
// @Param       id path string true "id buku"
// @Success     200 {object} model.ResponseOK
// @Failure     400 {object} model.ResponseError
// @Router      /books/{id} [get]
func GetBookById(e echo.Context) error {
	//Proses menerima parameter id
	id := e.Param("id")

	book := model.Book{}

	//Proses get data dari database
	err := configuration.DB.First(&book, id).Error
	if err != nil {
		return e.JSON(http.StatusBadRequest, model.ResponseError{
			Message: "Get Data Gagal!",
			Error:   err.Error(),
		})
	}

	return e.JSON(http.StatusOK, model.ResponseOK{
		Message: "Sukses!",
		Data:    book,
	})

}

// Update Data Buku
// @Summary     Update Data Buku
// @Description Semua wajib diisi kecuali keterangan
// @ID          update-buku
// @Accept      json
// @Produce     json
// @Tags        Book
// @Param       id path string true "id buku"
// @Param       Input body model.BookRequest true "Request Body"
// @Success     200 {object} model.ResponseOK
// @Failure     400 {object} model.ResponseError
// @Router      /books/{id} [put]
func UpdateBook(e echo.Context) error {
	//Proses menerima parameter id
	id := e.Param("id")

	bookRequest := model.BookRequest{}
	book := model.Book{}

	//Proses pengecekan data apakah ada atau tidak
	err := configuration.DB.First(&book, id).Error
	if err != nil {
		return e.JSON(http.StatusBadRequest, model.ResponseError{
			Message: "Get Data Gagal!",
			Error:   err.Error(),
		})
	}

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

	//Proses perubahan data
	book.No_ISBN = bookRequest.No_ISBN
	book.Judul = bookRequest.Judul
	book.Penerbit = bookRequest.Penerbit
	book.Penulis = bookRequest.Penulis
	book.Tanggal_penerimaan = bookRequest.Tanggal_penerimaan
	book.Tahun_terbit = bookRequest.Tahun_terbit
	book.Jumlah_halaman = bookRequest.Jumlah_halaman
	if bookRequest.Keterangan != "" {
		book.Keterangan = &bookRequest.Keterangan
	}

	//Proses update data ke database
	errSave := configuration.DB.Save(&book).Error
	if errSave != nil {
		return e.JSON(http.StatusBadRequest, model.ResponseError{
			Message: "Update Data Gagal!",
			Error:   errSave.Error(),
		})
	}

	return e.JSON(http.StatusOK, model.ResponseOK{
		Message: "Sukses!",
		Data:    book,
	})

}

// Delete Data Buku
// @Summary     Delete Data Buku
// @ID          delete-buku
// @Produce     json
// @Tags        Book
// @Param       id path string true "id buku"
// @Success     200 {object} model.ResponseOK
// @Failure     400 {object} model.ResponseError
// @Router      /books/{id} [delete]
func DeleteBook(e echo.Context) error {
	//Proses menerima parameter id
	id := e.Param("id")

	book := model.Book{}

	//Proses pengecekan data apakah ada atau tidak
	err := configuration.DB.First(&book, id).Error
	if err != nil {
		return e.JSON(http.StatusBadRequest, model.ResponseError{
			Message: "Get Data Gagal!",
			Error:   err.Error(),
		})
	}

	//Proses delete data ke database
	errDelete := configuration.DB.Delete(&book).Error
	if errDelete != nil {
		return e.JSON(http.StatusBadRequest, model.ResponseError{
			Message: "Delete Data Gagal!",
			Error:   errDelete.Error(),
		})
	}

	return e.JSON(http.StatusOK, model.ResponseOK{
		Message: "Sukses!",
		Data:    "Data dengan id " + id + " berhasil dihapus!",
	})

}
