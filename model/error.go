package model

type Error struct {
	Message	string	`json:"message"`
}

type ErrorEnvelope struct {
	Error Error	`json:"error"`
}

func (e Error) GetAsEnvelope() ErrorEnvelope {
	return ErrorEnvelope{
		Error: e,
	}
}