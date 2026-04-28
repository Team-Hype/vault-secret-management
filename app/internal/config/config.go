package config

import "fmt"

type Config struct {
	Postgres Postgres
}

type Postgres struct {
	Host     string `env:"POSTGRES_HOST"`
	Port     string `env:"POSTGRES_PORT"`
	DB       string `env:"POSTGRES_DB"`
	User     string `env:"POSTGRES_USER"`
	Password string `env:"POSTGRES_PASSWORD"`
}

func (p *Postgres) ToDSN() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		p.User,
		p.Password,
		p.Host,
		p.Port,
		p.DB,
	)
}
