package config

import (
	"log"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	Env        string `yaml:"env" env:"ENV" env-default:"development"`
	DbPath     string `yaml:"db_path" env-required:"true"`
	HTTPServer `yaml:"server_config"`
}

type HTTPServer struct {
	Address        string        `yaml:"address" env-default:"localhost:3000"`
	RequestTimeout time.Duration `yaml:"req_timeout" env-default:"4s"`
	IdleTimeout    time.Duration `yaml:"idle_timeout" env-default:"60s"`
}

func MustLoad() *Config {
	cfgPath := os.Getenv("CONFIG_PATH")
	if cfgPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	if _, err := os.Stat(cfgPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", cfgPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		log.Fatalf("cannot read config: %s", err)
	}

	return &cfg
}
