package controllers

import (
	"errors"
	"go_crud/initializers"
	model "go_crud/models"
	"log"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func Create(res *gin.Context) {

	var body struct {
		Name     string
		FullName string
		Contact  uint64
		Email    string
		Address  string
		Gender   string
		Password string
	}
	res.Bind(&body)
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)
	if err != nil {
		res.Status(500)
		return
	}
	user := model.User{
		Name:     body.Name,
		FullName: body.FullName,
		Contact:  body.Contact,
		Email:    body.Email,
		Address:  body.Address,
		Gender:   body.Gender,
		Password: string(passwordHash),
	}
	addUser := initializers.DB.Create(&user)

	if addUser.Error != nil {
		res.Status(400)
		return
	}

	res.JSON(200, gin.H{"user": user})
}

func Login(c *gin.Context) {
	// id := c.Param("id")
	var body struct {
		Email    string
		Password string
	}
	c.Bind(&body)
	var user model.User
	err := initializers.DB.Where("Email = ?", body.Email).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(401, gin.H{"error": "User not found"})
			return
		} 
	}
	log.Println("sddd", user)
	comparePass := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if comparePass != nil {
		c.JSON(401, gin.H{"error": "Invalid password"})
		return
	}
	c.JSON(200, gin.H{
		"message": "login successFully",
		"user":    user,
	})

}
