package dao

import (
	"TatkalBookingGo/internal/db"
	"TatkalBookingGo/internal/model"
	"database/sql"
	"fmt"
	"log"
)

type UserDao struct {
	//tableName string
	db db.Database
}

func NewUserDao(db db.Database) *UserDao {
	return &UserDao{db: db}
}

func (d *UserDao) GetAllUsers() ([]model.User, error) {
	database := d.db.GetDatabase()
	query := fmt.Sprintf("SELECT * FROM users")
	rows, err := database.Query(query)
	if err != nil {
		return nil, err
	}
	var users []model.User
	for rows.Next() {
		u := model.User{}
		err = rows.Scan(&u.Uid, &u.FirstName, &u.LastName, &u.Age)
		if err != nil {
			log.Println("Error scanning row", err)
		}
		users = append(users, u)
	}
	return users, nil
}

func (d *UserDao) GetUserById(userId int64) (model.User, error) {
	database := d.db.GetDatabase()
	row := database.QueryRow("SELECT * FROM users WHERE uid = $1", userId)
	if row.Err() != nil {
		return model.User{}, row.Err()
	}
	u := model.User{}
	err := row.Scan(&u.Uid, &u.FirstName, &u.LastName, &u.Age)
	if err != nil {
		log.Println("Error scanning row", err)
	}
	return u, nil
}

func (d *UserDao) AddUser(firstName, lastName string, age int) error {
	database := d.db.GetDatabase()
	_, err := database.Exec("INSERT INTO users(first_name, last_name, age) VALUES ($1, $2, $3)",
		firstName, lastName, age)
	if err != nil {
		return err
	}
	return nil
}

type TrainDao struct {
	//tableName string
	db db.Database
}

func NewTrainDao(db db.Database) *TrainDao {
	return &TrainDao{db: db}
}

func (t *TrainDao) ListAllTrains() ([]model.Train, error) {
	database := t.db.GetDatabase()
	rows, err := database.Query("SELECT *  FROM trains")
	if err != nil {
		return nil, err
	}
	var trains []model.Train
	for rows.Next() {
		t := model.Train{}
		err := rows.Scan(&t.TrainNumber, &t.Src, &t.Dest, &t.Capacity)
		if err != nil {
			log.Println("Error scanning row", err)

		}
		trains = append(trains, t)
	}
	return trains, nil
}

func (t *TrainDao) AddTrain(trainNumber int, from, to string, capacity int) error {
	database := t.db.GetDatabase()
	_, err := database.Exec("INSERT INTO trains(train_number, src, dest, capacity) VALUES ($1, $2, $3, $4)",
		trainNumber, from, to, capacity)
	if err != nil {
		return err
	}
	return nil
}

type BookingDao struct {
	//tableName string
	db db.Database
}

func NewBookingDao(db db.Database) *BookingDao {
	return &BookingDao{db: db}
}

func (b *BookingDao) AddSeat(trainNumber int, seatNumber string) error {
	database := b.db.GetDatabase()
	_, err := database.Exec("INSERT INTO bookings(train_number, seat_number) VALUES ($1, $2)",
		trainNumber, seatNumber)
	if err != nil {
		return err
	}
	return nil
}

func (b *BookingDao) GenerateSeats(trainNumber int) {
	for c := 'A'; c < 'K'; c++ {
		for i := 1; i <= 10; i++ {
			seatNumber := fmt.Sprintf("%c%d", c, i)
			err := b.AddSeat(trainNumber, seatNumber)
			if err != nil {
				log.Println("Error adding seat", seatNumber, "train no:", trainNumber, err)
			}
		}
	}
}

func (bd *BookingDao) BookTicket(userId int) {
	tx, _ := bd.db.GetDatabase().Begin()
	defer tx.Rollback()
	row := tx.QueryRow("SELECT booking_id, train_number, seat_number FROM bookings WHERE uid IS NULL LIMIT 1 FOR UPDATE SKIP LOCKED ")
	b := model.Booking{}
	err := row.Scan(&b.BookingId, &b.TrainNumber, &b.SeatNumber)
	if err == sql.ErrNoRows {
		log.Println("No seats left")
		return
	} else if err != nil {
		log.Fatal("Error row to object", err)
	}
	_, err = tx.Exec("UPDATE bookings SET uid = $1 WHERE uid IS NULL AND seat_number = $2", userId, b.SeatNumber)
	if err != nil {
		log.Fatal("Error booking ticket", err)
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal("Error committing transaction", err)
	}
	log.Printf("%d booked %s in train no %d", userId, b.SeatNumber, b.TrainNumber)
}
