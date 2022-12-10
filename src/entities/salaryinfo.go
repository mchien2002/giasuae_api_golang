package entities

type Salaryinfo struct {
	ID             int         `json:"id"`
	Type_teacher   int         `json:"type_teacher"`
	Two_sessions   string      `json:"two_sessions"`
	Three_sessions string      `json:"three_sessions"`
	Four_sessions  string      `json:"four_sessions"`
	Five_sessions  string      `json:"five_sessions"`
	ID_category    interface{} `json:"id_category"`
	Created_at     string      `json:"createdAt"`
}
