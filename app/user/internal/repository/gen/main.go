package main

import (
	"gorm.io/gen"
	"user/conf"
	"user/internal/repository"
)

// Dynamic SQL
type Querier interface {
	// SELECT * FROM @@table WHERE name = @name{{if role !=""}} AND role = @role{{end}}
	FilterWithNameAndRole(name, role string) ([]gen.T, error)
}

func main() {
	conf.InitConfig()
	repository.InitDB()

	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/repository/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(repository.DB)

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	g.ApplyBasic(repository.User{})
	//g.ApplyInterface(func(Querier) {}, repository.User{})

	// Generate the code
	g.Execute()
}
