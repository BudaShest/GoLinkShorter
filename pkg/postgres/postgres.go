package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

const (
	defMaxPoolSize  int           = 1
	defConnAttempts int           = 10
	defConnTimeout  time.Duration = 5 * time.Second
)

type Postgres struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration
	Pool         *pgxpool.Pool
}

func New(dsn string) (*Postgres, error) {
	pg := &Postgres{
		maxPoolSize:  defMaxPoolSize,
		connAttempts: defConnAttempts,
		connTimeout:  defConnTimeout,
	}

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("parse config error:\t%w", err)
	}

	poolConfig.MaxConns = int32(pg.maxPoolSize)

	for ; pg.connAttempts > 0; pg.connAttempts-- {
		pg.Pool, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
		if err == nil {
			break
		}
		log.Printf("Postgres is trying to connect, attemts left: %d", pg.connAttempts)

		time.Sleep(pg.connTimeout)
	}

	if err != nil {
		return nil, fmt.Errorf("postgres - no conn attempts anymore:\t%w", err)
	}

	return pg, nil

}

func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}
