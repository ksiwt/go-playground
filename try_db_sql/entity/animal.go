package entity

type Animal struct {
	ID   uint64 `db:"id"`
	Name string `db:"name"`
}
