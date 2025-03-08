package models

import (
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

type User struct {
    gorm.Model
    Username string `json:"username" gorm:"unique;not null"`
    Password string `json:"password" gorm:"not null"`
}

// Hash password before storing
func (u *User) HashPassword() error {
    bytes, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    u.Password = string(bytes)
    return nil
}

// Check password during login
func (u *User) CheckPassword(password string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
    return err == nil
}