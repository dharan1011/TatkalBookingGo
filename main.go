package main

import (
	"TatkalBookingGo/internal/dao"
	"TatkalBookingGo/internal/db"
	"TatkalBookingGo/internal/model"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"sync"
)

func main() {
	pg, err := db.NewPostgresConnection("localhost", "postgres", "postgres", "railways")
	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {
			log.Println("Error closing database connection", err)
		}
		log.Println("database connection closed")
	}(pg.GetDatabase())
	if err != nil {
		log.Fatal("Error created postgres connection object", err)
	}
	userDao := dao.NewUserDao(pg)
	users, err := userDao.GetAllUsers()
	bookingDao := dao.NewBookingDao(pg)
	if err != nil {
		log.Println("Error getting all users")
	}
	wg := sync.WaitGroup{}
	for i, u := range users {
		wg.Add(1)
		go func(u model.User, bd *dao.BookingDao) {
			defer wg.Done()
			bd.BookTicket(u.Uid)
		}(u, bookingDao)
		if i >= 5 {
			break
		}
	}
	wg.Wait()
}
