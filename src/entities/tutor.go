package entities

type Tutor struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Address       string    `json:"address"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	School        string    `json:"school"`
	Department    string    `json:"department"`
	Teach_areas   string    `json:"teach_areas"`
	Gender        string    `json:"gender"`
	Graduate_year string    `json:"graduateYear"`
	Isnow         string    `json:"isnow"`
	Describe      string    `json:"describe"`
	Sobuoi        int       `json:"sobuoi"`
	Birth_year    string    `json:"birthYear"`
	Created_at    string    `json:"createdAt"`
	Classes       []Class   `json:"classes"`
	Subjects      []Subject `json:"subjects"`
}
