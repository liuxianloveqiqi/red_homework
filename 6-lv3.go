package main

//数据库准备
/*
//用户信息表
create table usermassage
(
    username         varchar(255) null,
    password         varchar(255) null,
    Id               bigint       not null
        primary key,
    secretProtection varchar(255) null
);

*/
//留言板表
/* create table messageboard
(
    senderID    int auto_increment
        primary key,
    sendername  varchar(255)  null,
    receiveID   int           null,
    receivename varchar(255)  null,
    massage     varchar(9999) null
);*/
import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
	"strconv"
)

type User struct {
	ID               int
	Name             string
	Password         string
	SecretProtection string
}
type Messgae struct {
	ID      int
	Name    string
	Messgae string
}

var registerID, loginID, postID int
var registerName, registerPassword, registerSecretProtection string
var loginName, loginPassword, loginSecretProtection string
var message, postUerName string
var db *sql.DB

// 使用数据库
func initDB() (err error) {
	dsn := "root:xian712525@tcp(127.0.0.1:3306)/go_db?charset=utf8mb4"
	// open函数只是验证格式是否正确，并不是创建数据库连接
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}
	// 与数据库建立连接
	err2 := db.Ping()
	if err2 != nil {
		return err2
	}
	return nil
}

// 使用中间件
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取客户端cookie并校验
		if cookie, err := c.Cookie("login"); err == nil {
			if cookie == "yes" {
				c.Next()
			}
		} else {
			// 返回错误
			c.JSON(http.StatusUnauthorized, gin.H{"error": "没有登录"})
			// 若验证不通过，不再调用后续的函数处理
			c.Abort()
		}
	}
}

// 插入数据
func insertData(i int, n string, p string, ps string) {
	sqlStr := "insert into userMassage(ID,username,password,secretProtection) values (?,?,?,?)"
	r, err := db.Exec(sqlStr, i, n, p, ps)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	ID, err2 := r.LastInsertId()
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
		return
	}
	fmt.Printf("ID: %v\n", ID)
}

// 查询密保
func queryRowData() (string, string) {
	var n, p string
	sqlStr := "select ID,secretProtection from userMassage where ID=? and secretProtection=? "
	var u User
	err := db.QueryRow(sqlStr, loginID, loginSecretProtection).Scan(&u.Name, &u.Password)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	n = u.Name
	p = u.Password
	return n, p
}

// 用户名和密码登录查询
func queryManyData() bool {
	is := false
	sqlStr := "select username,password from usermassage"
	r, err := db.Query(sqlStr)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	defer r.Close()
	// 循环读取结果集中的数据
	for r.Next() {
		var u2 User
		err2 := r.Scan(&u2.Name, &u2.Password)
		if err2 != nil {
			fmt.Printf("err: %v\n", err2)
		}
		if u2.Name == loginName && u2.Password == loginPassword {
			is = true
			break
		}
	}
	return is
}

// 更新数据 用户消息
func updateData() {
	sqlStr := "insert into messageboard(senderID,sendername,receiveID,receivename,message) values (?,?,?,?,?)"
	r, err := db.Exec(sqlStr, loginID, loginName, postID, postUerName, message)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	i2, err2 := r.LastInsertId()
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
		return
	}
	fmt.Printf("i2: %v\n", i2)
}
func main() {
	//连接数据库
	err := initDB()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Println("连接成功")
	}
	fmt.Printf("db: %v\n", db)
	//创建一个默认的路由引擎
	r := gin.Default()
	//注册
	r.POST("/register", func(c *gin.Context) {
		//获取数据
		ri := c.PostForm("ID")
		registerID, _ = strconv.Atoi(ri)
		registerName = c.PostForm("name")
		registerPassword = c.PostForm("password")
		registerSecretProtection = c.PostForm("secretProtection")
		//添加注册信息到文件
		insertData(registerID, registerName, registerPassword, registerSecretProtection)
		//返回结果
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "注册成功",
		})
	})
	//登录界面
	r.GET("/login", func(c *gin.Context) {
		c.SetCookie("login", "yes", 60, "/", "localhost", false, true)
		// 返回信息
		c.String(200, "Login success!")
	})
	//登录验证
	r.POST("/login", func(c *gin.Context) {
		//获取参数
		loginName = c.PostForm("name")
		loginPassword = c.PostForm("password")
		//验证(只需要用户名和密码)
		is := queryManyData()
		if is {
			c.JSON(http.StatusOK, gin.H{
				"code":    200,
				"message": "登录成功",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code":    404,
				"message": "登录失败",
			})
			return
		}

	})
	//使用密保查询用户名和密码
	r.POST("/secret", func(c *gin.Context) {
		li := c.PostForm("ID")
		loginID, _ = strconv.Atoi(li)
		loginSecretProtection = c.PostForm("secretProtection")
		//使用密保
		na, pa := queryRowData()
		c.JSON(http.StatusOK, gin.H{
			"你的用户名为：": na,
			"你的密码为：":  pa,
		})
	})
	//留言板功能
	//发送留言
	r.POST("/send", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"状态": 200,
		})
		p := c.PostForm("postID")
		postID, _ = strconv.Atoi(p)
		postUerName = c.PostForm("postname")
		message = c.PostForm("message")
		fmt.Println("------------------------")
		fmt.Println(postID, postUerName, message)

		updateData()
	})
	//查看留言
	r.GET("/look", AuthMiddleWare(), func(c *gin.Context) {
		sqlStr := "select senderID,sendername,message from messageboard where receivename=?"
		fmt.Println(loginName)
		r, err := db.Query(sqlStr, loginName)
		fmt.Println(r)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			return
		}
		defer r.Close()
		// 循环读取结果集中的数据
		for r.Next() {
			var m Messgae
			err2 := r.Scan(&m.ID, &m.Name, &m.Messgae)
			if err2 != nil {
				fmt.Printf("err2: %v\n", err2)
				return
			}
			fmt.Println(m.ID, m.Name, m.Messgae)
			c.JSON(http.StatusOK, gin.H{
				"发送者ID":   m.ID,
				"发送者Name": m.Name,
				"发来的消息":   m.Messgae,
			})
		}
	})
	r.Run()
}
