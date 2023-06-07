package main

import (
	"design-patterns-go/creational/singleton/database"
	"design-patterns-go/creational/singleton/reader"
	"fmt"
	"sync"
)

var once sync.Once
var instance database.Database
var filePath = ".\\creational\\singleton\\capitals.txt"

type PostgresDB struct {
	capitals map[string]int
}

func (db *PostgresDB) GetPopulation(name string) int {
	return db.capitals[name]
}

func GetPostgresDBInstance() database.Database {
	once.Do(func() {
		db := PostgresDB{}
		caps, err := reader.ReadData(filePath)
		if err == nil {
			db.capitals = caps
		}
		instance = &db
	})
	return instance
}

func main() {
	db1 := GetPostgresDBInstance()
	db2 := GetPostgresDBInstance()
	fmt.Printf("db1[%p] == db2[%p]\n", db1, db2)
}
