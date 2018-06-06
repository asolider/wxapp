package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "client:password@/chengyu?charset=utf8mb4")

	if err != nil {
		log.Printf("fail to connect mysql: %s", err)
	}
	Db.SetMaxOpenConns(50)
	Db.SetMaxIdleConns(20)
}

type word struct {
	item  string
	spell string
	desc  string
	from  string
	ps    string
}

func getOneByName(name string) (findWord word) {
	err := Db.QueryRow("select item,spell,desc,from,ps from word where item=?", name).Scan(&findWord.item, &findWord.spell, &findWord.desc, &findWord.from, &findWord.ps)
	if err == sql.ErrNoRows || err != nil {
		return nil
	}
	return
}
