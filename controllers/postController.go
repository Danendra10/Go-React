package controllers

import (
	"strconv"

	"github.com/danendra10/gowlang-first/database"
	"github.com/danendra10/gowlang-first/models"
	"github.com/gofiber/fiber/v2"
)

func CreatePost(c *fiber.Ctx) error {
	var blog_post models.Post

	if err := c.BodyParser(&blog_post); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Failed to parse body",
			"code":    fiber.StatusBadRequest,
		})
	}

	if err := database.DB.Create(&blog_post).Error; err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Failed to create post",
			"code":    fiber.StatusInternalServerError,
		})
	}

	return c.JSON(fiber.Map{
		"message": "Post created successfully",
		"code":    fiber.StatusOK,
		"data":    blog_post,
	})
}

func AllPost(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit := 5
	offset := (page - 1) * limit

	var total int64
	var get_post []models.Post

	database.DB.Preload("User").Offset(offset).Limit(limit).Find(&get_post)
	database.DB.Model(&models.Post{}).Count(&total)

	return c.JSON(fiber.Map{
		"posts": get_post,
		"meta": fiber.Map{
			"total":     total,
			"page":      page,
			"limit":     limit,
			"last_page": float64(int(total) / limit),
		},
	})
}

func GetPOst(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var post models.Post
	database.DB.Where("id=?", id).First(&post)
	return c.JSON(fiber.Map{
		"post":    post,
		"code":    fiber.StatusOK,
		"message": "success",
	})
}
