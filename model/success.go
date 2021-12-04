package model


type Success struct {
	Success	bool	`json:"success"`
}

func (s Success) True() Success {
	s.Success = true
	return s
}

func (s Success) False() Success {
	s.Success = false
	return s
}