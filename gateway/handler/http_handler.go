package handler

import (
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
	mux.HandleFunc("POST api/product", h.HandleGetProduct)
}

func (h *handler) HandleGetProduct(w http.ResponseWriter, r *http.Request) {
	h.client.GetProduct(r.Context(), &pb.GetProductRequest{
		ProductId: r.PathValue("id"),
	})
}
