package config

type Server struct {
	Address string `env:"ADDRESS"                yaml:"address"`
}

type Postgres struct {
	Dsn string `env:"POSTGRES_DSN" yaml:"dsn"`
}

type Config struct {
	Server   Server   `yaml:"server"`
	Postgres Postgres `yaml:"postgres"`
}
