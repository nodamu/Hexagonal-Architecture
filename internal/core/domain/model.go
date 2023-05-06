package domain
type Message struct {
   ID   string `json:"id"`
   Body string `json:"body"`
}

type User struct {
   ID   int `json:"id"`
   Email    string `json:"email"`
	Password string `json:"password"`
	Membership bool `json:"is_member"`
}