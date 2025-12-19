package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gosuri/uilive"
)

type Session struct {
	count int
}

func main() {
	focusTimer := "timers/gopher_focus.txt"
	shortBreakTimer := "timers/gopher_short_break.txt"
	longBreakTimer := "timers/gopher_long_break.txt"

	session := Session{}
	for {
		session.AddCount()
		switch session.count {
		case 1, 3, 5, 7:
			timer(25, getAsciiData(focusTimer))
		case 2, 4, 6:
			timer(5, getAsciiData(shortBreakTimer))
		case 8:
			timer(15, getAsciiData(longBreakTimer))
			session.count = 0
		}
	}
}

func (s *Session) AddCount() {
	s.count++
}

func timer(minutes int, data []byte) {
	arrowPosition := 2300

	writer := uilive.New()
	writer.Start()

	for range minutes {
		newData := make([]byte, 0, len(data)+1)
		newData = append(newData, data[:arrowPosition+1]...)
		newData = append(newData, []byte("^")...)
		newData = append(newData, data[arrowPosition+2:]...)

		fmt.Fprintf(writer, "%s", string(newData))

		arrowPosition = arrowPosition + 2

		time.Sleep(time.Minute)
	}

	writer.Stop()
}

func getAsciiData(filePath string) []byte {
	f, textOpenErr := os.Open(filePath)
	if textOpenErr != nil {
		fmt.Println(textOpenErr)
		fmt.Println("fail to open file")
	}

	data := make([]byte, 4000)
	_, fileReadErr := f.Read(data)
	if fileReadErr != nil {
		fmt.Println(fileReadErr)
		fmt.Println("fail to read file")
	}

	return data
}
