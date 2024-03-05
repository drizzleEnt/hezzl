package model

import "time"

type Good struct {
	Id          int
	ProjectId   int
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Priority    int64     `json:"priority"`
	Removed     bool      `json:"removed"`
	CreatedAt   time.Time `json:"createdAt"`
}

type Meta struct {
	Total   int
	Removed int
	Limit   int
	Offset  int
}
