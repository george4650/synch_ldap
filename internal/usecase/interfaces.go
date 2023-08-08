package usecase

import (
	"context"
	"myapp/internal/model"
)

type EmployeeOracleStrore interface {
	InsertEmployee(ctx context.Context, e model.Employee) error 
	InsertEmployees(ctx context.Context, employees []model.Employee) error
	ListEmployeeLoginPhone(ctx context.Context) ([]model.GetEmployeeLoginPhone, error)
}

type SippeersOracleStrore interface {
	InsertSippeers(ctx context.Context, sippers model.Sippers) error
	InsertSippeerses(ctx context.Context, sippeerses []model.Sippers) error
	ListSippeersLogin(ctx context.Context) ([]model.ListSippersLogin, error)
}

type UserOracleStrore interface {
	CreateUsers(ctx context.Context, users []model.User) error
}

type UserPostgresStrore interface {
	CreateUsers(ctx context.Context, users []model.User) error
}
