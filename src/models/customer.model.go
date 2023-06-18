package models

import "time"

type Customer struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"update_at"`
}
