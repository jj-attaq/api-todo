package model

import (
	"github.com/google/uuid"
	/* "gorm.io/gorm"
	"time" */
)

/* type User struct {
	UID        uuid.UUID `gorm:"primary_key; unique; type:uuid; column:id; default:uuid_generate_v4()"` // this needs to be associated to the todoIDs so we know which todo corresponds to whom.
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
    Email string
    Password string
    Name string
} */

type User struct {
    UID      uuid.UUID `db:"uid" json:"uid"`
    Email    string    `db:"email" json:"email"`
    Password string    `db:"password" json:"-"`
    Name     string    `db:"name" json:"name"`
    ImageURL string    `db:"image_url" json:"imageUrl"`
    Website  string    `db:"website" json:"website"`
}
