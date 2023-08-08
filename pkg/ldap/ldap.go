package pkg

import (
	"fmt"
	"log"
	"myapp/config"
	"strings"

	"github.com/go-ldap/ldap/v3"
)

func ConnectToServer(cfg config.Config) (*ldap.Conn, error) {

	var (
		l   *ldap.Conn
		err error
	)

	servers := strings.Split(cfg.LDAPServers, ",")

	for _, server := range servers {

		l, err = ldap.DialURL(server)

		if err != nil {
			log.Printf("Не удалось подключиться к серверу: %s, Ошибка: %s\n", server, err)
			continue
		} else {
			break
		}
	}

	if l == nil {
		return nil, fmt.Errorf("ldap - ConnectToServer - Connect: %s", "Не удалось подключиться ни к одному из адресов")
	}

	err = l.Bind(fmt.Sprintf("%s@nextcontact.ru", cfg.LDAPLogin), cfg.LDAPPassword)
	if err != nil {
		if err, ok := err.(*ldap.Error); ok {
			if err.ResultCode == ldap.LDAPResultInvalidCredentials {
				log.Println("invalid credentials")
				return nil, fmt.Errorf("ldap - ConnectToServer - Bind: %s", "invalid credentials")

			}
		}
		log.Printf(fmt.Sprintf("l.Bind: %s", err.Error()))
		return nil, fmt.Errorf("ldap - ConnectToServer - Bind: %w", err)
	}

	return l, nil
}
