package main

import (
	"bank/repository"
	"fmt"

	"bank/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:P@ssw0rd@tcp(13.73.163.73:3306)/banking")
	if err != nil {
		panic(err)
	}

	customerRepository := repository.NewCustomerRepositoryDB(db)
	customerService := service.NewCustomerService(customerRepository)

	customers, err := customerService.Getcustomers()
	if err != nil {
		panic(err)
	}

	fmt.Println(customers)

	customer, err := customerService.Getcustomer(2000)
	if err != nil {
		panic(err)
	}

	fmt.Println(customer)

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
