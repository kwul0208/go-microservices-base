package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	pb "github.com/kwul0208/common/api"
)

type handler struct {
	client pb.ProductServiceClient
}

func NewHandler(client pb.ProductServiceClient) *handler {
	return &handler{client}
}

func (h *handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/product", h.HandleCreateProduct)
	mux.HandleFunc("PUT /api/product", h.HandleUpdateProduct)
}

func (h *handler) HandleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var product *pb.ProductOnly
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &product)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	// log.Printf("Received product: %+v", product)
	h.client.CreateProduct(r.Context(), &pb.CreateProductRequest{
		ProductOnly: product,
	})

	// Berhasil, kirim response JSON dengan status created
	response := map[string]string{
		"status":  "success",
		"message": "Product created successfully",
	}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

func (h *handler) HandleUpdateProduct(w http.ResponseWriter, r *http.Request) {
	var product *pb.ProductOnly

	id := r.URL.Query().Get("id")

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &product)
	if err != nil {
		http.Error(w, "Failed to parse JSON", http.StatusBadRequest)
		return
	}

	h.client.UpdateProduct(r.Context(), &pb.UpdateProductRequest{
		ID:          id,
		ProductOnly: product,
	})

	response := map[string]string{
		"status":  "success",
		"message": "Product updated successfully",
	}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
