package db

import "github.com/karsharma10/learn_go/config"

func ConnectDb() error {
	// Connect to the database
	dbConfig := config.Configs{}
	getConfig := config.WithDb()
	getConfig(&dbConfig)
	return nil
}
