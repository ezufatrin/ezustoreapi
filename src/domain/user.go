package domain

import "time"

type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(100)" json:"name"`
	Email     string    `gorm:"uniqueIndex;type:varchar(100)" json:"email"`
	Password  string    `gorm:"type:varchar(255)" json:"-"` // hidden in response
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRepository interface {
	Register(user *User) error
	Update(user *User) error
	GetByID(id uint) (*User, error)
	GetByEmail(email string) (*User, error)
}

type UserUsecase interface {
	Register(name, email, password string) error
	Login(email, password string) (string, error) // returns JWT token
	GetProfile(userID uint) (*User, error)
	UpdateProfile(userID uint, name, email string) error
}
