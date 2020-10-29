package main

import (
	"fmt"
	"os"
	
	"github.com/ediyasaedi/dk-case/database"
	"github.com/ediyasaedi/dk-case/user"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	jwtware "github.com/gofiber/jwt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func authRequired() func(c *fiber.Ctx){
	jwtSecret := os.Getenv("JWT_SECRET")
	return jwtware.New(jwtware.Config{
		ErrorHandler: func(c *fiber.Ctx, err error){
			c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		},
		SigningKey: []byte(jwtSecret),
	})
}

func initDatabase()  {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "users.db")

	if err != nil {
		panic("Failed to connecting database!")
	}

	fmt.Println("Database connection successfully opened!")

	database.DBConn.AutoMigrate(&user.User{})
	fmt.Println("Database migrated!")
}

func setupRoutes(app *fiber.App)  {
	app.Post("/api/v1/signup", user.RegisterUser)
	app.Post("/api/v1/signin", user.LoginUser)
	app.Get("/api/v1/getone", authRequired(), user.GetOne)
	app.Get("/api/v1/getall", user.GetUsers)
}

func main(){
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()

	app.Use(middleware.Logger())

	setupRoutes(app)

	err := app.Listen(3000)
	if err != nil {
		panic(err)
	}
}