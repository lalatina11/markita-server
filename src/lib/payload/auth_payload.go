package payload

type RegisterAdditional struct {
	Display_name string `json:"display_name"`
	Role         string `json:"role"`
}

type RegisterPayload struct {
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Data     RegisterAdditional `json:"data"`
}

type SignInPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
