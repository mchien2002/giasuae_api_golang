package dto

type AccountWithToken struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"username"`
	Password string `json:"password"`
	Role     int    `json:"role"`
	Token    string `json:"token"`
}
