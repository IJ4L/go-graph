// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	InsertEmployee(ctx context.Context, arg InsertEmployeeParams) error
	InsertMenu(ctx context.Context, arg InsertMenuParams) error
	InsertUser(ctx context.Context, arg InsertUserParams) error
	RemoveEmployeeByID(ctx context.Context, id int32) error
	SelectAllEmployee(ctx context.Context) ([]SelectAllEmployeeRow, error)
	SelectAllMenu(ctx context.Context) ([]SelectAllMenuRow, error)
	SelectMenuByID(ctx context.Context, id int32) (SelectMenuByIDRow, error)
	SelectUserByEmail(ctx context.Context, email string) (Auth, error)
}

var _ Querier = (*Queries)(nil)