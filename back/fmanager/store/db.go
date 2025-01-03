package store

import (
	"log"
	"os"

	"github.com/icehardstone/fmanager/config"
	"gorm.io/gorm"

	"gorm.io/driver/postgres"
)

type ModelBase struct {
	gorm.Model
}

type Store struct {
	db *gorm.DB
}

func NewStore() *Store {
	db, err := gorm.Open(postgres.Open(config.AppConfig.Database.DNS), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return &Store{
		db: db,
	}
}

func AutoMigrate() {
	args := os.Args
	log.Printf("loading args: %s", args)

	for _, arg := range args {
		if arg == "seed" {
			db := NewStore().db
			log.Printf("begining to migrate database")
			db.AutoMigrate(&File{})
			break
		}
	}
}
