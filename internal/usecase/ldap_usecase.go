package usecase

import (
	"context"
	"database/sql"
	"fmt"
	"math/rand"
	"myapp/config"
	"myapp/internal/model"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"

	"github.com/go-ldap/ldap/v3"
	"github.com/google/uuid"
)

type LdapUseCases struct {
	databaseName        string
	postgresStore       UserPostgresStrore
	employeeOracleStore EmployeeOracleStrore
	sippeersOracleStore SippeersOracleStrore
	userOracleStore     UserOracleStrore
	l                   *ldap.Conn
	cfg                 config.Config
}

func NewLdapUsecases(postgresStore UserPostgresStrore, employeeOracleStore EmployeeOracleStrore, sippeersOracleStore SippeersOracleStrore, userOracleStore UserOracleStrore, databaseName string, l *ldap.Conn, cfg config.Config) *LdapUseCases {
	return &LdapUseCases{
		databaseName:        databaseName,
		postgresStore:       postgresStore,
		employeeOracleStore: employeeOracleStore,
		sippeersOracleStore: sippeersOracleStore,
		userOracleStore:     userOracleStore,
		l:                   l,
		cfg:                 cfg,
	}
}

var (
	ldap_user_search_filters = "(&(objectCategory=person)(objectClass=user)(!(userAccountControl:1.2.840.113556.1.4.803:=2))(memberOf:1.2.840.113556.1.4.1941:=cn=NextPhone,cn=builtin,dc=nextcontact,dc=ru))"
	ldap_attributes          = []string{
		"sAMAccountName",  // Логин
		"whencreated",     // Дата создания
		"name",            // Полное ФИО
		"givenName",       // Имя + Отчество
		"sn",              // Фамилия
		"l",               // город
		"telephoneNumber", // мобильный телефон
		"Title",           // должность
		"Department",      // отдел
		"mail",            // электронная почта
		"StreetAddress",   // адрес
	}
)

func (us *LdapUseCases) UpdateData() error {

	searchRequest := ldap.NewSearchRequest(

		us.cfg.LDAPDomains,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		ldap_user_search_filters,
		ldap_attributes,
		nil,
	)

	searchReq, err := us.l.Search(searchRequest)
	if err != nil {
		return fmt.Errorf("LdapUseCases - UpdateData - us.l.Search: %w", err)
	}
	if len(searchReq.Entries) == 0 {
		return fmt.Errorf("LdapUseCases - UpdateData - us.l.Search: %s", "нет данных")
	}

	switch us.databaseName {
	case "Oracle":

		// Employees
		log.Info().Msg("Статистика по employee:")

		// Cписок всех employee из AD
		employeesAD := setDataIntoEmployees(searchReq)
		log.Info().Msgf("Получено данных из AD: %v", len(employeesAD))

		// Получаем список всех логинов и телефонов employee из БД
		data, err := us.employeeOracleStore.ListEmployeeLoginPhone(context.Background())
		if err != nil {
			return fmt.Errorf("LdapUseCases - UpdateData - us.employeeOracleStore.ListEmployeeLogin: %w", err)
		}
		log.Info().Msgf("Количество Employee уже в базе: %v", len(data))

		var employeeSumAD []model.Employee

		// Делаем дополнительную проверку, ведь login или phone также могут повторяться в AD - вставить таких в таблицу мы не можем
		employeeLogin := make(map[string]bool, len(employeesAD))
		employeePhone := make(map[int64]bool, len(employeesAD))

		for _, e := range employeesAD {
			if employeeLogin[e.Login] == false && employeePhone[e.MobilePhone.Int64] == false {
				employeeLogin[e.Login] = true
				employeePhone[e.MobilePhone.Int64] = true
				employeeSumAD = append(employeeSumAD, e)
			} else if employeeLogin[e.Login] == false && employeePhone[e.MobilePhone.Int64] == true {
				employeeLogin[e.Login] = true
				e.MobilePhone = sql.NullInt64{}
				employeeSumAD = append(employeeSumAD, e)
			}
		}

		// Сохраняем тех employee которых ещё нет в базе
		var employeesToInsert []model.Employee

		// Очищаем с целью сохранения в мапу login и phone employee из базы
		employeeLogin = make(map[string]bool, len(data))
		employeePhone = make(map[int64]bool, len(data))

		// Добавляем в мапы логины и тел. employee из базы
		for _, d := range data {
			employeeLogin[d.Login] = true
			employeePhone[d.MobilePhone.Int64] = true
		}

		for _, e := range employeeSumAD {
			if employeeLogin[e.Login] == false && employeePhone[e.MobilePhone.Int64] == false {
				employeeLogin[e.Login] = true
				employeePhone[e.MobilePhone.Int64] = true
				employeesToInsert = append(employeesToInsert, e)
			} else if employeeLogin[e.Login] == false && employeePhone[e.MobilePhone.Int64] == true {
				employeeLogin[e.Login] = true
				e.MobilePhone = sql.NullInt64{}
				employeesToInsert = append(employeesToInsert, e)
			}
		}

		log.Printf("Готово к вставке в табличку employees: %d \n", len(employeesToInsert))

		if len(employeesToInsert) != 0 {
			err = us.employeeOracleStore.InsertEmployees(context.Background(), employeesToInsert)
			if err != nil {
				return fmt.Errorf("LdapUseCases - UpdateData - us.oracleStore.InsertEmployees: %w", err)
			}
		}

		log.Printf("Добавлено новых пользователей в табличку employees: %d \n\n", len(employeesToInsert))

		// Sippeers
		log.Info().Msg("Статистика по sippeers:")

		// Cписок всех sippeers из AD
		sippeersesAD := setDataIntoSippers(us.cfg.LDAPAsteriskHost, searchReq)
		log.Info().Msgf("Получено данных из AD: %v", len(sippeersesAD))

		// Получаем список всех логинов sippeers из БД
		dataSippeers, err := us.sippeersOracleStore.ListSippeersLogin(context.Background())
		if err != nil {
			return fmt.Errorf("LdapUseCases - UpdateData - us.sippeersOracleStore.ListSippeersLogin: %w", err)
		}
		log.Info().Msgf("Количество sippeers уже в базе: %v", len(dataSippeers))

		var sippeersSumAD []model.Sippers

		// Делаем дополнительную проверку, ведь login также могут повторяться в AD - вставить таких в таблицу мы не можем
		sippeersLogin := make(map[string]bool, len(sippeersesAD))
		for _, e := range sippeersesAD {
			if sippeersLogin[e.Name] == false {
				sippeersLogin[e.Name] = true
				sippeersSumAD = append(sippeersSumAD, e)
			}
		}

		// Очищаем с целью сохранения в мапу логинов sippeers из базы
		sippeersLogin = make(map[string]bool, len(dataSippeers))

		// Добавляем в мапы логины sippeers из базы
		for _, d := range dataSippeers {
			sippeersLogin[d.Name] = true
		}

		// Сохраняем тех sippeers которых ещё нет в базе
		var sippeersToInsert []model.Sippers
		for _, e := range sippeersSumAD {
			if sippeersLogin[e.Name] == false {
				sippeersLogin[e.Name] = true
				sippeersToInsert = append(sippeersToInsert, e)
			}
		}

		log.Printf("Готово к вставке в табличку sippeers: %d \n", len(sippeersToInsert))

		if len(sippeersToInsert) != 0 {
			err = us.sippeersOracleStore.InsertSippeerses(context.Background(), sippeersToInsert)
			if err != nil {
				return fmt.Errorf("LdapUseCases - UpdateData - us.sippeersOracleStore.InsertSippeerses: %w", err)
			}
		}

		log.Printf("Добавлено новых пользователей в табличку sippeers: %d \n", len(sippeersToInsert))

	case "Postgres":
		users := setDataIntoUsers(searchReq)
		err = us.postgresStore.CreateUsers(context.Background(), users)
		if err != nil {
			return fmt.Errorf("LdapUseCases - UpdateData - us..postgresStore.CreateUsers: %w", err)
		}
	}

	return nil
}

