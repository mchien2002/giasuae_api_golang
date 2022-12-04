package entities

type NewClass struct {
	ID         int       `json:"_id"`
	Address    string    `json:"address"`
	District   string    `json:"distrist"`
	Sobuoi     int       `json:"sobuoi"`
	Time       string    `json:"time"`
	Salary     int       `json:"salary"`
	Require    string    `json:"require"`
	Status     int       `json:"status"`
	Contact    string    `json:"contact"`
	Created_at string    `json:"createdAt"`
	Subjects   []Subject `json:"subjects"`
}

type ClassesOfNewClass struct {
	ID_class    int `json:"id_class"`
	ID_newclass int `json:"id_newclass"`
}

type SubjectsOfNewClass struct {
	ID_subject  int `json:"id_class"`
	ID_newclass int `json:"id_newclass"`
}

type CategoresOfNewClass struct {
	ID_category int `json:"id_category"`
	ID_newclass int `json:"id_newclass"`
}