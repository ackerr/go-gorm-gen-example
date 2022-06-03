package main

import (
	"fmt"

	"github.com/ackerr/go-gorm-example/db"
)

type User struct {
	ID   uint `gorm:"primarykey"`
	Name string
	Age  int8
}

func main() {
	db.DB.Table("user").AutoMigrate(&User{})
	fmt.Println("migrate table success")
}
