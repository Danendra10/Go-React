package controllers

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/danendra10/gowlang-first/database"
	"github.com/danendra10/gowlang-first/models"
	"github.com/danendra10/gowlang-first/utils"
	"github.com/gofiber/fiber/v2"
)

func validateEmail(email string) bool {
	Re := regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)

	return Re.MatchString(email)
}

func Register(c *fiber.Ctx) error {
	var data map[string]interface{}

	var userData models.User

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}

	if len(data["password"].(string)) <= 6 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password must be greater than 6 characters",
		})
	}

	if !validateEmail(strings.TrimSpace(data["email"].(string))) {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Invalid Email Address",
		})
	}

	// check if email exist
	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(userData)

	if userData.Id != 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email already exist",
		})
	}

	user := models.User{
		Username: data["username"].(string),
		Email:    strings.TrimSpace(data["email"].(string)),
	}

	user.SetPassword(data["password"].(string))

	err := database.DB.Create(&user)
	if err != nil {
		log.Println(err)
	}

	c.Status(200)
	return c.JSON(fiber.Map{
		"message": "Register Success",
		"user":    user,
	})
}

func Login(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		fmt.Println("Unable to parse body")
	}

	var user models.User

	database.DB.Where("email=?", data["email"]).First(&user)

	if user.Id == 0 {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Email address does not exist",
		})
	}

	if err := user.PasswordCompare(data["password"]); err != nil {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "Password does not match",
		})
	}

	token, err := utils.GenerateJwt(strconv.Itoa(user.Id))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return nil
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie((&cookie))

	return c.JSON(fiber.Map{
		"message": "Login Success",
		"user":    user,
		"token":   token,
	})
}
