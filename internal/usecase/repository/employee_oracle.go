package repository

import (
	"context"
	"database/sql"
	"fmt"
	"myapp/internal/model"

	"github.com/rs/zerolog/log"

	go_oracle "gitlabnew.nextcontact.ru/r.alfimov/go-oracle"
)

func NewEmployeeOracleDB(ora *go_oracle.Oracle) *EmployeeOracle {
	return &EmployeeOracle{ora}
}

type EmployeeOracle struct {
	*go_oracle.Oracle
}

func (r *EmployeeOracle) InsertEmployee(ctx context.Context, e model.Employee) error {

	sqlText := `INSERT INTO employees_test (ID, FID_ROLE, FID_LOCATION, FID_LANGUAGE, SURNAME, NAME, PATRONYMIC, LOGIN, 
		DEPARTMENT, POST, MOBILE_PHONE, E_MAIL, ADDRESS, DATE_OF_BIRTH, COMMENTS)
	VALUES (:ID, :FID_ROLE, :FID_LOCATION, :FID_LANGUAGE, :SURNAME, :NAME, :PATRONYMIC, :LOGIN, :DEPARTMENT, :POST, :MOBILE_PHONE, 
		:E_MAIL, :ADDRESS, :DATE_OF_BIRTH, :COMMENTS)`

	values := []interface{}{e.ID, e.FidRole, e.FidLocation, e.FidLanguage, e.Surname, e.Name, e.Patronymic, e.Login, e.Department, e.Post, e.MobilePhone,
		e.Mail, e.Address, e.DateOfBirth, e.Comments}

	if err := go_oracle.Insert(ctx, r.Oracle, sqlText, values); err != nil {
		return fmt.Errorf("Oracle - InsertEmployee - oracle.Insert: %w", err)
	}

	return nil
}

func (r *EmployeeOracle) InsertEmployees(ctx context.Context, employees []model.Employee) error {

	sqlText := `INSERT INTO employees_test (ID, FID_ROLE, FID_LOCATION, FID_LANGUAGE, SURNAME, NAME, PATRONYMIC, LOGIN, 
		DEPARTMENT, POST, MOBILE_PHONE, E_MAIL, ADDRESS, DATE_OF_BIRTH, COMMENTS)
	VALUES (:ID, :FID_ROLE, :FID_LOCATION, :FID_LANGUAGE, :SURNAME, :NAME, :PATRONYMIC, :LOGIN, :DEPARTMENT, :POST, :MOBILE_PHONE, 
		:E_MAIL, :ADDRESS, :DATE_OF_BIRTH, :COMMENTS)`

	values := make([]interface{}, len(employees))
	for idx, e := range employees {
		values[idx] = []interface{}{e.ID, e.FidRole, e.FidLocation, e.FidLanguage, e.Surname, e.Name, e.Patronymic, e.Login, e.Department, e.Post, e.MobilePhone,
			e.Mail, e.Address, e.DateOfBirth, e.Comments}
	}

	if err := go_oracle.InsertMany(ctx, r.Oracle, sqlText, values); err != nil {
		return fmt.Errorf("Oracle - InsertEmployees - oracle.InsertMany: %w", err)
	}

	return nil
}

func (r *EmployeeOracle) ListEmployeeLoginPhone(ctx context.Context) ([]model.GetEmployeeLoginPhone, error) {

	sqlText := `SELECT login, mobile_phone FROM employees_test`

	employeeLogin, err := go_oracle.SelectMany[model.GetEmployeeLoginPhone](ctx, r.Oracle, sqlText, []interface{}{})
	if err != nil {
		if err == sql.ErrNoRows {
			log.Info().Msg("databse is empty")
			return nil, nil
		}
		return nil, fmt.Errorf("Oracle - ListEmployeeLogin - go_oracle.SelectMany: %w", err)
	}
	return employeeLogin, nil
}
