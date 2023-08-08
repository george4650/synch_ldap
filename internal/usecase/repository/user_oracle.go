package repository

import (
	"context"
	"fmt"
	"myapp/internal/model"

	go_oracle "gitlabnew.nextcontact.ru/r.alfimov/go-oracle"
)

func NewUserOracleDB(ora *go_oracle.Oracle) *UserOracle {
	return &UserOracle{ora}
}

type UserOracle struct {
	*go_oracle.Oracle
}

func (r *UserOracle) CreateUsers(ctx context.Context, users []model.User) error {

	/*sqlText := `INSERT INTO QUEUES (NAME, MUSICONHOLD, STRATEGY, TIMEOUT, ANNOUNCE_FREQUENCY, ANNOUNCE_ROUND_SECONDS, RETRY, WRAPUPTIME,
					   MAXLEN, SERVICELEVEL, JOINEMPTY, LEAVEWHENEMPTY, REPORTHOLDTIME, MEMBERDELAY, WEIGHT, TIMEOUTRESTART)
	VALUES (:name, :music, :strategy, 45, 0, 0, 5, 3, 0, 60, 0, 'no', 0, 0, 0, 0)`
	*/

	sqlText := ``

	values := []interface{}{users}

	if err := go_oracle.InsertMany(ctx, r.Oracle, sqlText, values); err != nil {
		return fmt.Errorf("Oracle - CreateUsers - oracle.Insert: %w", err)
	}

	return nil
}
