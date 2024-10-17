package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	pb "github.com/kwul0208/common/api"
)

type handler struct {
	client pb.ProductServiceClient
}

func NewHandler(client pb.ProductServiceClient) *handler {
	return &handler{client}
}

func (h *handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /api/products", h.HandlerGetProduct)
	mux.HandleFunc("GET /api/product/detail", h.HandlerGetProductById)
	mux.HandleFunc("PUT /api/product", h.HandleUpdateProduct)
	mux.HandleFunc("POST /api/product", h.HandleCreateProduct)
	mux.HandleFunc("DELETE /api/product/delete", h.HandleDeleteProduct)
}

func (h *handler) HandlerGetProduct(w http.ResponseWriter, r *http.Request) {

	product, _ := h.client.GetProducts(r.Context(), &pb.Empty{})

	response := map[string]interface{}{
		"status":  "success",
		"message": "getting product successfully",
		"data":    product,
	}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonResponse)
}

func (h *handler) HandlerGetProductById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)

	if err != nil {
		http.Error(w, "Failed parse int", http.StatusBadRequest)
		return
	}
	product, _ := h.client.GetProductById(r.Context(), &pb.ProductID{ID: id})

	response := map[string]interface{}{
		"status":  "success",
		"message": "getting product successfully",
		"data":    product,
	}
	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

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

	h.client.CreateProduct(r.Context(), &pb.CreateProductRequest{
		ProductOnly: product,
	})

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

	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, "Failed parse int", http.StatusBadRequest)
		return
	}

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

func (h *handler) HandleDeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.URL.Query().Get("id"), 10, 64)
	if err != nil {
		http.Error(w, "Failed parse int", http.StatusBadRequest)
		return
	}
	product, _ := h.client.DeleteProduct(r.Context(), &pb.ProductID{ID: id})

	response := map[string]interface{}{
		"status":  "success",
		"message": "delete product successfully",
		"data":    product,
	}

	jsonResponse, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
