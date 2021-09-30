package database

import (
	"github.com/pkg/errors"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func MySQLDB() (gorm.DB, error) {
	if db == nil {
		if err := initDatabase(dsn); err != nil {
			return gorm.DB{}, err
		}
	}
	return *db, nil
}

func initDatabase(dsn string) error {
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
