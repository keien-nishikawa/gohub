package model

import (
	"log"
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

func Username(){
	log.Print("username: keien-nishikawa")
}
