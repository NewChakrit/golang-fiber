package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
		DryRun: false, //Test virsual order if true
	})
	if err != nil {
		panic(err)
	}

	// db.Migrator().CreateTable(Gender{}, Test{}, Customer{})
	// db.AutoMigrate(Gender{}, Test{})
	// db.Migrator().CreateTable(Test{})

	// CreateGender("xxxx")
	// GetGenders()
	// GetGender(1)
	// GetGenderByName("Male")
	// UpdateGender(4, "yyyy")
	// UpdateGender(4, "zzzz")
	// DeleteGender(4)

	// CreateTest(0, "Test1")
	// CreateTest(0, "Test2")
	// CreateTest(0, "Test3")

	// DeleteTest(3)
	// GetTests()

	// db.Migrator().CreateTable(Customer{})

	// CreateCustomer("Oum", 2)

	// GetCustomers()
	UpdateGender2(1, "Male")
}

type Customer struct {
	// gorm.Model
	ID       uint
	Name     string
	Gender   Gender
	GenderID uint
}

func CreateCustomer(name string, genderID uint) {
	customer := Customer{
		Name:     name,
		GenderID: genderID,
	}

	tx := db.Create(&customer)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println(customer)
}

func GetCustomers() {
	customers := []Customer{}
	// tx := db.Preload("Gender").Find(&customers) // preload
	tx := db.Preload(clause.Associations).Find(&customers) // preload every table
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	// fmt.Println(customers, "\n")
	for _, customer := range customers {
		fmt.Printf("%v|%v|%v\n", customer.ID, customer.Name, customer.Gender.Name)
	}
}

func CreateTest(code uint, name string) {
	test := Test{Code: code, Name: name}
	db.Create(&test)

}

func GetTests() {
	tests := []Test{}
	db.Find(&tests)
	for _, t := range tests {
		fmt.Printf("%v|%v\n", t.ID, t.Name)
	}
}

// func DeleteTest(id uint) {
// 	db.Delete(&Test{}, id) // soft delete
// }

func DeleteTest(id uint) {
	db.Unscoped().Delete(&Test{}, id) // resl delete
}

// ----- Create ------
func CreateGender(name string) {
	gender := Gender{Name: name}
	tx := db.Create(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}

	fmt.Println(gender)
}

// ----- Read ------

func GetGenders() { // All genders
	genders := []Gender{}
	// tx := db.Find(&genders)
	tx := db.Order("id").Find(&genders) // order by _
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}

	fmt.Println(genders)
}

func GetGenderByName(name string) { // All genders
	genders := []Gender{}
	// tx := db.Find(&genders)
	// tx := db.Order("id").Find(&genders, "name=?", name) // order by _  ==== or ====
	tx := db.Where("name=?", name).Find(&genders) // order by _
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}

	fmt.Println(genders)
}

func GetGender(id uint) {
	gender := Gender{}
	tx := db.First(&gender, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}

	fmt.Println(gender)
}

// ------ Update ------

func UpdateGender(id uint, name string) {
	gender := Gender{}
	tx := db.First(&gender, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	gender.Name = name
	tx = db.Save(&gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	GetGender(id)

}

func UpdateGender2(id uint, name string) {
	gender := Gender{Name: name}
	tx := db.Model(&Gender{}).Where("id=@myid", sql.Named("myid", id)).Updates(gender)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	GetGender(id)
}

//  ------ Delete ------

func DeleteGender(id uint) {
	tx := db.Delete(&Gender{}, id)
	if tx.Error != nil {
		fmt.Println(tx.Error)
		return
	}
	fmt.Println("Deleted")
	GetGender(id)
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

// Row SQL
// Manual SQL Query
