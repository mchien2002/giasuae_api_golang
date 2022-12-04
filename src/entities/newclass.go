package entities

type NewClass struct{

}

type ClassesOfNewClass struct{
	ID_class int `json:"id_class"`
	ID_newclass int `json:"id_newclass"`
}

type SubjectsOfNewClass struct{
	ID_subject int `json:"id_class"`
}

type CategoresOfNewClass struct{

}