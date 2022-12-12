package api

import "net/http"

type Cookie struct {
	w http.ResponseWriter
	r *http.Request
}

func NewCookie(w http.ResponseWriter, r *http.Request) *Cookie {
	return &Cookie{
		w: w,
		r: r,
	}
}

// 设置cookie
func (c *Cookie) Set(name, value string) {
	http.SetCookie(c.w, &http.Cookie{
		Name:  name,
		Value: value,
	})
}

// 读取cookie
func (c *Cookie) Get(name string) (string, error) {
	cookie, err := c.r.Cookie(name)
	if err != nil {
		return "", err
	}
	return cookie.Value, nil
}

// 删除cookie
func (c *Cookie) Delete(name string) {
	http.SetCookie(c.w, &http.Cookie{
		Name:   name,
		MaxAge: -1, // 直接给负数把它删掉
	})
}
