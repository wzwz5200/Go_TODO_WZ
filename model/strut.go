package model

import (
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model

	Title      string
	Descripion string

	Status bool
}
