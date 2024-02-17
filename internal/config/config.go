package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Env               string
	ServerAddress     string
	ServerTimeout     time.Duration
	ServerIdleTimeout time.Duration
	DBHost            string
	DBPort            int
	DBUser            string
	DBPassword        string
	DBName            string
}

// временно
func Setenv() {
	os.Setenv("APP_ENV", "local")
	os.Setenv("APP_SERVERADDRESS", "localhost:8081")
	os.Setenv("APP_SERVERTIMEOUT", "4s")
	os.Setenv("APP_SERVERIDLETIMEOUT", "60s")
	os.Setenv("APP_DBHOST", "localhost")
	os.Setenv("APP_DBPORT", "5432")
	os.Setenv("APP_DBUSER", "postgres")
	os.Setenv("APP_DBPASSWORD", "admin")
	os.Setenv("APP_DBNAME", "postgres")
}

func MustLoad() Config {

	Setenv()

	cfg := Config{}
	if err := envconfig.Process("APP", &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	os.Setenv("STORAGE_CONFIG",
		fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName))

	return cfg
}
