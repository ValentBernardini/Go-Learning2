package main

import (
	"fmt"
	"gomysql/db"
	"gomysql/models"
)

func main() {
	db.Connect()
	defer db.Close()
	//fmt.Println(db.ExistTable("users"))
	//db.CreateTable(models.UserSchema, "users")
	user := models.CreateUsers("nashe", "140226", "nashe@gmail.com")
	fmt.Println(user)
	//db.TruncateTable("users")
	fmt.Println(models.ListUsers())

}
