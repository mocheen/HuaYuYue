package main

import (
	"fmt"
	"gorm.io/gen"
	"role/conf"
	"role/internal/repository"
)

type Querier interface {
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {
	c := conf.GetConf()
	fmt.Println(c.MySQL)
	repository.InitDB()

	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/repository/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(repository.DB)

	g.ApplyBasic(repository.AdminAPL{}, repository.UserRole{})

	g.Execute()
}
