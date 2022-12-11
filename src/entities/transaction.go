package entities

type Transactionhistories struct {
	ID         int         `json:"id"`
	Amount     int         `json:"amount"`
	Content    string      `json:"content"`
	ID_account interface{} `json:"id_account"`
	Status     int         `json:"status"`
	Created_at string      `json:"created_at"`
}
