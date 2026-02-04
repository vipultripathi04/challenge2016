package repository

import (
	"encoding/csv"
	"fmt"
	"golang/internal/model"
	"os"
	"path/filepath"
)

// CitiesRepository handles reading and managing cities data
type CitiesRepository struct {
	filePath string
}

// NewCitiesRepository creates a new CitiesRepository instance
func NewCitiesRepository(filePath string) *CitiesRepository {
	return &CitiesRepository{
		filePath: filePath,
	}
}

// ReadLocations reads and returns all locations from the CSV file
func (r *CitiesRepository) ReadLocations() ([]model.Location, error) {
	file, err := os.Open(r.filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open CSV file: %w", err)
	}
	defer file.Close()

	var locationsArray []model.Location
	reader := csv.NewReader(file)

	// Skip headers
	if _, err := reader.Read(); err != nil {
		return nil, fmt.Errorf("failed to read CSV headers: %w", err)
	}

	for {
		row, err := reader.Read()
		if err != nil {
			break
		}

		if len(row) > 5 && row[0] != "" {
			location := model.Location{
				City:     row[3],
				Province: row[4],
				Country:  row[5],
			}
			locationsArray = append(locationsArray, location)
		}
	}

	return locationsArray, nil
}

// GetCitiesFilePath returns the path to the cities CSV file
func GetCitiesFilePath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to get working directory: %w", err)
	}
	return filepath.Join(dir, "data", "cities.csv"), nil
}
