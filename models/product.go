package models

// import "github.com/go-ozzo/ozzo-validation"

// Artist represents an artist record.
type Product struct {
	Id   int    `json:"id" db:"id"`
	Name string `json:"name" db:"name"`
}
