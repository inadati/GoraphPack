package gqlkit

import (
	"context"
	"os"
	"strings"

	"github.com/inadati/gqlkit/models"
	"github.com/jinzhu/gorm"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

var SECRET_KEY1 = os.Getenv("SECRET_KEY1")
var GORM_SETUP = strings.Trim(os.Getenv("GORM_SETUP"), "\"")

type Resolver struct{}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateTodo(ctx context.Context, input models.NewTodo) (*models.Todo, error) {
	db, _ := gorm.Open("postgres", GORM_SETUP)
	defer db.Close()
	panic("not implemented")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	panic("not implemented")
}
