package dto

type AccountWithToken struct {
	ID         int    `json:"_id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Role       int    `json:"role"`
	Token      string `json:"token"`
}
