// This package is contains users
package users

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username   string `gorm:"primaryKey"`
	Haspass    string
	Email      string
	CreatedOn  time.Time
	ModifiedOn time.Time
}
