package model

type (
	AllLob struct {
		Id      string `json:"Id"`
		LobName string `json:"LobName"`
		Desc    string `json:"Desc"`
	}
	GetLob struct {
		Id      int `json:"Id"`
	}
)
