package main

import "context"

type service struct {
	store ProductStore
}

func NewService(store ProductStore) *service {
	return &service{store}
}

func (s *service) CreateProduct(context.Context) error {
	return nil
}
