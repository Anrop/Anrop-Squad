package db

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Arma Arma   `json:"arma"`
}

type Arma struct {
	ID   string `json:"id"`
	Nick string `json:"nick"`
}
