/*
 * @Author: GG
 * @Date: 2022-09-19 10:14:49
 * @LastEditTime: 2022-09-19 15:38:41
 * @LastEditors: GG
 * @Description:
 * @FilePath: \captcha\main.go
 *
 */
package main

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

// SESSION START
// session 配置
func sessionConfig() sessions.Store {
	sessionSecret := "golang-gin-captcha"
	sessionMaxAge := 3600
	store := cookie.NewStore([]byte([]byte(sessionSecret)))
	store.Options(sessions.Options{
		MaxAge: sessionMaxAge, // 秒
		Path:   "/",
	})
	return store
}

// 中间件
func session(sessionKey string) gin.HandlerFunc {
	store := sessionConfig()
	return sessions.Sessions(sessionKey, store)
}

// SESSION END
// CHPTCHA START
/*
 * 生成图片
 *
 * id  captchaId
 * ext 图片扩展名
 * lang 语言，视频用，图片不需要
 * download 是否下载
 * width 宽
 * height 高
 */
func serve(w http.ResponseWriter, r *http.Request, id, ext, lang string, download bool, width, height int) error {
	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Pragma", "no-cache")
	w.Header().Set("Expires", "0")

	var content bytes.Buffer
	switch ext {
	case ".png":
		w.Header().Set("Content-Type", "image/png")
		_ = captcha.WriteImage(&content, id, width, height)
	case ".wav":
		w.Header().Set("Content-Type", "audio/x-wav")
		_ = captcha.WriteAudio(&content, id, lang)
	default:
		return captcha.ErrNotFound
	}

	if download {
		w.Header().Set("Content-Type", "application/octet-stream")
	}
	http.ServeContent(w, r, id+ext, time.Time{}, bytes.NewReader(content.Bytes()))
	return nil
}

// 生成验证码路由
func Captcha(c *gin.Context, length ...int) {
	u := captcha.DefaultLen
	w, h := 107, 36
	l := len(length)
	if l >= 3 {
		h = length[2]
	}

	if l >= 2 {
		w = length[1]
	}

	if l >= 1 {
		u = length[0]
	}

	fmt.Printf("u: %v\n", u)
	fmt.Printf("w: %v\n", w)
	fmt.Printf("h: %v\n", h)
	captchaId := captcha.NewLen(u)
	session := sessions.Default(c)
	session.Set("chptcha", captchaId)
	_ = session.Save()
	_ = serve(c.Writer, c.Request, captchaId, ".png", "zh", false, w, h)
}

// 校验验证码
func VerifyCaptcha(c *gin.Context, code string) bool {
	session := sessions.Default(c)

	if captchaId := session.Get("chptcha"); captchaId != nil {
		session.Delete("chptcha")
		session.Save()
		if captcha.VerifyString(captchaId.(string), code) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}

}

// CHPTCHA END
func main() {
	e := gin.Default()
	e.LoadHTMLGlob("./*.html")
	e.Use(session("captcha-test"))
	e.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{})
	})
	e.GET("/captcha", func(c *gin.Context) {
		Captcha(c, 4)
	})
	e.POST("/captcha/verify", func(c *gin.Context) {
		code, ok := c.GetPostForm("code")
		if !ok {
			c.JSON(http.StatusOK, gin.H{"status": 1, "msg": "failed"})
			return
		}

		if VerifyCaptcha(c, code) {
			c.JSON(http.StatusOK, gin.H{"status": 1, "msg": "success"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": 1, "msg": "failed"})

		}
	})
	e.Run()
}
