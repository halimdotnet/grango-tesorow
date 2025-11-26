package dto

type AccountTypeResponse struct {
	ID        int    `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	DCPattern string `json:"dc_pattern"`
}

type AccountCategoryResponse struct {
	ID                   int     `json:"id"`
	Code                 string  `json:"code"`
	Name                 string  `json:"name"`
	Classification       *string `json:"classification,omitempty"`
	IsActive             bool    `json:"is_active"`
	AccountTypeCode      string  `json:"account_type_code"`
	AccountTypeName      string  `json:"account_type_name"`
	AccountTypeDCPattern string  `json:"account_type_dc_pattern"`
}
