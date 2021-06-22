package graph

import (
	"sync"

	"github.com/tochukaso/graphql-server/graph/model"
)

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	ProductObservers map[string]chan *model.Product
	mu               sync.Mutex
}
