package torq

type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	RoleID       string `json:"role_id"`
	Status       string `json:"status"`
	SsoProvision bool   `json:"sso_provision"`
}

type Role struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
