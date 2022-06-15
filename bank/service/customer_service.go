package service

import (
	"bank/errs"
	"bank/logs"
	"bank/repository"
	"database/sql"
	"log"
	"net/http"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) customerService {
	return customerService{custRepo: custRepo}
}

func (s customerService) Getcustomers() ([]CustomerResponse, error) {
	customers, err := s.custRepo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	custResponses := []CustomerResponse{}
	for _, customer := range customers {
		custResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		custResponses = append(custResponses, custResponse)
	}

	return custResponses, nil
}

func (s customerService) Getcustomer(id int) (*CustomerResponse, error) {
	customer, err := s.custRepo.GetById(id)

	if err != nil {

		if err == sql.ErrNoRows {
			// return nil, errors.New("customer not found")
			return nil, errs.AppError{
				Code: http.StatusNotFound,
				Message: "Customer not found",
			}
		}

		// log.Println(err)
		logs.Error(err)
		return nil, err
	}

	custResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &custResponse, nil
}
