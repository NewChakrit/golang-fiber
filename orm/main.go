package main

import (
	"context"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type SqlLogger struct {
	logger.Interface
}

func (l SqlLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	sql, _ := fc()
	fmt.Printf("%v\n===================================\n", sql)
}

var db *gorm.DB

func main() {
	dsn := "root:Chanew25@tcp(127.0.0.1:3306)/New?parseTime=true"
	dial := mysql.Open(dsn)

	var err error
	db, err = gorm.Open(dial, &gorm.Config{
		Logger: &SqlLogger{},
		DryRun: true, //Test virsual order
	})
	if err != nil {
		panic(err)
	}

	// db.Migrator().CreateTable(Gender{})
	// db.AutoMigrate(Gender{}, Test{})
	// db.Migrator().CreateTable(Test{})

	CreateGender("Male")
}

func CreateGender(name string) {
	gender := Gender{Name: name}
	tx := db.Create(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}

	fmt.Println(gender)
}

type Gender struct {
	ID   uint
	Name string `gorm:"unique;size:10"`
}

type Test struct {
	gorm.Model
	Code uint `gorm:"comment: This is Code"`
	// Name string `gorm:"column:myname;type:varchar(50)"` // config column name, confix type
	Name string `gorm:"column:myname;size:20;unique;default:Hello;not null"`
}

func (t Test) TableName() string {
	return "MyTest"
}
