package graph

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"github.com/LonelySakura/surely/graph/model"
)


type Storage struct{
	posts     map[string]*model.Post
	comments  map[string]*model.Comment

    childrenMap   map[string][]string
    postComments  map[string][]string
}

func NewStorage() *Storage {
	return &Storage{
		posts:    make(map[string]*model.Post),
		comments: make(map[string]*model.Comment),
		
		childrenMap: make(map[string][]string),
		postComments: make(map[string][]string),
	}
}

type Resolver struct{
	storage Storage
}

func NewResolver() *Resolver {
	return &Resolver{
		storage: *NewStorage(),
	}
}