func setDataIntoSippers(asteriskHost string, searchReq *ldap.SearchResult) []model.Sippers {
	rand.Seed(time.Now().UnixNano())

	var Sipperses []model.Sippers

	for _, s := range searchReq.Entries {
		var Sippers model.Sippers
		for _, sippersAttr := range s.Attributes {

			switch sippersAttr.Name {
			case "sAMAccountName":
				Sippers.Name = sippersAttr.Values[0]
				Sippers.CallerId = sippersAttr.Values[0]
			} // switch end
		} // s.Attributes end

		if Sippers.Name == "" {
			continue
		}
		Sippers.Host = "dynamic"
		Sippers.Type = "friend"
		Sippers.Context = "from-internal"
		Sippers.Secret = generateRandomPassword()
		Sippers.Transport = "wss"
		Sippers.DtmFMode = "rfc2833"
		Sippers.Nat = "no"
		Sippers.Trustrprid = "yes"
		Sippers.CallCounter = "yes"
		Sippers.VidoSupport = "no"
		Sippers.SessionTimers = "refuse"
		Sippers.Qualify = "no"
		Sippers.SendRpid = "yes"
		Sippers.QualifyFreq = "60"
		Sippers.FaxDetect = "no"
		Sippers.CanReinvite = "no"
		Sippers.Avpf = "yes"
		Sippers.ForceAvp = "yes"
		Sippers.IceSupport = "yes"
		Sippers.RtcpMux = "yes"
		Sippers.Encryption = "yes"
		Sippers.DtlsEnable = "yes"
		Sippers.DtlsVerify = "fingerprint"
		Sippers.DtlsCertfile = fmt.Sprintf("/etc/letsencrypt/live/%s/cert.pem", asteriskHost)
		Sippers.DtlsPrivateKey = fmt.Sprintf("/etc/letsencrypt/live/%s/privkey.pem", asteriskHost)
		Sippers.DtlsSetup = "actpass"
		Sippers.DtlsRekey = "0"
		Sipperses = append(Sipperses, Sippers)
	}
	return Sipperses
}

