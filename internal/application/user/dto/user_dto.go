package dto

type UserDTO struct {
	ID        string  `json:"id"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	FullName  string  `json:"fullName"`
	Email     string  `json:"email"`
	CreatedAt string  `json:"createdAt"`
	UpdatedAt *string `json:"updatedAt,omitempty"`
}

type CreateUserDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
