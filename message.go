package main

import (
	"encoding/xml"
	"time"
)

// 接受到的消息
type receiveMsg struct {
	ToUserName   string `xml:"ToUserName"`
	FromUserName string `xml:"FromUserName"`
	CreateTime   int    `xml:"CreateTime"`
	MsgType      string `xml:"MsgType"`
	Content      string `xml:"Content"`
	MsgId        int    `xml:"MsgId"`
}

// 发送的消息
type sendMsg struct {
	XMLName      xml.Name `xml:"xml"`
	ToUserName   string
	FromUserName string
	CreateTime   int64
	MsgType      string
	Content      string
}

// 构造返回数据格式
func responseText(msg receiveMsg, responseContent string) sendMsg {
	return sendMsg{
		ToUserName:   msg.FromUserName,
		FromUserName: msg.ToUserName,
		CreateTime:   time.Now().Unix(),
		MsgType:      "text",
		Content:      responseContent,
	}
}
