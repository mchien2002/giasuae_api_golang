package entities

type TutorDetail struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Address       string    `json:"address"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	School        string    `json:"school"`
	Department    string    `json:"department"`
	Teach_areas   string    `json:"teach_areas"`
	Gender        string    `json:"gender"`
	Graduate_year string    `json:"graduate_year"`
	Isnow         string    `json:"isnow"`
	Describe      string    `json:"describe"`
	Sobuoi        int       `json:"sobuoi"`
	Birth_year    string    `json:"birth_year"`
	ID_account    int       `json:"id_account"`
	Created_at    string    `json:"created_at"`
	Classes       []Class   `json:"classes"`
	Subjects      []Subject `json:"subjects"`
}

type TutorSet struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	School        string `json:"school"`
	Department    string `json:"department"`
	Teach_areas   string `json:"teach_areas"`
	Gender        string `json:"gender"`
	Graduate_year string `json:"graduateYear"`
	Isnow         string `json:"isnow"`
	Describe      string `json:"describe"`
	Sobuoi        int    `json:"sobuoi"`
	Birth_year    string `json:"birthYear"`
	ID_account    int    `json:"id_account"`
	Created_at    string `json:"created_at"`
	Classes       string `json:"classes"`
	Subjects      string `json:"subjects"`
}

type TutorReq struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	School        string `json:"school"`
	Department    string `json:"department"`
	Teach_areas   string `json:"teach_areas"`
	Gender        string `json:"gender"`
	Graduate_year string `json:"graduateYear"`
	Isnow         string `json:"isnow"`
	Describe      string `json:"describe"`
	Sobuoi        int    `json:"sobuoi"`
	Birth_year    string `json:"birthYear"`
	ID_account    int    `json:"id_account"`
	Created_at    string `json:"created_at"`
	Classes       []int  `json:"classes"`
	Subjects      []int  `json:"subjects"`
}

type TutorDefault struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	School        string `json:"school"`
	Department    string `json:"department"`
	Teach_areas   string `json:"teach_areas"`
	Gender        string `json:"gender"`
	Graduate_year string `json:"graduateYear"`
	Isnow         string `json:"isnow"`
	Describe      string `json:"describe"`
	Sobuoi        int    `json:"sobuoi"`
	Birth_year    string `json:"birthYear"`
	ID_account    int    `json:"id_account"`
	Created_at    string `json:"created_at"`
}

type ClassesOfTutor struct {
	ID_class int `json:"id_class"`
	ID_tutor int `json:"id_tutor"`
}

type SubjectsOfTutor struct {
	ID_subject int `json:"id_subject"`
	ID_tutor   int `json:"id_tutor"`
}
