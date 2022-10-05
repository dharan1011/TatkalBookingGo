package mock

import (
	"TatkalBookingGo/internal/dao"
	"fmt"
	"log"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func GenerateMockTrainData(trainsDao *dao.TrainDao) {
	src := [...]string{
		"CCT",
		"SC",
		"BNC",
		"NZM",
	}
	dest := [...]string{
		"SLO",
		"RJY",
		"HWH",
		"LTT",
	}
	var start = 1700
	for _, s := range src {
		for _, d := range dest {
			err := trainsDao.AddTrain(start, s, d, 100)
			if err != nil {
				log.Println("Error adding train", err)
				continue
			}
			start++
			err = trainsDao.AddTrain(start, d, s, 100)
			if err != nil {
				log.Println("Error adding train", err)
				continue
			}
			start++
		}
	}
}

func GenerateMockUserData(userDao *dao.UserDao) {
	for i := 0; i < 1e3; i++ {
		err := userDao.AddUser(randomString(10), randomString(10), rand.Intn(100)+1)
		if err != nil {
			log.Fatal("Error adding user", err)
		}
	}
}

func randomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, length)
	rand.Read(b)
	return fmt.Sprintf("%x", b)[:length]
}
