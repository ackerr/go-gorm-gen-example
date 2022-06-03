package main

import (
	"context"
	"fmt"

	"github.com/ackerr/go-gorm-example/dal/model"
	"github.com/ackerr/go-gorm-example/dal/query"
	"github.com/ackerr/go-gorm-example/db"
)

func main() {
	ctx := context.Background()

	u := query.Use(db.DB).User

	// create
	err := u.WithContext(ctx).Debug().Create(&model.User{Name: "demo1", Age: 20})
	if err != nil {
		panic(err)
	}

	// select
	count, err := u.WithContext(ctx).Debug().Where(u.Name.Like("demo%")).Count()
	if err != nil {
		panic(err)
	}
	fmt.Printf("user count: %d\n", count)

	// transaction
	err = query.Use(db.DB).Transaction(func(tx *query.Query) error {
		err := tx.User.WithContext(ctx).Debug().Create(&model.User{Name: "demo2"})
		if err != nil {
			return err
		}

		user, err := tx.User.WithContext(ctx).Debug().Where(u.Name.Like("demo%")).Last()
		if err != nil {
			return err
		}
		fmt.Printf("user name: %v\n", user.Name)
		return nil

	})
	if err != nil {
		panic(err)
	}
}
