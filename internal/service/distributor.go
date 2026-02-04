package service

import (
	"fmt"
	"golang/internal/model"
	"golang/internal/repository"
	"strings"
)

// DistributorService handles business logic for distributor permissions
type DistributorService struct {
	citiesRepo *repository.CitiesRepository
}

// NewDistributorService creates a new DistributorService instance
func NewDistributorService(citiesRepo *repository.CitiesRepository) *DistributorService {
	return &DistributorService{
		citiesRepo: citiesRepo,
	}
}

// Service interface defines the contract for distributor services
type Service interface {
	GetDistributorPermissions(req []model.DistributorRequest) ([]model.DistributorResponse, []model.ValidationResponse)
}

// GetDistributorPermissions is the main function that processes distributor requests
func (s *DistributorService) GetDistributorPermissions(req []model.DistributorRequest) ([]model.DistributorResponse, []model.ValidationResponse) {
	errorsArr := make([]model.ValidationResponse, 0)

	// Read all locations from CSV
	allLocations, err := s.citiesRepo.ReadLocations()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		errorsArr = append(errorsArr, model.ValidationResponse{Code: "400", Message: err.Error()})
		return nil, errorsArr
	}

	// Validate request and generate response
	finalResponse, errValidationResp := s.validateRequest(req, allLocations)
	if len(errValidationResp) > 0 {
		return nil, errValidationResp
	}

	return finalResponse, nil
}

// validateRequest validates the incoming requests and generates responses
func (s *DistributorService) validateRequest(req []model.DistributorRequest, allLocations []model.Location) ([]model.DistributorResponse, []model.ValidationResponse) {
	errorsArray := make([]model.ValidationResponse, 0)
	response := make([]model.DistributorResponse, 0)

	// Helper function to check if two locations match
	isRegionMatch := func(location model.Location, region model.Location) bool {
		if region.Country != "" && !strings.EqualFold(region.Country, location.Country) {
			return false
		}
		if region.Province != "" && !strings.EqualFold(region.Province, location.Province) {
			return false
		}
		if region.City != "" && !strings.EqualFold(region.City, location.City) {
			return false
		}
		return true
	}

	// Process each distributor request
	for _, distributor := range req {
		if distributor.Distributor == "" {
			errorsArray = append(errorsArray, model.ValidationResponse{
				Code:    "400",
				Message: "Distributor name is required",
			})
			continue
		}

		// Check if distributor is included in any location
		isIncluded := false
		for _, location := range allLocations {
			for _, includeRegion := range distributor.Include {
				if isRegionMatch(location, includeRegion) {
					isIncluded = true
					break
				}
			}
			if isIncluded {
				break
			}
		}

		// Check if distributor is excluded from any location
		isExcluded := false
		for _, location := range allLocations {
			for _, excludeRegion := range distributor.Exclude {
				if isRegionMatch(location, excludeRegion) {
					isExcluded = true
					break
				}
			}
			if isExcluded {
				break
			}
		}

		// Determine permission status
		permission := "DENY"
		if isIncluded && !isExcluded {
			permission = "ALLOW"
		}

		response = append(response, model.DistributorResponse{
			Distributor: distributor.Distributor,
			Permission:  permission,
		})
	}

	return response, errorsArray
}
