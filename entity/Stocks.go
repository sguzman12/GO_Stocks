package entity

type Stock struct {
	AlphaID     string  `json:"alphaID"`
	CompanyName string  `json:"companyName"`
	Value       float32 `json:"value"`
}
