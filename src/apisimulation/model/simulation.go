package model

type (
	MdownPaymemt struct {
		Price int `json:"Price"`
		DP    int `json:"DP"`
		Tenor int `json:"Tenor"`
	}
	MsubmitPengajuan struct {
		AmountIntalment string `json:"AmountInstalment"`
		TypeProduct     string `json:"TypeProduct"`
		Instalment      string `json:"Instalment"`
		TotalAmopunt    string `json:"TotalAmount"`
	}
	ListAgreement struct {
		Agrement         string  `json:"AGREMENT"`
		AmountInstalment float64 `json:"AMOUNT_INSTALMENT"`
		Product          string  `json:"PRODUCT"`
		Instalment       string  `json:"INSTALMENT"`
		DueDate          string  `json:"DUE_DATE"` 
		//TotalAmount      float64 `json:"TOTAL_AMOUNT"`
	}
)
