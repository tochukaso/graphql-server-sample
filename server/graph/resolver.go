package graph

import (
	"math/rand"
	"sync"

	"github.com/tochukaso/graphql-server/graph/model"
)

//go:generate go run github.com/99designs/gqlgen

type Resolver struct {
	ProductObservers map[string]chan *model.Product
	mu               sync.Mutex
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
