package timer

import (
	"fmt"
	"time"

	"github.com/gosuri/uilive"
    "github.com/masaya-nishimura-09/pomodoro/model"
)

func Timer(timer model.Timer) {
	arrowPosition := 2300

	writer := uilive.New()
	writer.Start()

	for range timer.Time {
		newData := make([]byte, 0, len(timer.Data)+1)
		newData = append(newData, timer.Data[:arrowPosition+1]...)
		newData = append(newData, []byte("^")...)
		newData = append(newData, timer.Data[arrowPosition+2:]...)

		fmt.Fprintf(writer, "%s", string(newData))

		arrowPosition = arrowPosition + 2

		time.Sleep(time.Minute)
	}

	writer.Stop()
}
