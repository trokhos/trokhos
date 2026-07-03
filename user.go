package trokhos

type UserType uint

const (
	Admin UserType = iota
	Librarian
	Borrower
)

type User struct {
}
