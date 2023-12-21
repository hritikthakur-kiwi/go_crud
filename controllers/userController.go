package controllers

import (
	"errors"
	"go_crud/initializers"
	model "go_crud/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
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

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	log.Println("dsdd", tokenString, "dsds", token)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create token",
		})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(200, gin.H{
		"message": "login successFully",
		"user":    user,
	})
}

func UpdateUser(c *gin.Context) {

	id, err := c.Get("user_id")

	if err {
		c.JSON(500, gin.H{"error": "cannot find user"})
	}

	var body struct {
		Name     string
		FullName string
		Contact  uint64
		Email    string
		Address  string
		Gender   string
		Password string
	}
	c.Bind(&body)

	var user model.User

	userDetails := initializers.DB.First(&user, id)

	if userDetails.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user not found"})
	}
	user.Name = body.Name
	user.FullName = body.FullName
	user.Contact = body.Contact
	user.Email = body.Email
	user.Address = body.Address
	user.Gender = body.Gender
	initializers.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"message": "user updated successfully", "user": user})
}
