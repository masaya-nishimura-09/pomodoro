package notification

import (
	"github.com/Bishwas-py/notify"
	"log"
	"time"
    "os"
    "path/filepath"
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
	ex, err := os.Executable()
	if err != nil {
		log.Printf("Error: %v", err)
	}
	exeDir := filepath.Dir(ex)
	soundPath := filepath.Join(exeDir, "notification", "gopher_sound.wav")
	notification.SetSoundByPath(soundPath)

	// Display notification and wait for response
    _, triggerErr := notification.Trigger()
	if triggerErr != nil {
		log.Printf("Error: %v", triggerErr)
	}

}
