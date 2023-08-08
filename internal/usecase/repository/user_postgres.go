package repository

import (
	"context"
	"fmt"
	"log"
	"myapp/internal/model"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
)

func NewPostgresDB(Bun *bun.DB) *Postgres {
	return &Postgres{
		Bun: Bun,
	}
}

type Postgres struct {
	Bun *bun.DB
}

func (db *Postgres) CreateUsers(ctx context.Context, users []model.User) error {
	_, err := db.Bun.NewInsert().Model(&users).Exec(context.Background())
	if err != nil {
		log.Println(err)
		return fmt.Errorf("Postgres - CreateUsers - db.Bun.NewInsert: %w", err)
	}
	return nil
}
