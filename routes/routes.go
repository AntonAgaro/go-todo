package routes

import (
	"github.com/gin-gonic/gin"
	"go-todo/controllers"
)

func RegisterRotes(r *gin.Engine) {
	todo := r.Group("/todo")
	{
		todo.GET("/", controllers.GetTodos)
		todo.GET("/:id", controllers.GetById)
		todo.POST("/", controllers.CreateTodo)
		todo.PUT("/:id", controllers.UpdateTodo)
		todo.DELETE("/:id", controllers.DeleteTodo)
	}
}
