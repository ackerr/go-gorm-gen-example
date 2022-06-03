package main

import (
	"github.com/ackerr/go-gorm-example/db"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "dal/query",
		WithUnitTest: true,
	})
	g.UseDB(db.DB)

	g.ApplyBasic(g.GenerateModelAs("user", "User"))

	g.Execute()
}
