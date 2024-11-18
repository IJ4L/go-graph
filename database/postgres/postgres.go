package database

import (
	"context"
	"fmt"
	"os"

	"github.com/ij4l/foodCatalog/util"
	"github.com/jackc/pgx/v5"
)

func ConnectPostgreSql(data util.Config) (db *pgx.Conn, err error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable", data.DBHost, data.DBPort, data.DBUser, data.DBPassword)
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	defer conn.Close(context.Background())

	return conn, nil
}