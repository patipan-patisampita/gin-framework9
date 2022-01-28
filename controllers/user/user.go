package usercontroller

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/matthewhartstonge/argon2"
	"github.com/patipan-patisampita/gin-framework9/configs"
	"github.com/patipan-patisampita/gin-framework9/models"
	"github.com/patipan-patisampita/gin-framework9/utils"
)

// CRUD
//Read
func GetAll(c *gin.Context) {
	var users []models.User
	///configs.DB.Order("id desc").Find(&users)

	//SQL
	//configs.DB.Raw("SELECT * FROM users ORDER by id desc").Scan(&users)
	configs.DB.Preload("Blogs").Find(&users)


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

	var input InputLogin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Email:    input.Email,
		Password: input.Password,
	}

	userAccount := configs.DB.Where("email = ?", input.Email).First(&user)
	if userAccount.RowsAffected < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ไม่พบผู้ใช้งานในระบบ"})
		return
	}

	//เปรียบเทียบรหัสผ่านที่ส่งมา
	ok, _ := argon2.VerifyEncoded([]byte(input.Password), []byte(user.Password))
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "รหัสผ่านไม่ถูกต้อง"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message":      "เข้าสู่ระบบสำเร็จ",
		"access_token": "token",
	})

	//สร้าง token
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24 * 2).Unix(),
	})

	jwtSecret := os.Getenv("JWT_SECRET")
	token, _ := claims.SignedString([]byte(jwtSecret))

	c.JSON(http.StatusCreated, gin.H{
		"message":      "เข้าระบบสำเร็จ",
		"access_token": token,
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

func SearchByFullname(c *gin.Context) {
	fullname := c.Query("fullname") //?fullname=John

	var users []models.User
	result := configs.DB.Where("fullname LIKE ?", "%"+fullname+"%").Scopes(utils.Paginate(c)).Find(&users)
	if result.RowsAffected < 1 {
		c.JSON(http.StatusNotFound, gin.H{"error": "ไม่พบข้อมูลนี้"})
		return
	}

	c.JSON(200, gin.H{
		"data": users,
	})
}

func GetProfile(c *gin.Context) {
	user := c.MustGet("user")
	c.JSON(http.StatusOK, gin.H{
		"data": user,
	})
}
