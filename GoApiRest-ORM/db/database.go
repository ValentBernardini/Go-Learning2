package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Realizar coneccion a base de datos con GORM
var dsn = "root:valen1402@tcp(localhost:3306)/goweb_db?charset=utf8mb4&parseTime=True&loc=Local"
var DataBase = func() (db *gorm.DB) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("failed to connect database", err)
		panic(err)
	}
	fmt.Println("Conexion exitosa")
	return db
}()
