package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joseph-beck/aircraft-departure-simulator/ads-backend/generated/sqlc"
)

type Credentials struct {
	// Hostname of the database.
	// For example localhost, or an IP.
	Host string
	// Port number of the database.
	// For example default for Postgres is 5432.
	Port string
	// Username for the database.
	User string
	// Password for the database.
	Password string
}

func NewCredentials(host, port, user, password string) Credentials {
	return Credentials{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
	}
}

func DefaultCredentials() Credentials {
	host := os.Getenv("PG_HOST")
	port := os.Getenv("PG_PORT")
	user := os.Getenv("PG_USER")
	pass := os.Getenv("PG_PASS")

	return NewCredentials(host, port, user, pass)
}

type Connecter interface {
	Connect() (DB, error)
}

type pgConnecter struct {
	// Name of the database being connected to.
	// This is just for logging purposes.
	name string
	// Credentials for connecting to the database.
	creds Credentials
}

// Generate a new Postgres connecter.
func NewPGConnecter(name string, creds Credentials) *pgConnecter {
	return &pgConnecter{
		name:  name,
		creds: creds,
	}
}

// Generate a connection url from the connection credentials.
func (p *pgConnecter) url() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s",
		p.creds.User,
		p.creds.Password,
		p.creds.Host,
		p.creds.Port,
	)
}

// Connect to the Postgres database.
func (p *pgConnecter) Connect() (DB, error) {
	conn, err := pgx.Connect(context.Background(), p.url())
	if err != nil {
		return DB{}, err
	}

	return DB{
		ctx:     context.Background(),
		conn:    conn,
		queries: sqlc.New(conn),
	}, nil
}
