package main

import (
	"fmt"
	"os"
)

// a.使用os.Create函数，在你的项目目录下创建一个"plan.txt"文件，
//
// b.使用file.Write将"I’m not afraid of difficulties and insist on learning programming",写入"plan.txt"中。
//
// c.使用file.Read函数读取"plan.txt"的内容，并打印到控制台
func main() {
	file, _ := os.Create("plan.txt")
	file.Write([]byte("I’m not afraid of difficulties and insist on learning programming"))
	file, _ = os.Open("plan.txt")
	buf := make([]byte, 100)
	file.Read(buf)
	fmt.Println(string(buf))
}
