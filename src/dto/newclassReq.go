package dto

type Newclasses struct {
	ID         int    `json:"_id"`
	Address    string `json:"address"`
	District   string `json:"distrist"`
	Sobuoi     int    `json:"sobuoi"`
	Time       string `json:"time"`
	Salary     int    `json:"salary"`
	Require    string `json:"require"`
	Status     int    `json:"status"`
	Contact    string `json:"contact"`
	Created_at string `json:"createdAt"`
	Subjects   []int
}
