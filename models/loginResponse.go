package models

// LoginResponse tiene el token que se devuelve en el login
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}
