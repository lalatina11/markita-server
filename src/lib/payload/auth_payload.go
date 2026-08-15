package payload

type RegisterAdditional struct {
	display_name string
}

type RegisterPayload struct {
	Email    string             `json:"email"`
	Password string             `json:"password"`
	Data     RegisterAdditional `json:"data"`
}
