package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

type User struct {
	Name     string
	Password interface{}
}

var name, password string

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端cookie并校验
		if cookie, err := c.Cookie("define"); err == nil {
			if cookie == name {
				c.Next()
				return
			}
		}
		// 返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "err"})
		// 若验证不通过，不再调用后续的函数处理
		c.Abort()
		return
	}
}

func main() {
	userAll := make([]User, 1)
	file, _ := os.Create("userMassage")
	defer file.Close()
	//创建一个默认的路由引擎
	r := gin.Default()
	//注册
	r.POST("/register", func(c *gin.Context) {
		//获取数据
		name = c.PostForm("name")
		password = c.PostForm("password")
		//添加注册信息到文件
		userAll = append(userAll, User{
			Name:     name,
			Password: password,
		})
		marsha1, _ := json.Marshal(&userAll)
		file.Write(marsha1)
		//返回结果
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "注册成功",
		})
	})
	//登录
	r.POST("/login", func(c *gin.Context) {
		//获取参数
		c.PostForm("name")
		c.PostForm("password")
		//验证
		os.Open("userMassage")
		var outMassage []byte
		var outData []User
		file.Read(outMassage)
		json.Unmarshal(outMassage, &outData)
		is := true
		for _, v := range outData {

			if v.Name == name && v.Password == password {
				c.JSON(http.StatusOK, gin.H{
					"code":    200,
					"message": "登录成功",
				})
				is = false
				break
			}
		}
		if is {
			c.JSON(http.StatusOK, gin.H{
				"code":    404,
				"message": "登录失败",
			})
			return
		}
		//登录成功设置cookie
		c.SetCookie("define", name, 3600, "/", "localhost", false, true)

	})
	//调用中间件
	r.Use(AuthMiddleWare())
	r.Run()
}
