package db

import (
	"database/sql"
	"fmt"
	"interactions/internal/config"
	"log"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// NewPostgresDB инициализирует новое подключение к PostgreSQL с использованием bun и возвращает объект *bun.DB
func NewPostgresDB(cfg config.DatabaseConfig) (*bun.DB, error) {
	// Формирование строки подключения (DSN)
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.SSLMode)

	// Создание SQL соединения через pgdriver
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))

	// Создание bun-клиента на основе *sql.DB
	db := bun.NewDB(sqldb, pgdialect.New())

	// Проверка соединения с базой данных
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("Ошибка проверки соединения с базой данных: %w", err)
	}

	log.Println("Успешное подключение к базе данных с использованием bun")
	return db, nil
}
