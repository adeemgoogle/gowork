package modell

type User struct {
	ID       int    `json:"ID"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}
