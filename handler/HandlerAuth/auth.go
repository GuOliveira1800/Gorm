package HandlerAuth

import (
	"errors"
	"store-backend/model/User"
	"time"

	"gorm.io/gorm"
	"store-backend/config"
	"store-backend/database"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByEmail(email string) (*User.Users, error) {
	db := database.DB
	var user User.Users
	if err := db.Where(&User.Users{Login: email}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// func isEmail(email string) bool {
// 	_, err := mail.ParseAddress(email)
// 	return err == nil
// }

// Login get user and password
func Login(c *fiber.Ctx) error {
	data := struct {
		Login string `json:"login"`
		Senha string `json:"senha"`
	}{}

	type UserData struct {
		ID    uint   `json:"id"`
		Login string `json:"login"`
		Senha string `json:"senha"`
	}
	var userData UserData

	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "data": err})
	}

	pass := data.Senha
	userModel, err := new(User.Users), *new(error)
	userModel, err = getUserByEmail(data.Login)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"status": "error", "message": "Internal Server Error", "data": err})
	} else if userModel == nil {
		CheckPasswordHash(pass, "")
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid identity or password", "data": err})
	} else {
		userData = UserData{
			Login: userModel.Login,
			Senha: userModel.Senha,
			ID:    userData.ID,
		}
	}

	if !CheckPasswordHash(pass, userData.Senha) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid identity or password", "data": nil})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	//JWT data
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userData.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	jwtToken, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	//testando
	c.Cookie(&fiber.Cookie{
		Name:  "token",
		Value: jwtToken,
	})
	return c.JSON(fiber.Map{"id": userData.ID, "email": userData.Login})
}
