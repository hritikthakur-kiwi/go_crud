package controllers

import (
	"go_crud/initializers"
	model "go_crud/models"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

func AddPost(c *gin.Context) {

	var body struct {
		postName string
		userId   uuid.UUID
		contents  uint64
		draft    string
	}
	c.Bind(&body)
	draftStatus := model.DraftStatus("pending")

	post := model.Post{
		PostName: "ksksjkddfddd",
		UserId:   "19476a76-eb23-4616-96aa-0baa2da4211a",
		Contents:  12,
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
