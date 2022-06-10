package main

import (
	"bank/handler"
	"bank/repository"
	"net/http"

	"bank/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Open("mysql", "root:P@ssw0rd@tcp(13.73.163.73:3306)/banking")
	if err != nil {
		panic(err)
	}

	customerRepositoryDB := repository.NewCustomerRepositoryDB(db)
	customerRepositoryMock := repository.NewCustomerRepositoryMock()

	_ = customerRepositoryDB
	customerService := service.NewCustomerService(customerRepositoryMock)
	customerHandler := handler.NewCustomerHandler(customerService)

	router := mux.NewRouter()
	router.HandleFunc("/customers", customerHandler.Getcustomers).Methods(http.MethodGet)
	router.HandleFunc("/customers/{customerID:[0-9]+}", customerHandler.Getcustomers).Methods(http.MethodGet)

	http.ListenAndServe(":8000", router)

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
