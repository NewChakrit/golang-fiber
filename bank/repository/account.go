package repository

type Account struct{
	AccountID int `db:"account_id"`
	CustomerID int `db:"customer_id"`
	OpeningDate int `db:"opening_date"`
	AccountType int `db:"account_type"`
	Amount int `db:"amount"`
	Status int `db:"status"`

}

type AccountRepository interface{
	Create(Account) (*Account, error)
	GetAll(int) ([]Account, error)
 }