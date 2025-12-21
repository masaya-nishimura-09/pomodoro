package notification

import (
	"github.com/Bishwas-py/notify"
	"log"
	"time"
)

func LongBreak() {
	// Create notification with an action button
	notification := notify.Notification{
		AppID:   "Pomodoro",
		Title:   "Time for a long break!",
		Body:    "Click OK to take a long break.",
		Timeout: int(10 * time.Second),
		Actions: notify.Actions{
			{
				Title:   "OK",
                Trigger: func(){},
			},
		},
	}

	// Display notification and wait for response
	_, err := notification.Trigger()
	if err != nil {
		log.Printf("Error: %v", err)
	}
}
