package profile

import (
	"context"
	"encoding/json"
	"net/http"
	"pinterest/domain"
	"strconv"

	"github.com/gorilla/mux"
	"go.uber.org/zap"
)

// Create product creates product using provided data
func (facade *ProductFacade) CreateProduct(w http.ResponseWriter, r *http.Request) {
	productInput := new(domain.Product)
	err := json.NewDecoder(r.Body).Decode(productInput)
	if err != nil {
		facade.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	productID, err := facade.productClient.CreateProduct(context.Background(), *productInput)

	if err != nil {
		// TODO: check for "Shop not found" errors
		facade.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	productIDOutput := domain.ProductIDResponse{ProductID: productID}
	responseBody, err := json.Marshal(productIDOutput)
	if err != nil {
		facade.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
}

// Edit product ...
func (facade *ProductFacade) EditProduct(w http.ResponseWriter, r *http.Request) {
	productInput := new(domain.Product)
	err := json.NewDecoder(r.Body).Decode(productInput)
	if err != nil {
		facade.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	productIDstr, passedID := vars[string(domain.IDKey)]
	if !passedID {
		facade.logger.Info("Could not get id from query params",
			zap.String("url", r.RequestURI),
			zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	productInput.Id, _ = strconv.ParseUint(productIDstr, 10, 64)

	err = facade.productClient.EditProduct(context.Background(), *productInput)

	if err != nil {
		facade.logger.Info(err.Error(), zap.String("url", r.RequestURI), zap.String("method", r.Method))
		switch err {
		case domain.ErrProductNotFound:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// Get product ...
func (facade *ProductFacade) GetProductByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	productIDStr, passedID := vars[string(domain.IDKey)]

	if !passedID {
		facade.logger.Info("Could not get id from query params",
			zap.String("url", r.RequestURI),
			zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	productID, _ := strconv.ParseUint(productIDStr, 10, 64)
	product, err := facade.productClient.GetProductByID(context.Background(), productID)
	if err != nil {
		facade.logger.Info(err.Error(),
			zap.String("url", r.RequestURI),
			zap.String("method", r.Method))
		if err == domain.ErrProductNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	productOutput := domain.ToProductOutput(product)
	responseBody, err := json.Marshal(productOutput)
	if err != nil {
		facade.logger.Info(err.Error(),
			zap.String("url", r.RequestURI),
			zap.String("method", r.Method))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(responseBody)
	return
}
