package handler

import (
	"encoding/json"
	"golang/internal/model"
	"golang/internal/service"
	"net/http"
)

// DistributorHandler handles HTTP requests related to distributors
type DistributorHandler struct {
	service service.Service
}

// NewDistributorHandler creates a new DistributorHandler instance
func NewDistributorHandler(svc service.Service) *DistributorHandler {
	return &DistributorHandler{
		service: svc,
	}
}

// GetPermissions handles the POST request for distributor permissions
func (h *DistributorHandler) GetPermissions(w http.ResponseWriter, req *http.Request) {
	errorsArr := make([]model.ValidationResponse, 0)
	distributorsReq := make([]model.DistributorRequest, 0)

	// Decode the request body
	errorReq := json.NewDecoder(req.Body).Decode(&distributorsReq)
	if errorReq != nil {
		errorsArr = append(errorsArr, model.ValidationResponse{Code: "400", Message: errorReq.Error()})
		writeResponse(w, http.StatusBadRequest, errorsArr)
		return
	}

	// Call service layer
	response, err := h.service.GetDistributorPermissions(distributorsReq)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err)
	} else {
		writeResponse(w, http.StatusOK, response)
	}
}

// writeResponse writes the HTTP response with appropriate headers
func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}
