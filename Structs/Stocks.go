package structs

type Stock struct {
	AlphaID     string  `json:"AlphaID"`
	CompanyName string  `json:"Name"`
	Value       float32 `json:"Value"`
}
