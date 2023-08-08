package main

import (
	"myapp/config"
	"myapp/internal/app"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func main() {

	var conf config.Config

	// Настройка логгера
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		TimeFormat: "02.01.2006 15:04:05",
	}
	log.Logger = log.Output(output)
	log.With().Timestamp().Logger()

	// Выбор базы данных
	pflag.String("database", "Oracle", "Choose database")
	pflag.Parse()
	err := viper.BindPFlags(pflag.CommandLine)
	if err != nil {
		log.Fatal().Msgf("Ошибка при парсинге флагов", err)
	}
	db := viper.GetString("database")

	// Configuration
	if _, err := os.Stat("./config/app.env"); err == nil {
		log.Info().Msg("Обнаружен локальный файл конфига. Грузим настройки из него")
		conf, err = config.LoadConfigFile("./config")
		if err != nil {
			log.Err(err)
		}
	} else {
		conf, _ = config.LoadConfig()
	}

	//log.Info().Msgf("%v", conf)
	
	// Run
	app.Run(db, conf)
}
