package types

type Product struct {
	ID          string `bun:",pk"`
	Name        string
	Description string
	Price       int
}
