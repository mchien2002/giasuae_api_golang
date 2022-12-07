package dto

type NewclasssesSet struct {
	ID         int    `json:"id"`
	Address    string `json:"address"`
	District   string `json:"district"`
	Sobuoi     int    `json:"sobuoi"`
	Time       string `json:"time"`
	Salary     int    `json:"salary"`
	Require    string `json:"require"`
	Status     int    `json:"status"`
	Contact    string `json:"contact"`
	Created_at string `json:"createdAt"`
}

type NewClassesReq struct {
	ID         int    `json:"id"`
	Address    string `json:"address"`
	District   string `json:"district"`
	Sobuoi     int    `json:"sobuoi"`
	Time       string `json:"time"`
	Salary     int    `json:"salary"`
	Require    string `json:"require"`
	Status     int    `json:"status"`
	Contact    string `json:"contact"`
	Created_at string `json:"createdAt"`
	Subjects   []int  `json:"subjects"`
}
