package db

// User database model
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Arma Arma   `json:"arma"`
}

// Arma database model
type Arma struct {
	ID   string `json:"id"`
	Nick string `json:"nick"`
}
