package dto

type AccountTypeResponse struct {
	ID        int    `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	DCPattern string `json:"dc_pattern"`
}
