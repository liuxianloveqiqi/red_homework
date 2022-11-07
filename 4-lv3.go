package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Create("user.date")
	defer file.Close()
	for {
		fmt.Println("请开始选择")
		fmt.Println("1.开始注册")
		fmt.Println("2.退出")
		var choose int
		fmt.Scan(&choose)
		if choose == 2 {
			break
		}
		if choose == 1 {
			var un, pw string
			symbol := ", .:!/;" //分隔符
			for {
				fmt.Println("请输入你的用户名")
				fmt.Scan(&un)
				buf := make([]byte, 999999999)
				file, _ := os.Open("user.date")
				file.Read(buf)
				if strings.Contains(string(buf), un) || strings.Contains(symbol, pw) {
					fmt.Println("输入的用户名重复或包含分隔符号，请重新输入")
				} else {
					fmt.Println("输入用户名成功")
					break
				}
			}
			for {
				fmt.Println("请输入你的密码(至少6位数)")
				fmt.Scan(&pw)
				if len(pw) < 6 || strings.Contains(symbol, pw) {
					fmt.Println("输入的密码少于6位或包含分隔符号，请重新输入")
				} else {
					fmt.Println("输入密码成功")
					break
				}
			}
			message := map[string]string{
				"User name": un,
				"password":  pw}
			marsha1, err := json.Marshal(&message)
			if err != nil {
				return
			}
			file.Write(marsha1)
		}

	}
}
