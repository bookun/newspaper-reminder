package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/bookun/newspaper-reminder/driver"
	"github.com/bookun/newspaper-reminder/model"
)

type Sender interface {
	Send(message string) error
}

func Remind(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().UnixNano())
	slack := driver.NewSlack(os.Getenv("channel"), os.Getenv("webhook"))
	people := []model.Person{
		//&model.Man{ID: "UAKS4S2D8"},
		//&model.Man{ID: "UCVLN2MEC"},
		&model.Man{ID: "UDP509ZC7"},
		//&model.Girl{ID: "UJM0Q6H60"},
	}
	picker := people[rand.Intn(len(people))]
	if err := slack.Send(picker.GetMessage()); err != nil {
		fmt.Fprintln(w, err)
	}
}
