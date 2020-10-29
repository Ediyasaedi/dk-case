package user

import (
	"fmt"
	"time"
	"os"
	
	"github.com/ediyasaedi/dk-case/database"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

// User ... 
type User struct {
	gorm.Model
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}

// GetUsers ...
func GetUsers(c *fiber.Ctx)  {
	db := database.DBConn
	var users []User
	db.Find(&users)
	c.JSON(users)
}

// RegisterUser ...
func RegisterUser(c *fiber.Ctx)  {
	db := database.DBConn
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		c.Status(503).Send(err)
		return
	}

	db.Create(user)
	var newUser = map[string]string{
        "email":    user.Email,
		"username": user.Username,
		"msg": "Register successfully",
    }
	c.JSON(newUser)
}

// GetOne ...
func GetOne(c *fiber.Ctx){
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	email := claims["email"].(string)
	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"username" : username,
		"email": email,
	})
}

// LoginUser ...
func LoginUser(c *fiber.Ctx) {

	jwtSecret := os.Getenv("JWT_SECRET")
	db := database.DBConn
	var body User
	err := c.BodyParser(&body)
	if err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
		return
	}

	var user User
	db.Where("Email = ? AND Password = ?", body.Email, body.Password).Find(&user)

	fmt.Println(user)

	if body.Email != user.Email || body.Password != user.Password {
		c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Bad credentials",
		})
		return
	}
	
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["exp"] = time.Now().Add(time.Minute * 30)

	s, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return
	}

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": s,
	})

}
