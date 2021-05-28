// Code generated by sqlc. DO NOT EDIT.
//go:generate  mockgen -source=./querier.go  -destination=./mock/querier.go

package gen

import (
	"context"
	"database/sql"
)

type Querier interface {
	CreateAuthor(ctx context.Context, arg CreateAuthorParams) (sql.Result, error)
	DeleteAuthor(ctx context.Context, id int32) error
	DeleteAuthorIn(ctx context.Context, id []int32) error
	GetAuthorsInCompany(ctx context.Context, arg GetAuthorsInCompanyParams) ([]Author, error)
	GetAuthorsInCompanyById(ctx context.Context, id []int32) ([]Author, error)
	GetAuthorsInOneCompany(ctx context.Context, id int32) ([]Author, error)
	GetCompanyName(ctx context.Context, companyname sql.NullString) (sql.NullString, error)
	GetOneAuthor(ctx context.Context, arg GetOneAuthorParams) (Author, error)
	InsertMulti(ctx context.Context, arg InsertMultiParams) (sql.Result, error)
	ListAllAuthors(ctx context.Context) ([]Author, error)
	ListAuthors(ctx context.Context, arg ListAuthorsParams) ([]Author, error)
	ListCompanyById(ctx context.Context, id []int32) ([]Company, error)
}

var _ Querier = (*Queries)(nil)