package model

import (
	"github.com/google/uuid"
)

// user define domain model and its json and db representations
type User struct {
	UID      uuid.UUID `db:"uid" json:"uid"`
	Email    string    `db:"email" json:"email"`
	Password string    `db:"password" json:"-"` // make sure that password will never send to user
	Name     string    `db:"name" json:"name"`
	ImageURL string    `db:"image_url" json:"image_url"`
	Website  string    `db:"website" json:"website"`
}
