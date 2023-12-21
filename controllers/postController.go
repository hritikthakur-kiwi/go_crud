package controllers

import (
	"go_crud/initializers"
	model "go_crud/models"

	"github.com/gin-gonic/gin"
)

func AddPost(c *gin.Context) {

	var body struct {
		postName string
		userId   uint
		content  uint64	
		draft    string
	}
	c.Bind(&body)
	draftStatus := model.DraftStatus(body.draft)

	post := model.Post{
		PostName: body.postName,
		UserId:   body.userId,
		Content:  body.content,
		Draft:    draftStatus,
	}
	addUser := initializers.DB.Create(&post)

	if addUser.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{"user": post})
}
	