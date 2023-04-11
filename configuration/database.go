package configuration

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {
	db_host := Get_env("DB_HOST")
	db_port := Get_env("DB_PORT")
	db_database := Get_env("DB_DATABASE")
	db_username := Get_env("DB_USERNAME")
	db_password := Get_env("DB_PASSWORD")

	//  dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", db_username, db_password, db_host, db_port, db_database)
	// fmt.Println(connection)
	DB, err = gorm.Open(mysql.Open(connection), &gorm.Config{})

	if err != nil {
		panic("Connection Failed!")
	}
	fmt.Println("Connection Success!")
}
