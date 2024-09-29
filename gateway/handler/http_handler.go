package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
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
}

func (h *handler) HandleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var product []*pb.ProductOnly
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

	log.Printf("Received product: %+v", product)
	h.client.CreateProduct(r.Context(), &pb.CreateProductRequest{
		ProductOnly: product,
	})
}
