package model

import (
	"time"
)

type User struct {
	id int
	firstName string
	familyName string
	birthday time.Time
	createdAt time.Time
	updatedAt time.Time
}

func Username() string{
	return "keien-nishikawa"
}
