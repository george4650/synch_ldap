package model

import "database/sql"

type GetEmployeeLoginPhone struct {
	Login       string        `db:"LOGIN"`
	MobilePhone sql.NullInt64 `db:"MOBILE_PHONE"`
}

type Employee struct {
	ID          string         `db:"name:ID"`
	FidRole     int            `db:"name:FID_ROLE"`
	FidLocation int            `db:"name:FID_LOCATION"`
	FidLanguage int            `db:"name:FID_LANGUAGE"`
	Name        string         `db:"name:NAME"`
	Surname     string         `db:"name:SURNAME"`
	Patronymic  sql.NullString `db:"name:PATRONYMIC"`
	Login       string         `db:"name:LOGIN"`
	Department  sql.NullString `db:"name:DEPARTMENT"`
	Post        sql.NullString `db:"name:POST"`
	MobilePhone sql.NullInt64  `db:"name:MOBILE_PHONE"`
	Mail        string         `db:"name:E_MAIL"`
	Address     sql.NullString `db:"name:ADDRESS"`
	DateOfBirth sql.NullTime   `db:"name:DATE_OF_BIRTH"`
	Comments    sql.NullString `db:"name:COMMENTS"`
}
