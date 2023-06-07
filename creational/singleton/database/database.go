package database

type Database interface {
	GetPopulation(name string) int
}
