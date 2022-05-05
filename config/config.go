package config

import (
	"log"

	"github.com/ilyakaznacheev/cleanenv"
)

type AppConfig struct {
	AppName      string `env:"APP_NAME"`
	Port         string `env:"PORT" env-default:"8080"`
	DBConnection string `env:"DB_CONNECTION" env-default:"mysql"`
	DBDataBase   string `env:"DB_DATABASE" env-required:"true"`
	DBPort   	 string `env:"DB_PORT" env-default:"3306"`
	DBUser       string `env:"DB_USER" env-required:"true"`
	DBPassword   string `env:"DB_PASSWORD" env-required:"true"`
}

func LoadConfigration() AppConfig {
	config := AppConfig{}

	err := cleanenv.ReadConfig(".env", &config)

	if err != nil {
		log.Fatalf("Cannot load [.env] file [%v]", err)
	}

	return config
}