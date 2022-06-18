package handler

import (
	"bank/errs"
	"bank/service"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type customerHandler struct {
	custSrv service.CustomerService
}

func NewCustomerHandler(custSrv service.CustomerService) customerHandler {
	return customerHandler{custSrv: custSrv}
}

func (h customerHandler) Getcustomers(w http.ResponseWriter, r *http.Request) {
	customers, err := h.custSrv.Getcustomers()
	if err != nil {
		appErr, ok := err.(errs.AppError)
		if ok {
			w.WriteHeader(appErr.Code)
			fmt.Fprint(w, appErr.Message)
		}


		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customers)
}

func (h customerHandler) Getcustomer(w http.ResponseWriter, r *http.Request) {
	customerID, _ := strconv.Atoi(mux.Vars(r)["custmerID"])
	customer, err := h.custSrv.Getcustomer(customerID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, err)
		return
	}
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(customer)
}
