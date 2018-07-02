package graph

// User is a basic struct to represent a Graph entry
type User struct {
	ID      int
	friends []int
}

// New return a user with a initialized friends map
func New(id int) *User {
	return &User{ID: id}
}
