package app

import (
	"fmt"
	"log"
	"myapp/config"
	"myapp/internal/usecase"
	"myapp/internal/usecase/repository"
	pkg "myapp/pkg/ldap"
	"myapp/pkg/oracle"
	"myapp/pkg/postgres"
)

func Run(database string, cfg config.Config) {

	//log.Println("cfg: ", cfg)
	//LDAP
	l, err := pkg.ConnectToServer(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	// Repository
	var (
		userPostgresRepository   *repository.Postgres
		sippeersOracleRepository *repository.SippeersOracle
		employeeOracleRepository *repository.EmployeeOracle
		userOracleRepository     *repository.UserOracle
	)

	switch database {
	case "Oracle":
		or, err := oracle.New(cfg)
		if err != nil {
			log.Fatal(fmt.Errorf("app - Run - Oracle.New: %w", err))
		}
		defer or.Close()

		employeeOracleRepository = repository.NewEmployeeOracleDB(or)
		sippeersOracleRepository = repository.NewSippeersOracleDB(or)
		userOracleRepository = repository.NewUserOracleDB(or)
	case "Postgres":

		pg, err := postgres.New(cfg)
		if err != nil {
			log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
		}
		defer pg.Close()

		userPostgresRepository = repository.NewPostgresDB(pg)
	default:
		log.Fatal("Не валидное значение базы данных")
	}

	// Use case
	us := usecase.NewLdapUsecases(userPostgresRepository, employeeOracleRepository, sippeersOracleRepository, userOracleRepository, database, l, cfg)

	if err := us.UpdateData(); err != nil {
		log.Fatal(err)
	}

}
