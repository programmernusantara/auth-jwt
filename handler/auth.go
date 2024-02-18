package handler

import (
	"errors"
	"net/mail"
	"time"

	"api-fiber-gorm/config"   // Import package config untuk mengambil konfigurasi aplikasi
	"api-fiber-gorm/database" // Import package database untuk akses database
	"api-fiber-gorm/model"    // Import package model untuk akses model data

	"gorm.io/gorm" // Import package gorm untuk ORM

	"github.com/gofiber/fiber/v2"  // Import package Fiber untuk framework web
	"github.com/golang-jwt/jwt/v5" // Import package jwt untuk otentikasi token JWT
	"golang.org/x/crypto/bcrypt"   // Import package bcrypt untuk pengaturan password
)

// CheckPasswordHash digunakan untuk membandingkan password dengan hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// getUserByEmail digunakan untuk mendapatkan pengguna berdasarkan alamat email
func getUserByEmail(e string) (*model.User, error) {
	db := database.DB
	var user model.User
	if err := db.Where(&model.User{Email: e}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Tidak ada pengguna yang ditemukan dengan alamat email yang diberikan
		}
		return nil, err // Terjadi kesalahan saat mengakses database
	}
	return &user, nil // Mengembalikan pengguna yang ditemukan
}

// getUserByUsername digunakan untuk mendapatkan pengguna berdasarkan nama pengguna
func getUserByUsername(u string) (*model.User, error) {
	db := database.DB
	var user model.User
	if err := db.Where(&model.User{Username: u}).Find(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil // Tidak ada pengguna yang ditemukan dengan nama pengguna yang diberikan
		}
		return nil, err // Terjadi kesalahan saat mengakses database
	}
	return &user, nil // Mengembalikan pengguna yang ditemukan
}

// isEmail digunakan untuk memeriksa apakah string yang diberikan adalah alamat email yang valid
func isEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil // Mengembalikan true jika alamat email valid, false jika tidak valid
}

// Login digunakan untuk proses login pengguna
func Login(c *fiber.Ctx) error {
	type LoginInput struct {
		Identity string `json:"identity"` // Identity bisa berupa email atau username
		Password string `json:"password"`
	}
	type UserData struct {
		ID       uint   `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	input := new(LoginInput)
	var userData UserData

	if err := c.BodyParser(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}

	identity := input.Identity
	pass := input.Password
	userModel, err := new(model.User), *new(error)

	if isEmail(identity) {
		userModel, err = getUserByEmail(identity)
	} else {
		userModel, err = getUserByUsername(identity)
	}

	if userModel == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "data": err})
	} else {
		userData = UserData{
			ID:       userModel.ID,
			Username: userModel.Username,
			Email:    userModel.Email,
			Password: userModel.Password,
		}
	}

	if !CheckPasswordHash(pass, userData.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = userData.Username
	claims["user_id"] = userData.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
}
