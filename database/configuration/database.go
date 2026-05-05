package configuration

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

func ConnectDB() *sql.DB {
	if os.Getenv("AWS_LAMBDA_FUNCTION_NAME") == "" {
		_ = godotenv.Load()
	}

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")

	if host == "" || port == "" || name == "" || user == "" {
		log.Fatal("variáveis de banco não configuradas corretamente")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		user,
		password,
		host,
		port,
		name,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("erro ao abrir conexão com banco: ", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("erro ao conectar no banco: ", err)
	}

	return db
} 
