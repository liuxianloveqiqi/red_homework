package main

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name     string
	Password interface{}
}

var registerName, registerPassword, loginName, loginPassword string

// 使用中间件
func AuthMiddleWare() gin.HandlerFunc {

	return func(c *gin.Context) {
		// 获取客户端cookie并校验
		if cookie, err := c.Cookie("define"); err == nil {
			if cookie == loginName {
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
	//创建一个默认的路由引擎
	r := gin.Default()
	//注册
	r.POST("/register", func(c *gin.Context) {
		//获取数据
		registerName = c.PostForm("name")
		registerPassword = c.PostForm("password")
		//添加注册信息到文件
		userAll = append(userAll, User{
			Name:     registerName,
			Password: registerPassword,
		})
		marsha1, _ := json.Marshal(&userAll)

		file.Write(marsha1)
		defer file.Close()
		//返回结果
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "注册成功",
		})
	})
	//登录
	r.POST("/login", func(c *gin.Context) {
		//获取参数
		loginName = c.PostForm("name")
		loginPassword = c.PostForm("password")
		//验证
		userFile, _ := os.Open("userMassage")
		defer userFile.Close()
		//outMessage用于存放读取的文件内容，为什么需要make，因为outMessage是一个切片，需要初始化后才能存储，并且需要给予一定的容量
		var outMassage []byte = make([]byte, 4096)
		//用于存储Json解密后的数据
		var outJsonData []User
		//读取文件内容
		n, _ := userFile.Read(outMassage)
		//JSON解密成结构体切片，为什么此处需要使用[:n]，因为outMassage是一个Byte切片，含有多余的数据，n是读取的长度，所以需要使用[:n]来截取
		json.Unmarshal(outMassage[:n], &outJsonData)

		is := true
		for _, v := range outJsonData {

			if v.Name == loginName && v.Password == loginPassword {
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
		c.SetCookie("define", loginName, 3600, "/", "localhost", false, true)

	})
	//调用中间件
	r.Use(AuthMiddleWare())
	r.Run()
}
