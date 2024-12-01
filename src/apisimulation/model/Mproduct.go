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
	ProductName struct {
		Category string `json:"Category"`
	}
	Price struct {
		Category string `json:"Category"`
		ProductName string `json:"ProductName"`
	}
)
