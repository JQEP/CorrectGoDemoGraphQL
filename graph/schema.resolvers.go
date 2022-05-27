package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-graphql-mongodb-api/database"
	"go-graphql-mongodb-api/graph/generated"
	"go-graphql-mongodb-api/graph/model"

	_ "github.com/99designs/gqlgen/graphql/handler"
	_ "github.com/99designs/gqlgen/graphql/playground"
)

func (r *mutationResolver) CreateCourse(ctx context.Context, input model.NewCourse) (*model.Course, error) {
	return db.InsertCourseById(input), nil
}

func (r *queryResolver) Course(ctx context.Context, id string) (*model.Course, error) {
	return db.FindCourseById(id), nil
}

func (r *queryResolver) Courses(ctx context.Context) ([]*model.Course, error) {
	return db.AllCourses(), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

var db = database.Connect("mongodb://localhost:27017/")
