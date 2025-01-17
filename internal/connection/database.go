package connection

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/bayuuat/go-sprint-2/internal/config"
	_ "github.com/lib/pq"
)

func GetDatabase(conf config.Database) *sql.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=%s sslmode=disable",
		conf.Host,
		conf.User,
		conf.Pass,
		conf.Name,
		conf.Port,
		conf.Tz,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("failed open connection to db: ", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("failed open connection to db: ", err.Error())
	}

	return db
}
