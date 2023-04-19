package dao

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"entgo.io/ent/dialect"
	"fmt"
	"github.com/TBXark/chat-bot-go/configs"
	"github.com/TBXark/chat-bot-go/pkg/dao/ent"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"modernc.org/sqlite"
)

type sqliteDriver struct {
	*sqlite.Driver
}

func (d sqliteDriver) Open(name string) (driver.Conn, error) {
	conn, err := d.Driver.Open(name)
	if err != nil {
		return conn, err
	}
	c := conn.(interface {
		Exec(stmt string, args []driver.Value) (driver.Result, error)
	})
	if _, err := c.Exec("PRAGMA foreign_keys = on;", nil); err != nil {
		_ = conn.Close()
		return nil, fmt.Errorf("failed to enable enable foreign keys: %w", err)
	}
	return conn, nil
}

func init() {
	sql.Register(dialect.SQLite, sqliteDriver{Driver: &sqlite.Driver{}})
}

func NewDatabase(config *configs.Config) *ent.Client {
	client, err := ent.Open(config.Database.Type, config.Database.Path)
	if err != nil {
		log.Fatalf("failed opening connection to %s: %v", config.Database.Type, config.Database.Path)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func WithTx[T any](db *ent.Client, exe func(tx *ent.Tx) (*T, error)) (*T, error) {
	tx, tErr := db.BeginTx(context.Background(), nil)
	if tErr != nil {
		return nil, tErr
	}
	defer func() {
		if e := recover(); e != nil {
			_ = tx.Rollback()
			return
		}
	}()
	result, err := exe(tx)
	if err != nil {
		if rErr := tx.Rollback(); rErr != nil {
			return nil, rErr
		}
		return nil, err
	}
	if cErr := tx.Commit(); cErr != nil {
		return nil, cErr
	}
	return result, err
}
