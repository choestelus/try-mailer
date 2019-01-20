package api

import "time"

type History struct {
	ID         string    `json:"id"`
	Email      string    `json:"email"`
	AccessFrom string    `json:"access_from"`
	Status     string    `json:"status"`
	Failed     bool      `json:"failed"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
