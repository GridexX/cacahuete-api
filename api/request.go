package api

type UserCreationRequest struct {
	Email     string `json:"email" validate:"required"`
	Username  string `json:"username" validate:"required,min=4"`
	Password  string `json:"password" validate:"required,min=8"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	// Street     string `json:"street"`
	// PostalCode string `json:"postalCode"`
	// City       string `j`
}