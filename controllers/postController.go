package controllers

import (
	"go_crud/initializers"
	model "go_crud/models"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func AddPost(c *gin.Context) {
	log.Println(c.Get("user"));
	userId, _ := c.Get("user")
     
	log.Println(userId)
	var body struct {
		PostName string
		userId   uuid.UUID
		Contents  uint64
		draft    string
	}
	c.Bind(&body)
	draftStatus := model.DraftStatus("pending")

	post := model.Post{
		PostName: body.PostName,
		UserId:   userId.(string),
		Contents: body.Contents,
		Draft:    draftStatus,
	}
	log.Print(post)
	addUser := initializers.DB.Create(&post)

	if addUser.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{"user": post})
}
