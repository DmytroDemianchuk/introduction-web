package horoshop

type Auth struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
type Domain struct {
	URL   string `json:"domain"`
	Token string `json:"token"`
}
