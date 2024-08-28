package config

import (
	"context"
	"fmt"
	"github.com/Rawipass/chat-service/logger"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
	"log"
	"net/url"
	"os"
	"time"
)

var DB *pgxpool.Pool

func ConnectDatabase() {
	// Init Database Config
	userName := viper.GetString("Database.Username")
	password := viper.GetString("Database.Password")
	host := viper.GetString("Database.Host")
	port := viper.GetInt("Database.Port")
	databaseName := viper.GetString("Database.DatabaseName")
	databaseSchema := viper.GetString("Database.DatabaseSchema")
	connectionTimeout := viper.GetInt("Database.ConnectionTimeout")
	escapedPassword := url.QueryEscape(password)
	databaseUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?search_path=%s&connect_timeout=%d",
		userName, escapedPassword, host, port, databaseName, databaseSchema, connectionTimeout,
	)
	config, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		log.Fatalf("Unable to parse database URL: %v\n", err)
	}

	config.MaxConns = viper.GetInt32("Database.MaxConnection")
	config.MinConns = viper.GetInt32("Database.MinConnection")
	config.MaxConnLifetime = time.Hour
	dbpool, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		logger.Logger.Errorf("Unable to parse config for database: %s\n", err)
		os.Exit(1)
	}
	DB = dbpool
	log.Println("Database connected!")
}

func DisconnectDatabase() {
	DB.Close()
}
