package model

import (
	"errors"
	"github.com/alainmucyo/my_brand/config"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID        uint64    `gorm:"primary_key,auto_increment" json:"id" validate:"required"`
	Name      string    `json:"name" gorm:"size:255"`
	Email     string    `json:"email" gorm:"unique,size:255" validate:"required,email"`
	Image     string    `json:"image" gorm:"size:255"`
	Password  string    `json:"password" gorm:"size:255"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type UserResponse struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (user *User) Prepare() {
	hashedPassword, _ := Hash(user.Password)
	user.Password = string(hashedPassword)
}

func (user User) Save() (User, error) {
	if config.Database.Save(&user).Error != nil {
		return User{}, errors.New("Can't create")
	}
	return user, nil
}

func (user User) CheckUser() (User, error) {
	password := user.Password
	if config.Database.Where("email = ?", user.Email).Take(&user).Error != nil {
		return User{}, errors.New("Wrong credentials")
	}
	err := VerifyPassword(user.Password, password)
	if err != nil {
		return User{}, errors.New("Wrong credentials")
	}
	user.Password = "hidden"
	return user, nil
}

func (User) FindById(id uint64) (User, error) {
	var user User
	if config.Database.Where("id = ?", id).Take(&user).Error != nil {
		return User{}, errors.New("User not found")
	}
	return user, nil
}
