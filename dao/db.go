package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"sync"
)

var instance *gorm.DB
var once sync.Once

func Instance() *gorm.DB {
	once.Do(func() {
		//dbConfig := viper.Get("database.db")
		//fmt.Println(dbConfig)

		db, err := gorm.Open("mysql", "homestead:secret@tcp(127.0.0.1:33060)/blog")
		if err != nil {
			log.Fatal("failed to connect database:", err)
		}
		fmt.Println("连接了一次...")
		instance = db
	})
	return instance
}
