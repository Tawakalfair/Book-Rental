package auth

import (
	"os"
	"time"

	"github.com/Tawakalfair/Book-Rental/book"
	"github.com/Tawakalfair/Book-Rental/database"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Hash the password with a high cost for security
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := book.User{
		Username: data["username"],
		Password: string(password),
	}

	database.DB.Create(&user)

	return c.JSON(user)
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user book.User

	database.DB.Where("username = ?", data["username"]).First(&user)

	if user.ID == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User not found",
		})
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect password",
		})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    string(rune(user.ID)),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 1 day
	})

	secretKey := os.Getenv("SECRET_KEY")
	token, err := claims.SignedString([]byte(secretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Could not log in",
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

// ShowRegisterPage renders the registration form
func ShowRegisterPage(c *fiber.Ctx) error {
	return c.Render("register", nil)
}

// ShowLoginPage renders the login form
func ShowLoginPage(c *fiber.Ctx) error {
	return c.Render("login", nil)
}

// RegisterUser handles the form submission from the registration page
func RegisterUser(c *fiber.Ctx) error {
	password, _ := bcrypt.GenerateFromPassword([]byte(c.FormValue("password")), 14)
	user := book.User{
		Username: c.FormValue("username"),
		Password: string(password),
	}
	database.DB.Create(&user)
	return c.Redirect("/login")
}

// LoginUser authenticates a user and sets a cookie
func LoginUser(c *fiber.Ctx) error {
	var user book.User
	database.DB.Where("username = ?", c.FormValue("username")).First(&user)

	if user.ID == 0 || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(c.FormValue("password"))) != nil {
		return c.Status(fiber.StatusUnauthorized).Render("login", fiber.Map{"error": "Invalid username or password"})
	}

	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    string(rune(user.ID)),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)), // 1 day
	})

	token, err := claims.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).Render("error", fiber.Map{"message": "Could not log in"})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	})

	return c.Redirect("/dashboard")
}

// LogoutUser logs the user out by clearing the cookie
func LogoutUser(c *fiber.Ctx) error {
	c.Cookie(&fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	})
	return c.Redirect("/login")
}

// ShowDashboard shows a protected dashboard page
func ShowDashboard(c *fiber.Ctx) error {
	// You can get user info from the JWT if needed
	// For now, we just show a generic dashboard
	return c.Render("dashboard", fiber.Map{
		"username": "User", // In a real app, you'd extract this from the JWT
	})
}
