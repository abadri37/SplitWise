package entity

type User struct {
	ID    string
	Name  string
	Email string
	Bank  string
}

type Contribution struct {
	UserID string
	Amount float64
	Paid   float64
}
