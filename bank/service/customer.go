package service

type CustomerResponse struct {
	CustomerID int    `json:"customer_id"`
	Name       string `json:"name"`
	Status     int    `json:"status"`
}

type CustomerService interface {
	Getcustomers() ([]CustomerResponse, error)
	Getcustomer(int) (*CustomerResponse, error)
}
