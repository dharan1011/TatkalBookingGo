package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

const (
	PostgresPort = 5432
)

type Database interface {
	StopHealthCheck()
	StartHealthCheck()
	GetDatabase() *sql.DB
}

type Postgres struct {
	db                 *sql.DB
	DatabaseName       string
	healthCheckChannel chan bool
}

func (pg *Postgres) StopHealthCheck() {
	pg.healthCheckChannel <- true
}

func (pg *Postgres) StartHealthCheck() {
	go databaseHealthCheck(pg.db, pg.healthCheckChannel, time.Second*5)
}

func (pg *Postgres) GetDatabase() *sql.DB {
	return pg.db
}

func NewPostgresConnection(host, user, password, dbname string) (*Postgres, error) {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, PostgresPort, user, password, dbname)
	log.Println("Database connection string", connectionString)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Println("Error creating database object")
		return nil, err
	}
	return &Postgres{
		db:                 db,
		DatabaseName:       dbname,
		healthCheckChannel: make(chan bool),
	}, nil
}

func databaseHealthCheck(db *sql.DB, healthCheckChannel <-chan bool, interval time.Duration) {
	for {
		select {
		case r := <-healthCheckChannel:
			if r {
				log.Println("Stopping db health check")
				break
			}
		case <-time.Tick(interval):
			err := db.Ping()
			if err != nil {
				log.Println("db health check failed", err)
			}
		}
	}
}
