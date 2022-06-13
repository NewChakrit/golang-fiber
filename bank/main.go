package main

import (
	"bank/handler"
	"bank/logs"
	"bank/repository"
	"fmt"
	"net/http"
	"strings"
	"time"

	"bank/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

func main() {
	initTimeZone()
	initConfig()
	db := initDatabase()

	customerRepositoryDB := repository.NewCustomerRepositoryDB(db)
	customerRepositoryMock := repository.NewCustomerRepositoryMock()

	// _ = customerRepositoryDB
	_ = customerRepositoryMock

	customerService := service.NewCustomerService(customerRepositoryDB)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.Getcustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.Getcustomers).Methods(http.MethodGet)

	// log.Panicf("Banking service started at port %v", viper.GetInt("app.port"))
	logs.Info("Banking service started at port " + viper.GetString("app.port"))
	http.ListenAndServe(fmt.Sprintf(":%v", viper.GetInt("app.port")), router)

	// customers, err := customerService.Getcustomers()
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(customers)

	// customer, err := customerService.Getcustomer(2000)
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(customer)

	// _ = customerRepository

	// // ----- Get All Customers -----
	// // customers, err := customerRepository.GetAll()
	// // if err != nil {
	// // 	panic(err)
	// // }
	// // fmt.Println(customers)

	// // ----- Get a Customer -----
	// customer, err := customerRepository.GetById(2000)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(customer)
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) // Change config yaml
	// ex cli = "APP_PORT=5000"
	// ex cli = "DB_DRIVER=mongodb"

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func initTimeZone() {
	ict, err := time.LoadLocation("Asia/Bangkok")
	if err != nil {
		panic(err)
	}

	time.Local = ict
}

func initDatabase() *sqlx.DB {
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?parseTime=true",
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.database"),
	)
	db, err := sqlx.Open(viper.GetString("db.driver"), dsn)

	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(3 * time.Minute)
	db.SetMaxIdleConns(10)
	db.SetMaxIdleConns(10)

	return db
}
