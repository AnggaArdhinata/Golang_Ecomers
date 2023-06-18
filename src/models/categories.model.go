package models

import "time"

type Categories struct {
	Id         int       `json:"id"`
	Name       string    `json:"category"`
	Created_At time.Time `json:"created_at"`
	Updated_At time.Time `json:"update_at"`
}
