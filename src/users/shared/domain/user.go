package domain

type User struct {
	ID       int     `json: "id,omitempty"`
	FirtName *string `json: "first_name"`
	LastName *string `json: "last_name"`
	Nickname *string `json: "nickname"`
}
