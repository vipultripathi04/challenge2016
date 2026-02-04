package model

// DistributorResponse represents the response containing distributor permissions
type DistributorResponse struct {
	Distributor string `json:"distributor"`
	Permission  string `json:"permissions"`
}

// ValidationResponse represents validation error responses
type ValidationResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
