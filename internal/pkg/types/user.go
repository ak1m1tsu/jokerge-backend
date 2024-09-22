package types

// User пользователь системы
type User struct {
	ID        string `bun:",pk"`
	Email     string
	FirstName string
	LastName  string
	Password  string
}
