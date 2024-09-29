package main

import "context"

type ProductService interface {
	CreateProduct(context.Context) error
}

type ProductStore interface {
	Create(context.Context) error
}
