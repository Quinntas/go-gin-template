package createUser

type DTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewDTO(email string, password string) DTO {
	return DTO{
		Email:    email,
		Password: password,
	}
}

type ResponseDTO struct {
	InsertId int64
	Message  string
}

func NewResponseDTO(insertId int64, message string) ResponseDTO {
	return ResponseDTO{
		InsertId: insertId,
		Message:  message,
	}
}
