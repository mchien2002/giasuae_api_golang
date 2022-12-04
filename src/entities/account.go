package entities

type Account struct {
	ID         int    `json:"_id"`
	Name       string `json:"name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Role       int    `json:"role"`
	Created_at string `json:"createdAt"`
}
