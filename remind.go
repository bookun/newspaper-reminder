package newspaperreminder

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
		&model.Man{ID: "UAKK9NV0C"},
		&model.Man{ID: "U9VRB0LUA"},
		&model.Man{ID: "U9ZC5MLCB"},
		&model.Girl{ID: "UJ5KBMZ1C"},
		&model.Man{ID: "UAKAEHGSD"},
		&model.Man{ID: "UJ4FQ8LLX"},
	}
	if uid := r.PostFormValue("user_id"); uid != "" {
		for index, person := range people {
			if person.GetID() == uid {
				if err := slack.Send(person.GetResetMessage()); err != nil {
					fmt.Fprintln(w, err)
				}
				people = append(people[:index], people[index+1:]...)
			}
		}
	}
	picker := people[rand.Intn(len(people))]
	if err := slack.Send(picker.GetMessage()); err != nil {
		fmt.Fprintln(w, err)
	}
}
