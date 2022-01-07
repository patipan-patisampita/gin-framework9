package usercontroller

import "github.com/gin-gonic/gin"

func GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"data new": "users",
	})
}

func Register(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "register",
	})
}

func Login(c *gin.Context){
	c.JSON(200,gin.H{
		"data":"login",
	})
}

func GetById(c *gin.Context){
	c.JSON(200,gin.H{
		"data":"id",
	})
}

func SearchByFullName(c *gin.Context){
	fullname := c.Query("fullname")
	c.JSON(200,gin.H{
		"data":fullname,
	})
}
