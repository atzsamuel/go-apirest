package entity

// object for REST(CRUD)

type Student struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
	Year      int    `json:"year"`
}
