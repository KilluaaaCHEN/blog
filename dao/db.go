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

var dbInst *gorm.DB
var dbOnce sync.Once

type Db struct {
	Drive    string
	Host     string
	Port     int
	UserName string `mapstructure:"user_name"`
	Password string
	DbName   string `mapstructure:"db_name"`
}

func InstDB() *gorm.DB {
	dbOnce.Do(func() {
		var config Db
		if err := viper.UnmarshalKey("database.db", &config); err != nil {
			log.Fatal(err)
		}
		connAddr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", config.UserName, config.Password, config.Host, config.Port, config.DbName)

		db, err := gorm.Open("mysql", connAddr)
		if err != nil {
			log.Fatal("failed to connect database:", err)
		}

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

		fmt.Println("db connected...")
		dbInst = db
	})
	return dbInst
}
