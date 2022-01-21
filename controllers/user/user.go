package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/patipan-patisampita/gin-framework9/configs"
	"github.com/patipan-patisampita/gin-framework9/models"
)

func GetAll(c *gin.Context) {
	var users []models.User
	configs.DB.Find(&users)

	c.JSON(200, gin.H{
		"data new": users,
	})
}

func Register(c *gin.Context) {
	var input InputRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Fullname: input.Fullname,
		Email:    input.Email,
		Password: input.Password,
	}
	result := configs.DB.Debug().Create(&user)

	//db error
	if result.Error !=nil{
		c.JSON(http.StatusBadRequest,gin.H{"error":result.Error})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "สมัครสมาชิกสำเร็จแล้ว",
	})
}

func Login(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "login",
	})
}

func GetById(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "id",
	})
}

func SearchByFullName(c *gin.Context) {
	fullname := c.Query("fullname")
	c.JSON(200, gin.H{
		"data": fullname,
	})
}
