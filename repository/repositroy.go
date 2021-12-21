package repository

import "gorm.io/gorm"

type Repositories struct {
	Db *gorm.DB
}

var Repo *Repositories

func Get() *Repositories {
	return Repo
}
