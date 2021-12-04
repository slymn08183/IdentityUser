package model

type LoginResponse struct {
	Token 			*string `json:"token"`
	RefreshToken 	*string `json:"refreshToken"`
}

type LoginReceive struct {
	Message	string	`json:"message"`
}
