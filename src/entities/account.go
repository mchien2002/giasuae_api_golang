package entities

type Account struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"username"`
	Password   string `json:"password"`
	Role       int    `json:"role"`
	Created_at string `json:"createdAt"`
}
