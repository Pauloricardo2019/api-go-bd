package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type User struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Age      int8   `json:"age"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var (
	db    *gorm.DB
	err   error
	Users = []User{}
)

func main() {
	r := gin.Default()

	db, err = gorm.Open(sqlite.Open("mybook.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	db.AutoMigrate(&User{})

	mainRoutes := r.Group("/users")
	mainRoutes.GET("/", Read)
	mainRoutes.POST("/", Create)
	mainRoutes.PUT("/:id", Update)
	mainRoutes.DELETE("/:id", Delete)

	if err := r.Run(":3000"); err != nil {
		panic("Server is don't run")
	}
}

func Read(c *gin.Context) {
	var reqBody User
	result := db.First(&reqBody)
	c.JSON(200, gin.H{
		"error":   false,
		"message": result,
	})

}

func Create(c *gin.Context) {

	var reqBody User
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "Invalid body ",
		})
		return
	}
	db.Create(&reqBody)
	if err != nil {
		c.JSON(500, gin.H{
			"error":   true,
			"message": "Error to create ",
		})
		return
	}

	c.JSON(201, gin.H{
		"error":   false,
		"message": "Create as success",
	})

}

func Update(c *gin.Context) {
	id := c.Param("id")

	var reqBody User

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(422, gin.H{
			"error":   true,
			"message": "Invalid request body",
		})
		return
	}

	for _, u := range Users {
		if u.ID == id {
			u.Name = reqBody.Name
			u.LastName = reqBody.LastName
			u.Age = reqBody.Age
			u.Email = reqBody.Email
			u.Password = reqBody.Password

			c.JSON(200, gin.H{
				"error":   false,
				"message": "Change for success",
			})
			return
		}
	}

	c.JSON(404, gin.H{
		"error":   true,
		"message": "Invalid user id",
	})

}

func Delete(c *gin.Context) {
	id := c.Param("id")

	for _, u := range Users {
		if u.ID == id {
			db.Delete(&u)

			c.JSON(200, gin.H{
				"error":   false,
				"message": "Deleted for success",
			})
		}

	}

	c.JSON(404, gin.H{
		"error":   true,
		"message": "invalid user id",
	})

}
