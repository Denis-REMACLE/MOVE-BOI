package main

import (
	"time"
	"github.com/0xAX/notificator"
)

func doesElementExist(s []string, str string) bool {
	for _, v := range s {
	  if v == str {
		return true
	  }
	}
	return false
  }

func main() {
	var notify = notificator.New(notificator.Options{
		AppName:     "MOVE BOI",
	})

	// Modify here the exercises, reps, sets, messages and exercice days
	// Please keep the number of exercises, pauses and phrases equal
	reps := "8"
	sets := "3"
	exercises := []string{"Push ups", "Crunches", "Plank", "Squats", "Dynamic Lunges"}
	pauses := []string{"Go breathe some fresh air !", "Read from a book !", "Jot some ideas down !", "Solve a rubiks !", "Go make yourself a coffee !"}
	phrases := []string{"TIME TO MOVE BOI !", "I'LL MAKE YOU MOVE !", "STAND UP AND EXERCISE !", "NO PAIN ! NO GAIN !", "YOU CAN DO IT !"}
	exercise_days := []string{"Monday", "Wednesday", "Friday"}
	countdown := len(exercises)
	breaktime := 12

	is_exercise_day := doesElementExist(exercise_days, time.Now().Weekday().String())

	for {
		time.Sleep(time.Hour)
		hours, _, _ := time.Now().Clock()
		if (hours == breaktime){
			notify.Push("GO EAT !", "and remember gluttony is bad for you !", "/etc/move-boi/biceps.png", notificator.UR_CRITICAL)
		} else if (countdown == 0){
			notify.Push("YOU'VE DONE WELL SEE YA TOMORROW !", "Please go walk a bit after work and rest !", "/etc/move-boi/biceps.png", notificator.UR_CRITICAL)
			break
		} else if (is_exercise_day == true){
			countdown--
			notify.Push(phrases[countdown], "Do "+reps+" "+exercises[countdown]+" "+sets+" times", "/etc/move-boi/biceps.png", notificator.UR_CRITICAL)
		} else {
			countdown--
			notify.Push("PAUSE FOR AT LEAST 5 MINS", pauses[countdown], "/etc/move-boi/biceps.png", notificator.UR_CRITICAL)
		}
	}
}
