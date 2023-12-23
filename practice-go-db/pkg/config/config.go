package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

func LoadEnv(filenames ...string) error {
	const op = "pkg.config.LoadEnv"
	err := godotenv.Load(filenames...)
	if err != nil {
		return fmt.Errorf("%s: %s", op, err)
	}
	return nil
}

type Config struct {
	DB_host            string
	DB_port            string
	DB_name            string
	DB_user            string
	DB_pass            string
	DB_ssl_mode        string
	DB_max_connections int

	Http_port   string
	HttpTimeOut time.Duration
}

// postgres://jack:secret@pg.example.com:5432/myDB?sslmode=verify-ca&pool_max_conns=10
func (cfg Config) GetConnectionUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.DB_user, cfg.DB_pass, cfg.DB_host, cfg.DB_port, cfg.DB_name, cfg.DB_ssl_mode)
}

func (cfg Config) GetMaxConnections() int {
	return cfg.DB_max_connections
}

func (cfg Config) GetHttpPort() string {
	return cfg.Http_port
}

func NewConfig() *Config {
	cfg := &Config{
		DB_host:            "localhost",
		DB_port:            "5432",
		DB_name:            "postgres",
		DB_user:            "postgres",
		DB_pass:            "",
		DB_ssl_mode:        "disable",
		DB_max_connections: 10,
		Http_port:          "8080",
		HttpTimeOut:        5 * time.Second,
	}

	host := os.Getenv("POSTGRES_HOST")
	if host != "" {
		cfg.DB_host = host
	}
	port := os.Getenv("POSTGRES_PORT")
	if port != "" {
		cfg.DB_port = port
	}
	name := os.Getenv("POSTGRES_DB")
	if name != "" {
		cfg.DB_name = name
	}
	user := os.Getenv("POSTGRES_USER")
	if user != "" {
		cfg.DB_user = user
	}
	password := os.Getenv("POSTGRES_PASSWORD")
	if password != "" {
		cfg.DB_pass = password
	}
	ssl_mode := os.Getenv("POSTGRES_SSL_MODE")
	if ssl_mode != "" {
		cfg.DB_ssl_mode = ssl_mode
	}
	max_cons_str := os.Getenv("POOL_MAX_COONS")
	if max_cons_str != "" {
		val, err := strconv.Atoi(max_cons_str)
		if err != nil {
			cfg.DB_max_connections = val
		}
	}
	Http_port := os.Getenv("HTTP_PORT")
	if Http_port != "" {
		cfg.Http_port = Http_port
	}
	timeOut := os.Getenv("HTTP_TIMEOUT")
	if max_cons_str != "" {
		val, err := strconv.Atoi(timeOut)
		if err != nil && val > 0 {
			cfg.DB_max_connections = val
		}
		cfg.HttpTimeOut = time.Duration(val)
	}

	return cfg
}
