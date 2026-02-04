package model

// DistributorRequest represents the incoming request for distributor permissions
type DistributorRequest struct {
	Distributor string     `json:"distributor"`
	Include     []Location `json:"include"`
	Exclude     []Location `json:"exclude,omitempty"`
	Locations   []Location `json:"locations"`
}

// Location represents a geographic location
type Location struct {
	City     string `json:"city,omitempty"`
	Province string `json:"province,omitempty"`
	Country  string `json:"country"`
}
