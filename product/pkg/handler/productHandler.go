package handler

import (
	"context"
	"net/http"

	"github.com/kwul0208/product/pkg/db"
	"github.com/kwul0208/product/pkg/pb"
	"github.com/kwul0208/product/pkg/use_case"
)

type Server struct {
	H           db.Handler
	Product_guc use_case.ProductUseCaseGrpc
	pb.UnimplementedProductServiceServer
}

func (h *Server) FindAll(ctx context.Context, rr *pb.FindAllRequest) (*pb.FindAllResponse, error) {
	products, err := h.Product_guc.GetAll(rr)

	var outProduct []*pb.FindOneData
	for _, v := range products {
		var p pb.FindOneData
		p.Id = v.Id
		p.Name = v.Name
		p.Price = v.Price
		p.Stock = v.Stock

		outProduct = append(outProduct, &p)
	}

	if err != nil {
		return &pb.FindAllResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, err
	}

	return &pb.FindAllResponse{
		Status:   http.StatusOK,
		Products: outProduct,
	}, nil
}

func (h *Server) FindOne(ctx context.Context, rr *pb.FindOneRequest) (*pb.FindOneResponse, error) {
	// Fetch the product from the use case
	product, err := h.Product_guc.GetOne(rr)

	if err != nil {
		return &pb.FindOneResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, err
	}

	// Initialize the response and set data
	outProduct := &pb.FindOneResponse{
		Data: &pb.FindOneData{},
	}

	outProduct.Data.Id = product.Id
	outProduct.Data.Name = product.Name
	outProduct.Data.Price = product.Price
	outProduct.Data.Stock = product.Stock

	return &pb.FindOneResponse{
		Status: http.StatusOK,
		Data:   outProduct.Data,
	}, nil
}

func (h *Server) CreateProduct(ctx context.Context, rr *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {

	product, err := h.Product_guc.CreateProduct(rr)

	if err != nil {
		return &pb.CreateProductResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, err
	}

	return &pb.CreateProductResponse{
		Status: http.StatusCreated,
		Id:     product.Id,
	}, nil
}
