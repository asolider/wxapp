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
		msg := requestMsg{}
		xml.Unmarshal(body, &msg)

		sendContent := "你发送的数据为：" + msg.Content
		c.XML(http.StatusOK, ponseText(msg, sendContent))
	})

	route.Run(":8081") // listen and serve on 0.0.0.0:8080
}
