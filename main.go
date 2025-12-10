package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gosuri/uilive"
)

type Session struct {
	count int
}

func main() {
	f, textOpenErr := os.Open("gopher4.txt")
	if textOpenErr != nil {
		fmt.Println(textOpenErr)
		fmt.Println("fail to open file")
	}

	data := make([]byte, 10000)
	count, fileReadErr := f.Read(data)
	if fileReadErr != nil {
		fmt.Println(fileReadErr)
		fmt.Println("fail to read file")
	}

	fmt.Println(string(data[:count]))

	session := Session{}
	for {
		session.AddCount()
		switch session.count {
		case 1, 3, 5, 7:
			timer(24, "Time to focus")
		case 2, 4, 6:
			timer(4, "Short break")
		case 8:
			timer(14, "Long break")
			session.count = 0
		}
	}
}

func (s *Session) AddCount() {
	s.count++
}

func timer(minutes int, status string) {
	min := minutes
	sec := 59

	writer := uilive.New()
	writer.Start()

	for m := min; m >= 0; m-- {
		for s := sec; s >= 0; s-- {
			formattedM := strconv.Itoa(m)
			formattedS := strconv.Itoa(s)
			if m < 10 {
				formattedM = "0" + strconv.Itoa(m)
			}
			if s < 10 {
				formattedS = "0" + strconv.Itoa(s)
			}

			fmt.Fprintf(writer, "%s %s:%s\n", status, formattedM, formattedS)
			time.Sleep(time.Second)
		}
	}

	writer.Stop()
}
