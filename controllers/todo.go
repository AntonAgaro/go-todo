package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-todo/config"
	"go-todo/models"
	"net/http"
)

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	config.DB.Limit(10).Find(&todos)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"todos":  todos,
	})
}

func GetById(c *gin.Context) {
	id := c.Param("id")
	var todo *models.Todo
	if err := config.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Todo was not found!",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"todo":   todo,
	})
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	config.DB.Create(&todo)
	c.JSON(http.StatusCreated, gin.H{
		"status":  "ok",
		"message": "todo was created!",
		"todo":    todo,
	})
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")
	var todo *models.Todo

	if err := config.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Todo was not found!",
		})
		return
	}
	var input *models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	todo.Title = input.Title
	todo.Completed = input.Completed

	config.DB.Save(&todo)

	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
		"todo":   todo,
	})
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")
	var todo *models.Todo

	if err := config.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Todo was not found!",
		})
		return
	}

	config.DB.Delete(&todo)

	c.JSON(http.StatusOK, gin.H{
		"status":  "ok",
		"message": fmt.Sprintf("Task with id %d was deleted", todo.ID),
	})
}