func setDataIntoEmployees(searchReq *ldap.SearchResult) []model.Employee {

	cityList := map[string]int{
		"Смоленск":   1,
		"Волжский":   2,
		"Рязань":     3,
		"Энгельс":    4,
		"Волгодонск": 5,
		"Таганрог":   6,
		"Волгоград":  7,
		"Удалёнка 1": 8,
	}

	var employees []model.Employee

	for _, e := range searchReq.Entries {

		var city string
		var employee model.Employee

		for _, employeeAttr := range e.Attributes {
			employee.ID = uuid.NewString()

			switch employeeAttr.Name {
			case "sn":
				if employeeAttr.Values[0] != "" {
					employee.Surname = employeeAttr.Values[0]
				}
			case "givenName":

				fullName := employeeAttr.Values[0]
				if !strings.Contains(fullName, " ") && fullName != "" {
					employee.Name = fullName
				} else if strings.Contains(fullName, " ") {
					fullNameSplited := strings.Split(fullName, " ")
					employee.Name = fullNameSplited[0]
					employee.Patronymic.String = fullNameSplited[1]
					employee.Patronymic.Valid = true
				}

			//case "whenCreated":
			//	whenCreated := strings.Split(employeeAttr.Values[0], ".")
			//	date, err := time.Parse("20060102150405", whenCreated[0])
			//	if err != nil {
			//		log.Println(err)
			//	}
			//	employeeAttr.Values[0] = fmt.Sprint(date.Format("02.01.2006 15:04:05"))
			//	employeeAttr.CreatedAt = whenCreated[0]
			case "sAMAccountName":
				employee.Login = employeeAttr.Values[0]
			case "telephoneNumber":
				phone, err := strconv.Atoi(employeeAttr.Values[0])
				if err == nil {
					employee.MobilePhone.Int64 = int64(phone)
					employee.MobilePhone.Valid = true
				} else {
					employee.MobilePhone.Valid = false
				}

			case "department":
				employee.Department.String = employeeAttr.Values[0]
				employee.Department.Valid = true
			case "title":
				employee.Post.String = employeeAttr.Values[0]
				employee.Post.Valid = true
			case "l":
				city = employeeAttr.Values[0]
			//case "name":
			//
			//	fullName := employeeAttr.Values[0]
			//	if !strings.Contains(fullName, " ") && fullName != "" {
			//		employee.Name = fullName
			//	} else if strings.Contains(fullName, " ") {
			//		fullNameSplited := strings.Split(fullName, " ")
			//		employee.Name = fullNameSplited[0]
			//		employee.Surname = fullNameSplited[1]
			//		if len(fullNameSplited) > 2 {
			//			employee.Patronymic.String = fullNameSplited[2]
			//			employee.Patronymic.Valid = true
			//		}
			//	}
			case "mail":
				employee.Mail = employeeAttr.Values[0]
			case "streetaddress":
				employee.Address.String = employeeAttr.Values[0]
				employee.Address.Valid = true
			} // switch end
		} // e.Attributes end

		locationId := cityList[city]
		if locationId <= 0 || locationId > 8 {
			locationId = 1
		}

		employee.FidRole = 1
		employee.FidLocation = locationId
		employee.FidLanguage = 1

		if employee.Login == "" || employee.Name == "" || employee.Surname == "" {
			log.Printf("Пропущен пользователь с незаполненными основными данными: %v", employee)
			continue
		}
		employees = append(employees, employee)
	}
	return employees
}

func setDataIntoUsers(searchReq *ldap.SearchResult) []model.User {
	var Users []model.User

	for _, user := range searchReq.Entries {
		var User model.User
		for _, userAttr := range user.Attributes {

			User.UUID = uuid.NewString()
			switch userAttr.Name {
			case "sn":
				User.Surname = userAttr.Values[0]
			case "givenName":
				User.GivenName = userAttr.Values[0]
			case "whenCreated":

				whenCreated := strings.Split(userAttr.Values[0], ".")

				date, err := time.Parse("20060102150405", whenCreated[0])

				if err != nil {
					log.Err(err)
				}

				userAttr.Values[0] = fmt.Sprint(date.Format("02.01.2006 15:04:05"))

				User.CreatedAt = whenCreated[0]

			case "sAMAccountName":
				User.SAMAccountName = userAttr.Values[0]
			case "telephoneNumber":
				User.TelephoneNumber = userAttr.Values[0]
			case "department":
				User.Department = userAttr.Values[0]
			case "title":
				User.Title = userAttr.Values[0]
			case "l":
				User.City = userAttr.Values[0]
			case "mail":
				User.Mail = userAttr.Values[0]
			} // switch end

		}

		Users = append(Users, User)
	}
	fmt.Printf("\nall actual users count: %d\n\n", len(searchReq.Entries))
	return Users
}

func generateRandomPassword() string {

	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz")
	length := 16
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}
	return b.String()
}
