package model

// LoginInfo - to submit login payload
type LoginInfo struct {
	Username string
	Password string
}

type LoginResponse struct {
	Success bool
	UUID    string
	Message string
}

// RegistrationInfo -  submit registration info
type RegistrationInfo struct {
	FullName string
	Status   string
	Username string
	Password string
}
