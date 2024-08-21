package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"srvo-cntrllr/database"
	"srvo-cntrllr/routers"

	_ "github.com/lib/pq"
)

// input
// go get -u "github.com/gin-gonic/gin"
// go get -u "github.com/lib/pq"
// go get -u "github.com/rubenv/sql-migrate"
// go get -u "github.com/joho/godotenv"

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "FARHAN"
// 	dbName   = "praktikum_mcs_bab7"
// )

func main() {
	var PORT = ":" + os.Getenv("PORT")

	// psqlInfo := fmt.Sprintf(
	// 	`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
	// 	host, port, user, password, dbName,
	// )

	psqlInfo := os.Getenv("DATABASE_URL")

	DB, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalf("Error Open DB: %v\n", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Eror pinging DB : %v\n", err)
	}

	database.DBMigrate(DB)

	defer DB.Close()

	routers.StartServer().Run(PORT)
	fmt.Println("berhasil konek...")
}
