package User

type RequestLogin struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
