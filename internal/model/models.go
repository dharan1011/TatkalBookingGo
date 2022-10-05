package model

type User struct {
	Uid       int    `json:"uid"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

type Train struct {
	TrainNumber int64  `json:"train_number"`
	Src         string `json:"src"`
	Dest        string `json:"dest"`
	Capacity    int    `json:"capacity"`
}

type Booking struct {
	BookingId   int    `json:"booking_id"`
	TrainNumber int    `json:"train_number"`
	SeatNumber  string `json:"seat_number"`
	UserId      int    `json:"user_id"`
}
