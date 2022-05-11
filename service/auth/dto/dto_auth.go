package dto

// InputLogin ...
type InputLogin struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required,min=6"`
}

// AuthData ...
type DataAuth struct {
	ID       uint64
	Email    string
	Password string
}
