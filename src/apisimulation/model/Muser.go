package model

type (
	Register struct {
		Name     string `json:"Name"`
		Phone    string `json:"Phone"`
		Email    string `json:"Email"`
		Password string `json:"Password"`
	}
	Login struct {
		Phone    string `json:"Phone"`
		Password string `json:"Password"`
	}
	User struct {
		Id       int    `json:"Id"`
		Name     string `json:"Name"`
		Email    string `json:"Email"`
		Password string `json:"Password"`
		
	}
)
