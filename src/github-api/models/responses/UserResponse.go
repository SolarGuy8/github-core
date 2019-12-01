package responses

type RegisterUser struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type User struct {
	Username string `json:"username"`
}

type Users struct {
	Array []User
}
