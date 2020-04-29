package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"log"
	"sync"
	"time"
)

var instance *gorm.DB
var once sync.Once

func Instance() *gorm.DB {
	once.Do(func() {
		dbConfig := viper.GetStringMapString("database.db")
		connAddr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", dbConfig["user_name"], dbConfig["password"], dbConfig["host"], dbConfig["port"], dbConfig["db_name"])
		db, err := gorm.Open("mysql", connAddr)

		cuAt := time.Now().Unix()

		db.Callback().Update().Register("gorm:update_time_stamp", func(scope *gorm.Scope) {
			if _, ok := scope.Get("gorm:update_column"); !ok {
				_ = scope.SetColumn("UpdatedAt", cuAt)
			}
		})
		db.Callback().Create().Register("gorm:create_time_stamp", func(scope *gorm.Scope) {
			if !scope.HasError() {
				if createdAtField, ok := scope.FieldByName("CreatedAt"); ok {
					if createdAtField.IsBlank {
						_ = createdAtField.Set(cuAt)
					}
				}
				if updatedAtField, ok := scope.FieldByName("UpdatedAt"); ok {
					if updatedAtField.IsBlank {
						_ = updatedAtField.Set(cuAt)
					}
				}
			}
		})

		db.LogMode(viper.GetBool("debug"))

		if err != nil {
			log.Fatal("failed to connect database:", err)
		}
		fmt.Println("连接了一次...")
		instance = db
	})
	return instance
}
