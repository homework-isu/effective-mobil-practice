package config

import (
	"fmt"
	"os"
	"strconv"

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
	db_host            string
	db_port            string
	db_name            string
	db_user            string
	db_pass            string
	db_ssl_mode        string
	db_max_connections int

	http_port string
}

// postgres://jack:secret@pg.example.com:5432/mydb?sslmode=verify-ca&pool_max_conns=10
func (cfg Config) GetConnectionUrl() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		cfg.db_user, cfg.db_pass, cfg.db_host, cfg.db_port, cfg.db_name, cfg.db_ssl_mode)
}

func (cfg Config) GetMaxConnections() int {
	return cfg.db_max_connections
}

func (cfg Config) GetHttpPort() string {
	return cfg.http_port
}

func NewConfig() *Config {
	cfg := &Config{
		db_host:            "localhost",
		db_port:            "5432",
		db_name:            "postgres",
		db_user:            "postgres",
		db_pass:            "",
		db_ssl_mode:        "disable",
		db_max_connections: 10,
		http_port: "8080",
	}

	host := os.Getenv("POSTGRES_HOST")
	if host != "" {
		cfg.db_host = host
	}
	port := os.Getenv("POSTGRES_PORT")
	if port != "" {
		cfg.db_port = port
	}
	name := os.Getenv("POSTGRES_DB")
	if name != "" {
		cfg.db_name = name
	}
	user := os.Getenv("POSTGRES_USER")
	if user != "" {
		cfg.db_user = user
	}
	password := os.Getenv("POSTGRES_PASSWORD")
	if password != "" {
		cfg.db_pass = password
	}
	ssl_mode := os.Getenv("POSTGRES_SSL_MODE")
	if ssl_mode != "" {
		cfg.db_ssl_mode = ssl_mode
	}
	max_cons_str := os.Getenv("POOL_MAX_COONS")
	if max_cons_str != "" {
		val, err := strconv.Atoi(max_cons_str)
		if err != nil {
			cfg.db_max_connections = val
		}
	}
	http_port := os.Getenv("HTTP_PORT")
	if http_port != "" {
		cfg.http_port = http_port
	}

	return cfg
}
