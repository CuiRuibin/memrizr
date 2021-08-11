package model

import (
	"github.com/google/uuid"
)

// User defines domain model and its json and db represent
type User struct {
	UID      uuid.UUID `db:"uid" json:"uid"`
	Email    string    `db:"mail" json:"email"`
	Password string    `db:"password" json:"-"`
	Name     string    `db:"name" json:"name"`
	ImageUrl string    `db:"image_url" json:"imageUrl"`
	WWebsite string    `db:"website" json:"website"`
}
