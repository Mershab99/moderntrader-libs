package main

import (
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/rawsql"
	"log"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./gen/sql",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})
	// https://github.com/go-gorm/rawsql/blob/master/tests/gen_test.go
	gormdb, err := gorm.Open(rawsql.New(rawsql.Config{
		//SQL:      rawsql,                      //create table sql
		FilePath: []string{
			//"./sql/user.sql", // create table sql file
			"../sql/fixtures.sql", // create table sql file directory
		},
	}))

	if err != nil {
		log.Fatalf("GORM ERROR: %s", err)

	}
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions

	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}
