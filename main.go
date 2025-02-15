package main

import (
	"net/http"
	"strconv"

	"examples/htmx-go/components"
	"examples/htmx-go/models"

	"github.com/gin-gonic/gin"
)

func main() {
	InitDatabase()
	defer DB.Close()

	e := gin.Default()
	e.LoadHTMLGlob("templates/*")

	e.GET("/", func(c *gin.Context) {
		todos := ReadToDoList()
		c.HTML(http.StatusOK, "index.html", gin.H{
			"todos": todos,
		})
	})

	e.POST("/todos", func(c *gin.Context) {
		title := c.PostForm("title")
		status := c.PostForm("status")
		id, _ := CreateToDo(title, status)

		/* c.HTML leverages Go native templating */
		// c.HTML(http.StatusOK, "task", gin.H{
		// 	"Title":  title,
		// 	"Status": status,
		// 	"Id":     id,
		// })

		c.Writer.Header().Set("Content-Type", "text/html; charset=utf-8")
		c.Status(http.StatusOK)
		component := components.Task(models.ToDo{Id: id, Title: title, Status: status})
		if err := component.Render(c.Request.Context(), c.Writer); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		}
	})

	e.DELETE("/todos/:id", func(c *gin.Context) {
		param := c.Param("id")
		id, _ := strconv.ParseInt(param, 10, 64)
		DeleteToDo(id)
	})

	e.Run(":8080")
}
