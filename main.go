package main

import (
	"encoding/xml"
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
		log.Printf("用户发送信息 %s", body)
		msg := receiveMsg{}
		xml.Unmarshal(body, &msg)

		if msg.MsgType != "text" {
			c.String(http.StatusOK, "")
		}
		word := getOneByName(msg.Content)
		sendContent := "你发送的数据为：" + msg.Content + word.desc
		c.XML(http.StatusOK, responseText(msg, sendContent))
	})

	route.Run(":8081") // listen and serve on 0.0.0.0:8080
}
