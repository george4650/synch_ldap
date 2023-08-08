package config

import (
	"log"

	"github.com/spf13/viper"
)

type (
	Config struct {
		LDAPServers      string `mapstructure:"LDAP_SERVERS"`
		LDAPDomains      string `mapstructure:"LDAP_DOMAINS"`
		LDAPLogin        string `mapstructure:"LDAP_LOGIN"`
		LDAPPassword     string `mapstructure:"LDAP_PASSWORD"`
		LDAPAsteriskHost string `mapstructure:"LDAP_ASTERISKHOST"`

		OracleDbHost string `mapstructure:"DB_HOST"`
		OracleDbPort int    `mapstructure:"DB_PORT"`
		OracleDbTns  string `mapstructure:"DB_TNS"`
		OracleDbName string `mapstructure:"DB_NAME"`
		OracleDbPass string `mapstructure:"DB_PASS"`

		PostgresHost     string `mapstructure:"PG_HOST"`
		PostgresPort     int    `mapstructure:"PG_PORT"`
		PostgresUser     string `mapstructure:"PG_USER"`
		PostgresPassword string `mapstructure:"PG_PASS"`
		PostgresDbname   string `mapstructure:"PG_DBNAME"`
		PostgresSslmode  string `mapstructure:"PG_SSLMODE"`
	}
)

func LoadConfig() (config Config, err error) {

	viper.AutomaticEnv()

	config.LDAPServers = viper.GetString("LDAP_SERVERS")
	config.LDAPDomains = viper.GetString("LDAP_DOMAINS")
	config.LDAPLogin = viper.GetString("LDAP_LOGIN")
	config.LDAPPassword = viper.GetString("LDAP_PASSWORD")
	config.LDAPAsteriskHost = viper.GetString("LDAP_ASTERISKHOST")

	config.OracleDbHost = viper.GetString("DB_HOST")
	config.OracleDbPort = viper.GetInt("DB_PORT")
	config.OracleDbTns = viper.GetString("DB_TNS")
	config.OracleDbName = viper.GetString("DB_NAME")
	config.OracleDbPass = viper.GetString("DB_PASS")

	config.PostgresHost = viper.GetString("PG_HOST")
	config.PostgresPort = viper.GetInt("PG_PORT")
	config.PostgresUser = viper.GetString("PG_USER")
	config.PostgresPassword = viper.GetString("PG_PASS")
	config.PostgresDbname = viper.GetString("PG_DBNAME")
	config.PostgresSslmode = viper.GetString("PG_SSLMODE")

	return config, err
}

func LoadConfigFile(path string) (config Config, err error) {

	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal("Не удалось загрузить локальный config файл ", err)
	}

	err = viper.Unmarshal(&config)
	return config, nil
}
