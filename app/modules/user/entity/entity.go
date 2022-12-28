package entity

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"regexp"
	"time"
)

type (
	User struct {
		Id        uint64    `json:"id,omitempty" gorm:"column:id"`
		Name      string    `json:"name,omitempty" gorm:"column:name"`
		Password  string    `json:"password,omitempty" gorm:"column:password"`
		Email     string    `json:"email,omitempty" gorm:"column:email"`
		PhoneNo   string    `json:"phoneNo,omitempty" gorm:"column:phoneNo"`
		CreatedAt time.Time `json:"createdAt,omitempty" gorm:"column:createdAt"`
		UpdatedAt time.Time `json:"updatedAt,omitempty" gorm:"column:updatedAt"`
	}
)

func (u *User) Validation() error {
	var regex, _ = regexp.Compile(`^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$`)
	if !regex.MatchString(u.Email) {
		return errors.New("format email salah")
	}
	return nil
}

func (u *User) HashPassword() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func (u *User) CheckPasswordHash(hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(u.Password))
	return err == nil
}
