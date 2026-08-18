package payload

type SignUpAdditional struct {
	Display_name string `json:"display_name" validate:"required,min=2,max=100"`
	Role         string `json:"role" validate:"required,oneof=admin user"`
}

type SignUpPayload struct {
	Email    string           `json:"email" validate:"required,email"`
	Password string           `json:"password" validate:"required,min=8"`
	Data     SignUpAdditional `json:"data" validate:"required"`
}

type SignInPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}
