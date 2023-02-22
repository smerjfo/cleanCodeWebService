package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
	"strings"
	"time"
)

type Settings struct {
	Host     string
	Port     int16
	Database string
	User     string
	Password string
	SSLMode  string
}

type Store struct {
	Pool *pgx.Conn
}

func (s Settings) toDSN() string {
	var args []string
	if len(s.Host) > 0 {
		args = append(args, fmt.Sprintf("host=%s", s.Host))
	}
	if s.Port > 0 {
		args = append(args, fmt.Sprintf("port=%d", s.Port))
	}
	if len(s.Database) > 0 {
		args = append(args, fmt.Sprintf("database=%s", s.Database))
	}
	if len(s.User) > 0 {
		args = append(args, fmt.Sprintf("user=%s", s.User))
	}
	if len(s.Password) > 0 {
		args = append(args, fmt.Sprintf("password=%s", s.Password))
	}
	if len(s.SSLMode) > 0 {
		args = append(args, fmt.Sprintf("sslmode=%s", s.SSLMode))
	}
	return strings.Join(args, " ")
}

func init() {
	if err := initDefaultEnv(); err != nil {
		panic(err)
	}
}

func initDefaultEnv() error {
	if len(os.Getenv("PGHOST")) == 0 {
		if err := os.Setenv("PGHOST", "postgres"); err != nil {
			return err
		}
	}
	if len(os.Getenv("PGPORT")) == 0 {
		if err := os.Setenv("PGPORT", "5432"); err != nil {
			return err
		}
	}
	if len(os.Getenv("PGDATABASE")) == 0 {
		if err := os.Setenv("PGDATABASE", "postgres"); err != nil {
			return err
		}
	}
	if len(os.Getenv("PGUSER")) == 0 {
		if err := os.Setenv("PGUSER", "postgres"); err != nil {
			return err
		}
	}
	if len(os.Getenv("PGPASSWORD")) == 0 {
		if err := os.Setenv("PGPASSWORD", "password"); err != nil {
			return err
		}
	}
	if len(os.Getenv("PGSSLMODE")) == 0 {
		if err := os.Setenv("PGSSLMODE", "disable"); err != nil {
			return err
		}
	}
	return nil
}

func New(settings Settings) (*Store, error) {
	config, err := pgx.ParseConfig(settings.toDSN())
	if err != nil {
		return nil, err
	}
	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err = conn.Ping(ctx); err != nil {
		return nil, err
	}
	return &Store{conn}, nil
}
