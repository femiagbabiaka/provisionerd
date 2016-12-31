package main

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type virtualMailer struct {
	gorm.Model
}
