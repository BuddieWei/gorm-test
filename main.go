package main

import (
	"context"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	// "github.com/buddiewei/gorm-test/pkg/orm/model"

	"github.com/buddiewei/gorm-test/pkg/orm/query"
)

func main() {
	db, _ := gorm.Open(mysql.Open("root:123456@(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"))
	q := query.Use(db)
	// name := "buddie"
	/* q.User.WithContext(context.TODO()).Omit(q.User.Class.Field()).Create(&model.User{
		ID:      "u-123",
		Name:    &name,
		ClassID: "c-123",
	}) */
	user, err := q.User.WithContext(context.TODO()).Where(q.User.ID.Eq("u-123")).Preload(q.User.Class).Take()
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	fmt.Printf("user: %+v\n", user)
}
