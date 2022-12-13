package api

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"sync"
)

var (
	sessions map[string]string
	mu       sync.Mutex
)

// 初始化
func init() {
	sessions = make(map[string]string)
}

// 生成session-id
func NewSessionID() string {
	// 生成一个新的 session-ID,这里先简单用随机数生成id
	b := make([]byte, 32)
	rand.Read(b)
	// 将随机字符串转换为十六进制表示的字符串
	s := hex.EncodeToString(b)
	// 输出随机密码看一下
	fmt.Println(s)
	return s
}

// 保存session并加锁
func SaveSession(id string, data string) error {
	// 保存session数据,用互斥锁防止资源竞争
	mu.Lock()
	defer mu.Unlock()
	sessions[id] = data
	return nil
}

// 读取session并加锁
func ReadSession(id string) (string, error) {
	mu.Lock()
	defer mu.Unlock()
	data, ok := sessions[id]
	if !ok {
		return "", errors.New("session没有找到")
	}
	// 返回读取到的session数据
	return data, nil
}
