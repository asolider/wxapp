package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	APPID     string = ""
	APPSECRET string = ""
)

func main() {
	route := gin.Default()
	route.GET("/api", func(c *gin.Context) {
		str := c.Query("echostr")

		log.Printf("%v", c.Request)
		log.Println(str)

		c.String(http.StatusOK, str)
	})

	route.POST("/api", func(c *gin.Context) {
		body, _ := c.GetRawData()
		msg := receiveMsg{}
		xml.Unmarshal(body, &msg)
		log.Printf("用户发送信息 %#v", msg)

		if msg.MsgType != "text" {
			c.String(http.StatusOK, "")
		}
		word := getOneByName(msg.Content)
		log.Printf("%#v", word)
		sendContent := "未查询到数据"
		if word.id != 0 {
			sendContent = fmt.Sprintf("成语：%s\n拼音：%s\n解释：%s\n出处：%s\n举例：%s", word.item, word.spell, word.desc, word.from, word.ps)
		}
		c.XML(http.StatusOK, responseText(msg, sendContent))
	})

	route.Run(":8081") // listen and serve on 0.0.0.0:8080
}
