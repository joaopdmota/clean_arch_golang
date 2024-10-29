package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.55

import (
	"context"

	"github.com/devfullcycle/20-CleanArch/internal/infra/graph/model"
	"github.com/devfullcycle/20-CleanArch/internal/usecase"
)

// CreateOrder is the resolver for the createOrder field.
func (r *mutationResolver) CreateOrder(ctx context.Context, input *model.OrderInput) (*model.Order, error) {
	dto := usecase.OrderInputDTO{
		ID:    input.ID,
		Price: float64(input.Price),
		Tax:   float64(input.Tax),
	}
	output, err := r.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &model.Order{
		ID:         output.ID,
		Price:      float64(output.Price),
		Tax:        float64(output.Tax),
		FinalPrice: float64(output.FinalPrice),
	}, nil
}

// GetOrders is the resolver for the getOrders field.
func (r *mutationResolver) GetOrders(ctx context.Context) ([]*model.Order, error) {
	result, err := r.GetOrdersUseCase.Execute()

	if err != nil {
		return nil, err
	}

	orders := make([]*model.Order, len(result.Orders))
	for i, orderEntity := range result.Orders {
		orders[i] = &model.Order{
			ID:         orderEntity.ID,
			FinalPrice: orderEntity.FinalPrice,
			Price:      orderEntity.Price,
			Tax:        orderEntity.Tax,
		}
	}

	return orders, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }