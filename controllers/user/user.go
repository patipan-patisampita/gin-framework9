package usercontroller

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/patipan-patisampita/gin-framework9/configs"
	"github.com/patipan-patisampita/gin-framework9/models"
)

// CRUD
//Read
func GetAll(c *gin.Context) {
	var users []models.User
	configs.DB.Order("id desc").Find(&users)

	//SQL
	configs.DB.Raw("SELECT * FROM users ORDER by id desc").Scan(&users)
	c.JSON(200, gin.H{
		"data new": users,
	})
}

//Create
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
	//check repeat email
	userExist := configs.DB.Where("email = ?", input.Email).First(&user)
	if userExist.RowsAffected == 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "มีผู้ใช้งานอีเมล์นี้ในระบบแล้ว"})
		return
	}
	result := configs.DB.Debug().Create(&user)

	//db error
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error})
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

//Read Get by id
func GetById(c *gin.Context) {
	id := c.Param("id")
	var user models.User
	result := configs.DB.First(&user, id)

	if result.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูล"})
		return
	}

	c.JSON(200, gin.H{
		"data": user,
	})
}

func SearchByFullName(c *gin.Context) {
	fullname := c.Query("fullname")

	var users[]models.User
	result := configs.DB.Where("fullname LIKE ?","%" + fullname + "%").Find(&users)
	if result.RowsAffected<1{
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูล"})
		return
	}

	c.JSON(200, gin.H{
		"data": fullname,
	})
}
