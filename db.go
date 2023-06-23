package main

import "gorm.io/gorm"

type dbData struct {
	gorm.Model
	GivenNumber int
	CalcNumber  int
}
