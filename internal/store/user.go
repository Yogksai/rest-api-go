package store

import (
	"github.com/Yogksai/rest-api-go/internal/controller"
	"github.com/Yogksai/rest-api-go/internal/repository"
)

type PostStore struct {
	UserRepo      repository.UserRepo
	UserContrller controller.RepoController
}
