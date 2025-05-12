package entity

type ExpenseGroup struct {
	ID      string
	Name    string
	Members map[string]*User
}
